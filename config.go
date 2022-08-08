package main

import (
	"github.com/nerdywoffy/vrchat-obs-controller/obs"
	"github.com/nerdywoffy/vrchat-obs-controller/osc"
)

type Configuration struct {
	OSC  osc.OscSetting
	OBS  OBSConfiguration
	Logs LogsConfiguration
}

type OBSConfiguration struct {
	WebsocketVersion string
	Credential       obs.OBSCredential
	Scenes           []obs.OBSScene
}

type LogsConfiguration struct {
	Level string
}
