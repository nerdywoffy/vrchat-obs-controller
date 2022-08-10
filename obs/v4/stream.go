package v4

import "github.com/andreykaipov/goobs/api/events"

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
