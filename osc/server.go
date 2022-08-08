package osc

import (
	"fmt"

	goosc "github.com/hypebeast/go-osc/osc"
	"github.com/sirupsen/logrus"
)

type Server struct {
	log        *logrus.Logger
	dispatcher goosc.Dispatcher
	server     *goosc.Server
}

func NewServer(log *logrus.Logger, server *goosc.Server, dispatcher goosc.Dispatcher) *Server {
	return &Server{
		log:        log,
		dispatcher: dispatcher,
		server:     server,
	}
}

func (s *Server) Start() {
	go func() {
		if err := s.server.ListenAndServe(); err != nil {
			fmt.Println(err)
		}
	}()
}

func (s *Server) Server() *goosc.Server {
	return s.server
}

func (s *Server) AddBooleanListener(addr string, callback func(bool)) {
	s.log.Debugf("Listening %s as Boolean Parameter", addr)
	s.dispatcher.(*goosc.StandardDispatcher).AddMsgHandler(addr, func(msg *goosc.Message) {
		s.log.Debugf("Incoming %s message with %d arguments", addr, msg.CountArguments())
		if msg.CountArguments() == 0 {
			return
		}

		if v, ok := msg.Arguments[0].(bool); ok {
			callback(v)
			return
		}
	})
}

func (s *Server) AddIntegerListener(addr string, callback func(int)) {
	s.dispatcher.(*goosc.StandardDispatcher).AddMsgHandler(addr, func(msg *goosc.Message) {
		s.log.Debugf("Incoming %s message with %d arguments: %+v", addr, msg.CountArguments(), msg.Arguments)
		if msg.CountArguments() == 0 {
			return
		}

		switch msg.Arguments[0].(type) {
		case int32:
			callback(int(msg.Arguments[0].(int32)))
			return
		case int64:
			callback(int(msg.Arguments[0].(int64)))
			return
		}

	})
}
