package v5

import (
	"fmt"

	"github.com/nerdywoffy/vrchat-obs-controller/obs"

	"github.com/andreykaipov/goobs"
	"github.com/sirupsen/logrus"
)

type obsv5 struct {
	log    *logrus.Logger
	client *goobs.Client
	state  obs.OBSState
	scenes map[int]obs.OBSScene
}

func New(log *logrus.Logger) obs.OBSWebsocketAPI {
	return &obsv5{
		log:    log,
		scenes: make(map[int]obs.OBSScene),
	}
}

func (v5 *obsv5) Start(credential obs.OBSCredential) error {
	options := []goobs.Option{}

	// Add password to options
	v5.log.Infof("Connecting to OBS Websocket v5 on %s:%d", credential.Host, credential.Port)
	if len(credential.Password) > 0 {
		v5.log.Infoln("Password set, will connecting with password!")
		options = append(options, goobs.WithPassword(credential.Password))
	}

	// Start OBS v5 Client
	client, err := goobs.New(
		fmt.Sprintf("%s:%d", credential.Host, credential.Port),
		options...,
	)
	if err != nil {
		return err
	}
	v5.client = client

	go v5.handleStreamUpdate()
	return nil
}

func (v5 *obsv5) GetState() (obs.OBSState, error) {
	return v5.state, nil
}
