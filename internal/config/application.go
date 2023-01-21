package config

import (
	"github.com/mcuadros/go-defaults"
	"os"

	"github.com/pelletier/go-toml/v2"
	"github.com/sirupsen/logrus"
)

type applicationConfig struct {
	Version  string `toml:"version" default:"0.0.1"`
	Language string `toml:"language" default:"zh_cn"`
	HTTP     struct {
		Port int `toml:"port" default:"12070"`
	} `toml:"http"`
	Log struct {
		Dir string `toml:"dir" default:"./logs"`
	} `toml:"log"`
	Database struct {
		Type    string `toml:"type" default:"sqlite3"`
		Sqlite3 struct {
			Dir string `toml:"dir" default:"./sekai.db"`
		} `toml:"sqlite3"`
		Mysql struct {
			Server   string `toml:"server" default:"localhost"`
			Port     int    `toml:"port" default:"3306"`
			Username string `toml:"username"`
			Password string `toml:"password"`
			Db       string `toml:"db"`
		} `toml:"mysql"`
	} `toml:"database"`
	SiteConfig struct {
		SiteRoot            string `toml:"siteRoot" default:"localhost:12070"`
		SiteHome            string `toml:"siteHome" default:"localhost:12070"`
		SiteName            string `toml:"siteName" default:"Sekai"`
		SiteDescription     string `toml:"siteDescription" default:"powered by Golang"`
		SiteBackStageTheme  string `toml:"siteBackStageTheme" default:"Plain-SeKai"`
		SiteFrontStageTheme string `toml:"siteFrontStageTheme" default:"Plain-SeKai"`
	} `toml:"siteConfig"`
	GoogleAuth struct {
		Salt string `toml:"salt"`
	} `toml:"googleAuth"`
}

var ApplicationConfig *applicationConfig

func initApplicationConfig() {
	ApplicationConfig = new(applicationConfig)
	// 读取application.toml
	if data, err := os.ReadFile("./configs/application.toml"); err != nil {
		// 读取错误，尝试写入defaultConfig
		defaults.SetDefaults(ApplicationConfig)
		if tomlData, err := toml.Marshal(ApplicationConfig); err != nil {
			// tomlMarshal错误
			logrus.Panic("tomlMarshal error: " + err.Error())
			return
		} else {
			// 写入错误
			err = os.WriteFile("./configs/application.toml", tomlData, 0775)
			if err != nil {
				logrus.Panic("application.toml write error: " + err.Error())
				return
			}
		}
	} else {
		// 读取正常
		if err := toml.Unmarshal(data, ApplicationConfig); err != nil {
			logrus.Panic(err)
		}
	}
}
