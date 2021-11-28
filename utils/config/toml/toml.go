package toml

import (
	"io/fs"

	"github.com/canbefree/magazine/utils"
	"github.com/canbefree/magazine/vars"
	"github.com/pelletier/go-toml"
	jww "github.com/spf13/jwalterweatherman"
)

const TOML_SUFFIX = ".toml"

type TomlLoader struct {
}

func NewTomlLoader() *TomlLoader {
	return &TomlLoader{}
}

func (l *TomlLoader) SaveConfig(configType string, configData interface{}) error {
	config, err := toml.Marshal(configData)
	if err != nil {
		return err
	}
	configPath := GetConfigPath(configType)
	return utils.WriteFile(configPath, config, fs.ModePerm)
}
func (l *TomlLoader) LoadConfig(key string, config interface{}) error {
	configPath := GetConfigPath(key)
	data, err := utils.ReadFile(configPath)
	if err != nil {
		return err
	}
	jww.TRACE.Println(&config)
	if err := toml.Unmarshal(data, &config); err != nil {
		return err
	}
	jww.TRACE.Println(&config)
	return nil
}

func GetConfigPath(configType string) string {
	if vars.ConfigPath[len(vars.ConfigPath)-1] != '/' {
		return vars.ConfigPath + "/" + configType + TOML_SUFFIX
	}
	return vars.ConfigPath + configType + TOML_SUFFIX
}
