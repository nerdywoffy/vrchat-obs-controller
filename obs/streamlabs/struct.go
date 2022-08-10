package streamlabs

type StreamlabsScene struct {
	ResourceID string `json:"resourceId"`
	ID         string `json:"id"`
	Name       string `json:"name"`
}

type StreamlabsScenesResult []StreamlabsScene

type StreamlabsScenes struct {
	Result StreamlabsScenesResult `json:"result"`
}

type StreamlabsStreamingState struct {
	StreamingStatus        string `json:"streamingStatus"`
	StreamingStatusTime    string `json:"streamingStatusTime"`
	RecordingStatus        string `json:"recordingStatus"`
	RecordingStatusTime    string `json:"recordingStatusTime"`
	ReplayBufferstatus     string `json:"replayBufferStatus"`
	ReplayBufferStatusTime string `json:"replayBufferStatusTime"`
}
