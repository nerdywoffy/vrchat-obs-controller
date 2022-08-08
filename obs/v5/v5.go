package v5

import (
	"github.com/nerdywoffy/vrchat-obs-controller/obs"

	"github.com/sirupsen/logrus"
)

type obsv5 struct {
	state obs.OBSState
}

func New(log *logrus.Logger) obs.OBSWebsocketAPI {
	return &obsv5{}
}

func (v4 *obsv5) Start(obs.OBSCredential) error {
	panic("not implemented yet")
}

func (v4 *obsv5) ToggleStream() error {
	panic("not implemented yet")

}

func (v4 *obsv5) ToggleRecord() error {
	panic("not implemented yet")

}

func (v4 *obsv5) ToggleInstantReplay() error {
	panic("not implemented yet")

}

func (v4 *obsv5) SaveInstantReplay() error {
	panic("not implemented yet")

}

func (v4 *obsv5) SetScene(scenes []obs.OBSScene) error {
	panic("not implemented yet")
}

func (v4 *obsv5) ToggleToSceneName(sceneName string) error {
	panic("not implemented yet")
}

func (v4 *obsv5) ToggleToSceneNumber(number int) error {
	panic("not implemented yet")
}

func (v5 *obsv5) GetState() (obs.OBSState, error) {
	return v5.state, nil
}

func (v5 *obsv5) GetStatusStream() (bool, error) {
	panic("not implemented yet")
}

func (v5 *obsv5) GetStatusRecord() (bool, error) {
	panic("not implemented yet")
}

func (v5 *obsv5) GetStatusInstantReplay() (bool, error) {
	panic("not implemented yet")
}

func (v5 *obsv5) GetCurrentSceneNumber() (int, error) {
	panic("not implemented yet")
}
