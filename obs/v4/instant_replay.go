package v4

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
