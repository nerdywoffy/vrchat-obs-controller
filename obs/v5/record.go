package v5

func (v5 *obsv5) GetStatusRecord() (bool, error) {
	data, err := v5.client.Record.GetRecordStatus()
	if err != nil {
		return false, err
	}

	return data.OutputActive, nil
}

func (v5 *obsv5) ToggleRecord() error {
	_, err := v5.client.Record.ToggleRecord()
	return err
}
