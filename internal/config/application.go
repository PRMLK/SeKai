package config

import (
	"github.com/pelletier/go-toml/v2"
	"github.com/sirupsen/logrus"
	"os"
)

type applicationConfig struct {
	Version string `toml:"Version"`
	Server  struct {
		Port string `toml:"Port"`
	} `toml:"Server"`
	Database struct {
		Type   string `toml:"type"`
		Sqlite struct {
			Dir string `toml:"dir"`
		} `toml:"sqlite"`
		Mysql struct {
			Server   string `toml:"server"`
			Port     int    `toml:"port"`
			Database string `toml:"database"`
			Username string `toml:"username"`
			Password string `toml:"password"`
		} `toml:"mysql"`
	} `toml:"database"`
}

var ApplicationConfig applicationConfig

func initApplicationConfig() {
	var data []byte
	if tempData, err := os.ReadFile("./configs/application.toml"); err != nil {
		logrus.Panic(err)
	} else {
		data = tempData
	}
	if err := toml.Unmarshal(data, &ApplicationConfig); err != nil {
		logrus.Panic(err)
	}
}
