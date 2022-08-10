package streamlabs

import (
	"encoding/json"
	"fmt"
)

func (sl *slobs) GetStatusInstantReplay() (bool, error) {
	var streamStatus StreamlabsStreamingState

	// Fetch RPC
	scenesList, rpcError, err := sl.rpc.Send("getModel", map[string]interface{}{
		"resource": "StreamingService",
	})
	if err != nil {
		return false, err
	}

	if rpcError != nil {
		return false, fmt.Errorf("getModel returning error: %s", rpcError.Message)
	}

	// Parse all Scenes list
	if err := json.Unmarshal(scenesList, &streamStatus); err != nil {
		return false, err
	}

	return (streamStatus.ReplayBufferstatus == "running" || streamStatus.ReplayBufferstatus == "saving"), nil
}

func (sl *slobs) ToggleInstantReplay() error {
	isInstantReplayEnabled, err := sl.GetStatusInstantReplay()
	if err != nil {
		return err
	}

	if isInstantReplayEnabled { // Enabled, need to stop
		_, rpcError, err := sl.rpc.Send("stopReplayBuffer", map[string]interface{}{
			"resource": "StreamingService",
		})
		if err != nil {
			return err
		}
		if rpcError != nil {
			return fmt.Errorf("stopReplayBuffer returning error: %s", rpcError.Message)
		}
	} else { // Disabled, need to start
		_, rpcError, err := sl.rpc.Send("startReplayBuffer", map[string]interface{}{
			"resource": "StreamingService",
		})
		if err != nil {
			return err
		}
		if rpcError != nil {
			return fmt.Errorf("startReplayBuffer returning error: %s", rpcError.Message)
		}
	}

	return nil
}

func (sl *slobs) SaveInstantReplay() error {
	_, rpcError, err := sl.rpc.Send("saveReplay", map[string]interface{}{
		"resource": "StreamingService",
	})
	if err != nil {
		return err
	}
	if rpcError != nil {
		return fmt.Errorf("saveReplay returning error: %s", rpcError.Message)
	}

	return nil
}
