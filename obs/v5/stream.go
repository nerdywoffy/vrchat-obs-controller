package v5

import "github.com/andreykaipov/goobs/api/events"

func (v5 *obsv5) handleStreamUpdate() {
	for event := range v5.client.IncomingEvents {
		switch e := event.(type) {
		case *events.ReplayBufferStateChanged:
			v5.state.IsReplayBuffer = e.OutputActive
		case *events.RecordStateChanged:
			v5.state.IsRecording = e.OutputActive
		case *events.StreamStateChanged:
			v5.state.IsStreaming = e.OutputActive
		}
	}
}

func (v5 *obsv5) GetStatusStream() (bool, error) {
	data, err := v5.client.Stream.GetStreamStatus()
	if err != nil {
		return false, err
	}

	return data.OutputActive, nil
}

func (v5 *obsv5) ToggleStream() error {
	_, err := v5.client.Stream.ToggleStream()
	return err
}
