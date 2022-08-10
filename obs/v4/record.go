package v4

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
