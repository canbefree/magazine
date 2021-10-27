package toml

import (
	"cc/tools/utils"
	"cc/vars"
	"io/fs"
	"os"

	"github.com/pelletier/go-toml"
)

const TOML_SUFFIX = ".toml"

type TomlLoader struct {
}

func NewTomlLoader() *TomlLoader {
	return &TomlLoader{}
}

func (l *TomlLoader) SaveConfig(key string, data interface{}) error {
	config, err := toml.Marshal(data)
	if err != nil {
		return err
	}
	configPath := GetConfigPath(key)
	return utils.WriteFile(configPath, config, fs.FileMode(os.O_TRUNC))
}
func (l *TomlLoader) LoadConfig(key string, config interface{}) error {
	configPath := GetConfigPath(key)
	data, err := utils.ReadFile(configPath)
	if err != nil {
		return err
	}
	if err := toml.Unmarshal(data, &config); err != nil {
		return err
	}
	return nil
}

func GetConfigPath(key string) string {
	return vars.ConfigPath + key + TOML_SUFFIX
}
