package main

import (
	"time"

	"github.com/hypebeast/go-osc/osc"
)

const (
	// Replay Buffer
	ReplayBufferToggle  = "/avatar/parameters/OBSReplayBufferToggle"
	ReplayBufferCapture = "/avatar/parameters/OBSReplayBufferCapture"
	ReplayBufferStatus  = "/avatar/parameters/OBSReplayBufferStatus"

	// Recording
	RecordToggle = "/avatar/parameters/OBSRecordToggle"
	RecordStatus = "/avatar/parameters/OBSRecordStatus"

	// Stream
	StreamToggle = "/avatar/parameters/OBSStreamToggle"
	StreamStatus = "/avatar/parameters/OBSStreamStatus"

	// Scene Switch
	SceneSwitchSelector = "/avatar/parameters/OBSSceneSwitchSelector"
)

func pollStatus() {
	poller := time.NewTicker(500 * time.Millisecond)
	for {
		<-poller.C
		runPoller()
	}
}

func runPoller() {
	// Get Stream Status
	isStreaming, err := _obs.GetStatusStream()
	if err != nil {
		_log.Error(err)
		return
	}
	msg := osc.NewMessage(StreamStatus)
	msg.Append(isStreaming)
	_osc.GetClient().Client().Send(msg)

	// Get Record Status
	isRecording, err := _obs.GetStatusRecord()
	if err != nil {
		_log.Error(err)
		return
	}
	msg = osc.NewMessage(RecordStatus)
	msg.Append(isRecording)
	_osc.GetClient().Client().Send(msg)

	// Get Instant Replay Status
	isInstantReplay, err := _obs.GetStatusInstantReplay()
	if err != nil {
		_log.Error(err)
		return
	}
	msg = osc.NewMessage(ReplayBufferStatus)
	msg.Append(isInstantReplay)
	_osc.GetClient().Client().Send(msg)

	// Get Current Scene
	currentSceneNumber, err := _obs.GetCurrentSceneNumber()
	if err != nil {
		_log.Error(err)
	}
	if currentSceneNumber >= 0 {
		msg = osc.NewMessage(SceneSwitchSelector)
		msg.Append(int32(currentSceneNumber))
		_osc.GetClient().Client().Send(msg)
	}
}

func start() {
	// Build Hooks for Replay Buffer
	_osc.GetServer().AddBooleanListener(ReplayBufferToggle, func(b bool) {
		_log.Debugf("ReplayBufferToggle Invoked with value %v", b)
		if !b {
			return
		}

		if err := _obs.ToggleInstantReplay(); err != nil {
			_log.Error(err)
		}
	})
	_osc.GetServer().AddBooleanListener(ReplayBufferCapture, func(b bool) {
		_log.Debugf("ReplayBufferCapture Invoked with value %v", b)
		if !b {
			return
		}

		if err := _obs.SaveInstantReplay(); err != nil {
			_log.Error(err)
		}
	})

	// Build Hooks for Record
	_osc.GetServer().AddBooleanListener(RecordToggle, func(b bool) {
		_log.Debugf("RecordToggle Invoked with value %v", b)
		if !b {
			return
		}

		if err := _obs.ToggleRecord(); err != nil {
			_log.Error(err)
		}
	})

	// Build Hooks for Stream
	_osc.GetServer().AddBooleanListener(StreamToggle, func(b bool) {
		_log.Debugf("StreamToggle Invoked with value %v", b)
		if !b {
			return
		}

		if err := _obs.ToggleStream(); err != nil {
			_log.Error(err)
		}
	})

	// Build Hooks for Scene Selector
	_osc.GetServer().AddIntegerListener(SceneSwitchSelector, func(v int) {
		_log.Debugf("StreamToggle Invoked with value %v", v)

		// If value less than 1, ignore
		if v < 1 {
			return
		}

		if err := _obs.ToggleToSceneNumber(v); err != nil {
			_log.Error(err)
		}
	})

	// Poll?
	go pollStatus()

	// Start Server
	_osc.GetServer().Start()
	<-hold()
}
