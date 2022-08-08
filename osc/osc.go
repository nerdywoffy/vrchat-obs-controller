package osc

import (
	"fmt"

	"github.com/hypebeast/go-osc/osc"
	"github.com/sirupsen/logrus"
)

type OSCProtocol struct {
	client *Client
	server *Server
}

type OscSetting struct {
	Sender   OscSenderSetting
	Receiver OscReceiverSetting
}

type OscSenderSetting struct {
	Host string
	Port int
}

type OscReceiverSetting struct {
	Host string
	Port int
}

func New(log *logrus.Logger, setting OscSetting) *OSCProtocol {
	// Create client
	log.Infof("Running OSC sender on %s port %d", setting.Sender.Host, setting.Sender.Port)
	client := osc.NewClient(setting.Sender.Host, setting.Sender.Port)

	// Create Server, along with dispatcher
	log.Infof("Running OSC receiver on %s port %d", setting.Receiver.Host, setting.Receiver.Port)
	oscServerDispatcher := osc.NewStandardDispatcher()
	oscServer := &osc.Server{
		Addr:       fmt.Sprintf("%s:%d", setting.Receiver.Host, setting.Receiver.Port),
		Dispatcher: oscServerDispatcher,
	}

	return &OSCProtocol{
		client: NewClient(client),
		server: NewServer(log, oscServer, oscServerDispatcher),
	}
}

func (proto *OSCProtocol) GetClient() *Client {
	return proto.client
}

func (proto *OSCProtocol) GetServer() *Server {
	return proto.server
}
