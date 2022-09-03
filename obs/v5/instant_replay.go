package v5

func (v5 *obsv5) GetStatusInstantReplay() (bool, error) {
	data, err := v5.client.Outputs.GetReplayBufferStatus()
	if err != nil {
		return false, err
	}

	return data.OutputActive, nil
}

func (v5 *obsv5) ToggleInstantReplay() error {
	// There's no such way to Toggle / Untoggle, read input first
	isReplayBufferRunning, err := v5.GetStatusInstantReplay()
	if err != nil {
		return err
	}

	if isReplayBufferRunning {
		_, err := v5.client.Outputs.StopReplayBuffer()
		return err
	}

	_, err = v5.client.Outputs.StartReplayBuffer()
	return err
}

func (v5 *obsv5) SaveInstantReplay() error {
	_, err := v5.client.Outputs.SaveReplayBuffer()
	return err
}
