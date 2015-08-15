package api

import (
	"io/ioutil"
	pathlib "path"

	"github.com/buckhx/pathutil"
	"gopkg.in/yaml.v2"
)

type Config struct {
	DenvHome   string
	IgnoreFile string
	InfoFile   string
}

var Settings Config

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func init() {
	// TODO: create a settings lib
	path := "settings.yml"
	if !pathutil.Exists(path) {
		// for tests to reference correct settings
		path = pathlib.Join("..", path)
	}
	config, err := ioutil.ReadFile(path)
	check(err)
	err = yaml.Unmarshal(config, &Settings)
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
}
