package v4

import (
	"github.com/nerdywoffy/vrchat-obs-controller/obs"

	"github.com/sirupsen/logrus"
)

type obsv4 struct {
	state obs.OBSState
}

func New(log *logrus.Logger) obs.OBSWebsocketAPI {
	panic("Support for OBSv4 (Legacy) dropped, please use v5 instead.")
}

func (v4 *obsv4) Start(obs.OBSCredential) error {
	panic("support dropped")
}

func (v4 *obsv4) ToggleStream() error {
	panic("support dropped")

}

func (v4 *obsv4) ToggleRecord() error {
	panic("support dropped")

}

func (v4 *obsv4) ToggleInstantReplay() error {
	panic("support dropped")

}

func (v4 *obsv4) SaveInstantReplay() error {
	panic("support dropped")

}

func (v4 *obsv4) SetScene(scenes []obs.OBSScene) error {
	panic("support dropped")
}

func (v4 *obsv4) ToggleToSceneName(sceneName string) error {
	panic("support dropped")
}

func (v4 *obsv4) ToggleToSceneNumber(number int) error {
	panic("support dropped")
}

func (v4 *obsv4) GetState() (obs.OBSState, error) {
	return v4.state, nil
}

func (v4 *obsv4) GetStatusStream() (bool, error) {
	panic("support dropped")
}

func (v4 *obsv4) GetStatusRecord() (bool, error) {
	panic("support dropped")
}

func (v4 *obsv4) GetStatusInstantReplay() (bool, error) {
	panic("support dropped")
}

func (v4 *obsv4) GetCurrentSceneNumber() (int, error) {
	panic("support dropped")
}
