package v4

import (
	"errors"
	"fmt"

	"github.com/nerdywoffy/vrchat-obs-controller/obs"

	"github.com/andreykaipov/goobs"
	"github.com/andreykaipov/goobs/api/events"
	"github.com/andreykaipov/goobs/api/requests/scenes"
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

func (v4 *obsv4) handleStreamUpdate() {
	for event := range v4.client.IncomingEvents {
		switch e := event.(type) {
		case *events.ReplayStarted:
			v4.state.IsReplayBuffer = true
		case *events.ReplayStopped:
			v4.state.IsReplayBuffer = false
		case *events.RecordingStarted:
			v4.state.IsRecording = true
		case *events.RecordingStopped:
			v4.state.IsRecording = false
		case *events.StreamStarted:
			v4.state.IsStreaming = true
		case *events.StreamStopped:
			v4.state.IsStreaming = false
		default:
			v4.log.Debugf("Unhandled Incoming Events: %s", e.GetUpdateType())
		}
	}
}

func (v4 *obsv4) GetStatusStream() (bool, error) {
	data, err := v4.client.Streaming.GetStreamingStatus()
	if err != nil {
		return false, err
	}

	return data.Streaming, nil
}

func (v4 *obsv4) ToggleStream() error {
	_, err := v4.client.Streaming.StartStopStreaming()
	return err
}

func (v4 *obsv4) GetStatusRecord() (bool, error) {
	data, err := v4.client.Recording.GetRecordingStatus()
	if err != nil {
		return false, err
	}

	return data.IsRecording, nil
}

func (v4 *obsv4) ToggleRecord() error {
	_, err := v4.client.Recording.StartStopRecording()
	return err
}

func (v4 *obsv4) GetStatusInstantReplay() (bool, error) {
	data, err := v4.client.ReplayBuffer.GetReplayBufferStatus()
	if err != nil {
		return false, err
	}

	return data.IsReplayBufferActive, nil
}

func (v4 *obsv4) ToggleInstantReplay() error {
	_, err := v4.client.ReplayBuffer.StartStopReplayBuffer()
	return err
}

func (v4 *obsv4) SaveInstantReplay() error {
	_, err := v4.client.ReplayBuffer.SaveReplayBuffer()
	return err
}

func (v4 *obsv4) SetScene(scenes []obs.OBSScene) error {
	// Assign scene number
	for _, scene := range scenes {
		if scene.Number <= 0 {
			continue
		}
		v4.scenes[scene.Number] = scene
	}

	v4.log.Debugf("Registered Scenes: %+v", v4.scenes)

	return nil
}

func (v4 *obsv4) ToggleToSceneName(sceneName string) error {
	for _, scene := range v4.scenes {
		if scene.Name == sceneName {
			v4.client.Scenes.SetCurrentScene(&scenes.SetCurrentSceneParams{
				SceneName: scene.Name,
			})
			return nil
		}
	}

	return errors.New("scene not available or not registered")
}

func (v4 *obsv4) ToggleToSceneNumber(number int) error {
	if v, ok := v4.scenes[number]; ok {
		v4.client.Scenes.SetCurrentScene(&scenes.SetCurrentSceneParams{
			SceneName: v.Name,
		})
		return nil
	}

	return errors.New("scene not available or not registered")
}

func (v4 *obsv4) GetState() (obs.OBSState, error) {
	return v4.state, nil
}

func (v4 *obsv4) GetCurrentSceneNumber() (int, error) {
	scenes, err := v4.client.Scenes.GetCurrentScene()
	if err != nil {
		return 0, err
	}

	// Find Scene Name based on registered scene
	for num, scene := range v4.scenes {
		if scene.Name == scenes.Name {
			return num, nil
		}
	}

	return 0, nil
}
