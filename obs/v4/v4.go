package v4

import (
	"fmt"

	"github.com/nerdywoffy/vrchat-obs-controller/obs"

	"github.com/andreykaipov/goobs"
	"github.com/sirupsen/logrus"
)

type obsv4 struct {
	log    *logrus.Logger
	client *goobs.Client
	state  obs.OBSState
	scenes map[int]obs.OBSScene
}

func New(log *logrus.Logger) obs.OBSWebsocketAPI {
	return &obsv4{
		log:    log,
		scenes: make(map[int]obs.OBSScene),
	}
}

func (v4 *obsv4) Start(credential obs.OBSCredential) error {
	options := []goobs.Option{}

	// Add password to options
	v4.log.Infof("Connecting to OBS Websocket v4 on %s:%d", credential.Host, credential.Port)
	if len(credential.Password) > 0 {
		v4.log.Infoln("Password set, will connecting with password!")
		options = append(options, goobs.WithPassword(credential.Password))
	}

	// Start OBS v4 Client
	client, err := goobs.New(
		fmt.Sprintf("%s:%d", credential.Host, credential.Port),
		options...,
	)
	if err != nil {
		return err
	}
	v4.client = client

	go v4.handleStreamUpdate()
	return nil
}

func (v4 *obsv4) GetState() (obs.OBSState, error) {
	return v4.state, nil
}
