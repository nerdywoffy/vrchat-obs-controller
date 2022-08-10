package slobsrpc

import (
	"encoding/json"
	"errors"
	"fmt"
	"math/rand"
	"net/url"
	"sync"
	"time"

	"github.com/gorilla/websocket"
	"github.com/sirupsen/logrus"
)

type SLOBSRPC struct {
	log  *logrus.Logger
	conn *websocket.Conn

	counterMutex sync.Mutex
	counter      int64

	ongoingQueue map[int64]*request
}

type SLOBSRPCRequest struct {
	JSONRPCVersion string          `json:"jsonrpc"`
	ID             int64           `json:"id"`
	Method         string          `json:"method"`
	Params         json.RawMessage `json:"params"`
}

type SLOBSRPCResponse struct {
	JSONRPCVersion string          `json:"jsonrpc"`
	ID             int64           `json:"id"`
	Result         json.RawMessage `json:"result,omitempty"`
	Error          *SLOBSRPCError  `json:"error,omitempty"`
}

type SLOBSRPCError struct {
	Code    int64  `json:"code"`
	Message string `json:"message"`
}

type request struct {
	payload  []byte
	response []byte
	err      *SLOBSRPCError
	onFinish chan error
}

func New(log *logrus.Logger) *SLOBSRPC {
	var rpc SLOBSRPC
	rpc.log = log
	rpc.ongoingQueue = make(map[int64]*request)

	return &rpc
}

func (rpc *SLOBSRPC) Dial(host string, port int) error {
	// Generate URL
	parsed, err := url.Parse(generateURL(host, port))
	if err != nil {
		return err
	}

	conn, _, err := websocket.DefaultDialer.Dial(parsed.String(), nil)
	if err != nil {
		return err
	}
	rpc.conn = conn

	go rpc.readMessage()
	return nil
}

func (rpc *SLOBSRPC) Send(method string, params interface{}) ([]byte, *SLOBSRPCError, error) {
	var rpcResponse []byte

	// Get Request ID
	id := rpc.GetCounter()
	// Encode params to JSON
	rpcParams, err := json.Marshal(&params)
	if err != nil {
		return nil, nil, err
	}

	// Create RPC Request
	rpcRequestObj := SLOBSRPCRequest{JSONRPCVersion: "2.0", ID: id, Method: method, Params: rpcParams}
	rpcRequest, err := json.Marshal(&rpcRequestObj)
	if err != nil {
		return nil, nil, err
	}

	// Wrap request inside array of string
	rpcOutputObj := []string{string(rpcRequest)}
	rpcOutput, err := json.Marshal(rpcOutputObj)
	if err != nil {
		return nil, nil, err
	}

	// Build Request
	ding := make(chan error, 1)
	outRequest := request{payload: rpcOutput, response: rpcResponse, onFinish: ding}
	go rpc.queue(id, &outRequest)

	queueError := <-ding
	if queueError != nil {
		return nil, nil, queueError
	}

	return outRequest.response, outRequest.err, nil
}

func (rpc *SLOBSRPC) GetCounter() int64 {
	rpc.counterMutex.Lock()
	defer rpc.counterMutex.Unlock()
	rpc.counter += 1
	return rpc.counter
}

func (rpc *SLOBSRPC) queue(id int64, outRequest *request) {
	// Register to queuer
	rpc.ongoingQueue[id] = outRequest

	// Write message
	err := rpc.conn.WriteMessage(websocket.TextMessage, outRequest.payload)
	if err != nil {
		outRequest.onFinish <- err
	}
}

func (rpc *SLOBSRPC) readMessage() {
	for {
		_, message, err := rpc.conn.ReadMessage()
		if err != nil {
			rpc.log.Error(err)
			return
		}
		if err := rpc.parseAndSendMessage(message); err != nil {
			rpc.log.Error(err)
		}
	}
}

func (rpc *SLOBSRPC) parseAndSendMessage(message []byte) error {
	// Ignore anything that doesn't start with a
	if len(message) == 0 || (len(message) > 0 && message[0] != 'a') {
		return nil
	}

	// Remove unnecessary a in front of message
	if len(message) > 0 && message[0] == 'a' {
		message = message[1:]
	}

	// Read as an array of string
	arrs := []string{}
	if err := json.Unmarshal(message, &arrs); err != nil {
		return err
	}

	if len(arrs) == 0 {
		return errors.New("no message detected")
	}

	// Parse JSON
	// rpc.log.Debugln(arrs[0])
	resp := SLOBSRPCResponse{}
	if err := json.Unmarshal([]byte(arrs[0]), &resp); err != nil {
		return err
	}

	// Read if there's any incoming queue for it?
	if target, ok := rpc.ongoingQueue[resp.ID]; ok {
		target.response = resp.Result
		target.err = resp.Error
		target.onFinish <- nil

		delete(rpc.ongoingQueue, resp.ID)
	}

	return nil
}

func generateURL(host string, port int) string {
	letterRunes := []rune("abcdefghijklmnopqrstuvwxyz")

	// Generate three random numbers
	randSource := rand.New(rand.NewSource(time.Now().UnixNano()))
	num := randSource.Intn(999-100) + 100

	// Generate eight random runes
	letter := make([]rune, 8)
	for i := range letter {
		letter[i] = letterRunes[randSource.Intn(len(letterRunes))]
	}

	return fmt.Sprintf("ws://%s:%d/api/%d/%s/websocket", host, port, num, string(letter))
}
