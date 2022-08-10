package streamlabs

import (
	"errors"
	"fmt"

	"github.com/nerdywoffy/vrchat-obs-controller/obs"
	"github.com/nerdywoffy/vrchat-obs-controller/slobsrpc"

	"github.com/sirupsen/logrus"
)

type slobs struct {
	log    *logrus.Logger
	rpc    *slobsrpc.SLOBSRPC
	state  obs.OBSState
	scenes map[int]obs.OBSScene

	sceneToIdMap map[string]string
}

func New(log *logrus.Logger) obs.OBSWebsocketAPI {
	return &slobs{
		log:          log,
		scenes:       make(map[int]obs.OBSScene),
		sceneToIdMap: make(map[string]string),
	}
}

func (sl *slobs) Start(credential obs.OBSCredential) error {
	// Create JSON RPC Connection
	sl.log.Debugln("Dialing SLOBS")
	sl.rpc = slobsrpc.New(sl.log)
	err := sl.rpc.Dial(credential.Host, int(credential.Port))
	if err != nil {
		return err
	}
	sl.log.Debugln("Connected to SLOBS")

	// Auth Myself
	authOutput, rpcError, err := sl.rpc.Send("auth", map[string]interface{}{
		"resource": "TcpServerService",
		"args":     []string{credential.Password},
	})
	if err != nil {
		return err
	}

	// Check if RPC error returns an error?
	if rpcError != nil {
		return fmt.Errorf("got error from RPC: %s", rpcError.Message)
	}

	// Lazy check, compare with string
	if string(authOutput) != "true" {
		return errors.New("authentication failed")
	}

	return nil
}

func (sl *slobs) GetState() (obs.OBSState, error) {
	return sl.state, nil
}
