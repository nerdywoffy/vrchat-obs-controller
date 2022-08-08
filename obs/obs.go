package obs

type OBSWebsocketAPI interface {
	Start(OBSCredential) error

	GetStatusStream() (bool, error)
	ToggleStream() error

	GetStatusRecord() (bool, error)
	ToggleRecord() error

	GetStatusInstantReplay() (bool, error)
	ToggleInstantReplay() error
	SaveInstantReplay() error

	SetScene(scenes []OBSScene) error
	ToggleToSceneName(sceneName string) error
	ToggleToSceneNumber(number int) error
	GetCurrentSceneNumber() (int, error)

	GetState() (OBSState, error)
}

type OBSCredential struct {
	Host     string
	Port     int64
	Password string
}

type OBSState struct {
	IsRecording    bool
	IsStreaming    bool
	IsReplayBuffer bool
}

type OBSScene struct {
	Number int
	Name   string
}
