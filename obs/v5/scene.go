package v5

import (
	"errors"

	"github.com/andreykaipov/goobs/api/requests/scenes"
	"github.com/nerdywoffy/vrchat-obs-controller/obs"
)

func (v5 *obsv5) SetScene(scenes []obs.OBSScene) error {
	// Assign scene number
	for _, scene := range scenes {
		if scene.Number <= 0 {
			continue
		}
		v5.scenes[scene.Number] = scene
	}

	v5.log.Debugf("Registered Scenes: %+v", v5.scenes)

	return nil
}

func (v5 *obsv5) ToggleToSceneName(sceneName string) error {
	for _, scene := range v5.scenes {
		if scene.Name == sceneName {
			v5.client.Scenes.SetCurrentProgramScene(&scenes.SetCurrentProgramSceneParams{
				SceneName: scene.Name,
			})
			return nil
		}
	}

	return errors.New("scene not available or not registered")
}

func (v5 *obsv5) ToggleToSceneNumber(number int) error {
	if v, ok := v5.scenes[number]; ok {
		v5.client.Scenes.SetCurrentProgramScene(&scenes.SetCurrentProgramSceneParams{
			SceneName: v.Name,
		})
		return nil
	}

	return errors.New("scene not available or not registered")
}

func (v5 *obsv5) GetCurrentSceneNumber() (int, error) {
	scenes, err := v5.client.Scenes.GetCurrentProgramScene()
	if err != nil {
		return 0, err
	}

	// Find Scene Name based on registered scene
	for num, scene := range v5.scenes {
		if scene.Name == scenes.CurrentProgramSceneName {
			return num, nil
		}
	}

	return 0, nil
}
