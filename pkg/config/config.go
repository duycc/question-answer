//===--------------------------- pkg/config/config.go - [question-answer] --------------------------------*- Go -*-===//
// Brief :
//
//
// Author: YongDu
// Date  : 2022-04-24
//===--------------------------------------------------------------------------------------------------------------===//

package config

import (
	"os"
	"path/filepath"
	"runtime"

	"github.com/BurntSushi/toml"
)

var (
	Config  = tomlConfig{}
	ViewDir string
	LogDir  string
)

type tomlConfig struct {
	RootDir string
	Port    string
	DBDebug bool
}

type configByOS struct {
	Windows tomlConfig
	OSX     tomlConfig
	Linux   tomlConfig
}

func init() {
	var configOS configByOS
	if _, err := toml.DecodeFile("config.toml", &configOS); err != nil {
		panic(err)
	}
	if runtime.GOOS == "windows" {
		Config = configOS.Windows
	} else if runtime.GOOS == "darwin" {
		Config = configOS.OSX
	} else {
		Config = configOS.Linux
	}

	ViewDir = filepath.Join(Config.RootDir, "templates")
	LogDir = filepath.Join(Config.RootDir, "logs")
	os.MkdirAll(LogDir, os.ModePerm)
}
