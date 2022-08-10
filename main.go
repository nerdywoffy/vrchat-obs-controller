package main

import (
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/nerdywoffy/vrchat-obs-controller/obs"
	"github.com/nerdywoffy/vrchat-obs-controller/obs/streamlabs"
	v4 "github.com/nerdywoffy/vrchat-obs-controller/obs/v4"
	v5 "github.com/nerdywoffy/vrchat-obs-controller/obs/v5"
	"github.com/nerdywoffy/vrchat-obs-controller/osc"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

var (
	_config *viper.Viper
	_log    *logrus.Logger
	_C      Configuration
	_obs    obs.OBSWebsocketAPI
	_osc    *osc.OSCProtocol
)

func hold() <-chan struct{} {
	wait := make(chan struct{})
	go func() {
		s := make(chan os.Signal, 1)
		signal.Notify(s, syscall.SIGINT, syscall.SIGTERM, syscall.SIGHUP)
		<-s
		close(wait)
	}()
	return wait
}

func main() {
	// Read configuration
	_config = viper.New()
	_config.SetConfigName("config")
	_config.SetConfigType("yaml")
	_config.AddConfigPath("./")
	if err := _config.ReadInConfig(); err != nil {
		panic(err)
	}
	if err := _config.Unmarshal(&_C); err != nil {
		panic(err)
	}

	// Initialize Logger
	_log = logrus.New()
	logrusLevel, err := logrus.ParseLevel(_C.Logs.Level)
	if err != nil {
		_log.Fatal(err)
	}
	_log.SetLevel(logrusLevel)

	// Build OBS
	switch strings.ToLower(_C.OBS.WebsocketVersion) {
	case "v4":
		_obs = v4.New(_log)
	case "v5":
		_obs = v5.New(_log)
	case "streamlabs":
		_obs = streamlabs.New(_log)
	default:
		_log.Fatal("unknown OBS version")
	}
	if err := _obs.Start(_C.OBS.Credential); err != nil {
		_log.Fatal(err)
	}
	if err := _obs.SetScene(_C.OBS.Scenes); err != nil {
		_log.Fatal(err)
	}

	// Build OSC
	_osc = osc.New(_log, _C.OSC)
	start()
}
