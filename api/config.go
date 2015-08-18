package api

import (
	pathlib "path"

	"github.com/buckhx/pathutil"
	"gopkg.in/yaml.v2"
)

type Config struct {
	DenvHome   string
	IgnoreFile string
	InfoFile   string
	SnapshotDenv	string
}

var Settings Config

//TODO move to a util file
func check(err error) {
	if err != nil {
		panic(err)
	}
}

func init() {
	// settings_yml generated into resources.go
	err := yaml.Unmarshal([]byte(settings_yml), &Settings)
	check(err)
	if len(Settings.DenvHome) < 1 {
		panic("Missing DenvHome setting")
	}
	Settings.DenvHome = pathutil.Expand(Settings.DenvHome)
	if len(Settings.InfoFile) < 1 {
		panic("Missing InfoFile setting")
	}
	Settings.InfoFile = pathlib.Join(Settings.DenvHome, Settings.InfoFile)
	if len(Settings.IgnoreFile) < 1 {
		panic("Missing IgnoreFile setting")
	}
	if len(Settings.SnapshotDenv) < 1 {
		panic("Missing SnapshotDenv setting")
	}
}
