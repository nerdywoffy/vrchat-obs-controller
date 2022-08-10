package v4

import (
	"errors"

	"github.com/andreykaipov/goobs/api/requests/scenes"
	"github.com/nerdywoffy/vrchat-obs-controller/obs"
)

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
