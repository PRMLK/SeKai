package config

import (
	"github.com/pelletier/go-toml/v2"
	"github.com/sirupsen/logrus"
	"os"
)

type languageConfig struct {
	Language     string `toml:"language"`
	ServerLogger struct {
		LoadConfigMessage         string `toml:"loadConfigMessage"`
		HTTPStartingMessage       string `toml:"httpStartingMessage"`
		HTTPStartingError         string `toml:"httpStartingError"`
		HTTPServerShutdownMessage string `toml:"httpServerShutdownMessage"`
		HTTPServerShutdownError   string `toml:"httpServerShutdownError"`
		HTTPServerExited          string `toml:"httpServerExited"`
		ChunkTemplateLoadedError  string `toml:"chunkTemplateLoadedError"`
		LoadPageError             string `toml:"loadPageError"`
		TomlUnmarshalError        string `toml:"tomlUnmarshalError"`
		ReadFileError             string `toml:"readFileError"`
		SameThemeExist            string `toml:"sameThemeExist"`
		InlineFileReadError       string `toml:"inlineFileReadError"`
		NewPostError              string `toml:"newPostError"`
	} `toml:"serverLogger"`
}

var LanguageConfig languageConfig

func initLanguageConfig() {
	var data []byte
	if tempData, err := os.ReadFile("./configs/language/" + ApplicationConfig.Language + ".toml"); err != nil {
		logrus.Panic(err)
	} else {
		data = tempData
	}
	if err := toml.Unmarshal(data, &LanguageConfig); err != nil {
		logrus.Panic(err)
	}
}
