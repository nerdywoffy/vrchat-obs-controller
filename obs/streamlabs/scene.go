package streamlabs

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/nerdywoffy/vrchat-obs-controller/obs"
)

func (sl *slobs) SetScene(scenes []obs.OBSScene) error {
	var slobsScenes StreamlabsScenesResult

	// Fetch RPC
	scenesList, rpcError, err := sl.rpc.Send("getScenes", map[string]interface{}{
		"resource": "ScenesService",
	})
	if err != nil {
		return err
	}

	if rpcError != nil {
		return fmt.Errorf("getScenes returning error: %s", rpcError.Message)
	}

	// Parse all Scenes list
	if err := json.Unmarshal(scenesList, &slobsScenes); err != nil {
		return err
	}

	for _, scene := range slobsScenes {
		sl.sceneToIdMap[scene.Name] = scene.ID
	}

	for _, scene := range scenes {
		if scene.Number <= 0 {
			continue
		}
		sl.scenes[scene.Number] = scene
	}

	sl.log.Debugf("Registered Scenes (From SLOBS): %+v", sl.sceneToIdMap)
	sl.log.Debugf("Registered Scenes (From Config): %+v", sl.scenes)
	return nil
}

func (sl *slobs) ToggleToSceneName(sceneName string) error {
	for _, scene := range sl.scenes {
		if scene.Name == sceneName {
			if v, ok := sl.sceneToIdMap[scene.Name]; !!ok {
				return errors.New("scene not available or not registered")
			} else {
				_, rpcError, err := sl.rpc.Send("makeSceneActive", map[string]interface{}{
					"resource": "ScenesService",
					"args":     []string{v},
				})
				if err != nil {
					return err
				}
				if rpcError != nil {
					return fmt.Errorf("makeSceneActive returning error: %s", rpcError.Message)
				}
			}
		}
	}

	return errors.New("scene not available or not registered")
}

func (sl *slobs) ToggleToSceneNumber(number int) error {
	if v, ok := sl.scenes[number]; ok {
		sl.log.Debugf("Got Scene Name: %s", v.Name)
		if v, ok := sl.sceneToIdMap[v.Name]; !ok {
			return errors.New("scene not available or not registered")
		} else {
			_, rpcError, err := sl.rpc.Send("makeSceneActive", map[string]interface{}{
				"resource": "ScenesService",
				"args":     []string{v},
			})
			if err != nil {
				return err
			}
			if rpcError != nil {
				return fmt.Errorf("makeSceneActive returning error: %s", rpcError.Message)
			}
		}
		return nil
	}

	return errors.New("scene not available or not registered")
}

func (sl *slobs) GetCurrentSceneNumber() (int, error) {
	var (
		slobsScenes      StreamlabsScene
		currentSceneName string
	)
	// Fetch RPC
	scenesList, rpcError, err := sl.rpc.Send("activeScene", map[string]interface{}{
		"resource": "ScenesService",
	})
	if err != nil {
		return 0, err
	}

	if rpcError != nil {
		return 0, fmt.Errorf("activeScene returning error: %s", rpcError.Message)
	}

	// Parse all Scenes list
	if err := json.Unmarshal(scenesList, &slobsScenes); err != nil {
		return 0, err
	}

	// Go through key
	for sceneName, sceneId := range sl.sceneToIdMap {
		if slobsScenes.ID == sceneId {
			currentSceneName = sceneName
			break
		}
	}

	// Find Scene Name based on registered scene
	for num, scene := range sl.scenes {
		if scene.Name == currentSceneName {
			return num, nil
		}
	}

	return 0, nil
}
