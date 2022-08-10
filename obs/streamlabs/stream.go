package streamlabs

import (
	"encoding/json"
	"fmt"
)

func (sl *slobs) handleStreamUpdate() {

}

func (sl *slobs) GetStatusStream() (bool, error) {
	var streamStatus StreamlabsStreamingState

	// Fetch RPC
	scenesList, rpcError, err := sl.rpc.Send("getModel", map[string]interface{}{
		"resource": "StreamingService",
	})
	if err != nil {
		return false, err
	}

	if rpcError != nil {
		return false, fmt.Errorf("getScenes returning error: %s", rpcError.Message)
	}

	// Parse all Scenes list
	if err := json.Unmarshal(scenesList, &streamStatus); err != nil {
		return false, err
	}

	return streamStatus.StreamingStatus == "live", nil
}

func (sl *slobs) ToggleStream() error {
	_, rpcError, err := sl.rpc.Send("toggleStreaming", map[string]interface{}{
		"resource": "StreamingService",
	})
	if err != nil {
		return err
	}
	if rpcError != nil {
		return fmt.Errorf("toggleStreaming returning error: %s", rpcError.Message)
	}

	return nil
}
