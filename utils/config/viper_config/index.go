package viper_config

import (
	"cc/utils"
	"cc/vars"
	"errors"
	"os"
	"reflect"

	"github.com/pelletier/go-toml"
	"github.com/spf13/jwalterweatherman"
	"github.com/spf13/viper"
)

const VIPER_CONFIG_TYPE = "toml"

var (
	ErrorNotSuppport = errors.New("config data type must be struct")
)

type Loader interface {
	SaveConfig(key string, data interface{}) error
	LoadConfig(key string, config interface{}) error
}

type ViperConfig struct {
}

func NewViperConfig() *ViperConfig {
	return &ViperConfig{}
}

/*
it panits if type's data is not struct
*/
func (l *ViperConfig) SaveConfig(configType string, data interface{}) error {
	t := reflect.TypeOf(data)
	v := reflect.ValueOf(data)
	switch t.Kind() {
	case reflect.Ptr:
		t.Elem().NumField()
	case reflect.Struct:
		for i := 0; i < t.NumField(); i++ {
			viper.Set(t.Field(i).Name, v.Field(i).Interface().(interface{}))
		}
	case reflect.Map:
		data, ok := data.(map[string]interface{})
		if !ok {
			return ErrorNotSuppport
		}
		for k, v := range data {
			viper.Set(k, v)
		}
	default:
		return ErrorNotSuppport
	}
	viper.SetConfigName(configType)
	configPath := GetConfigPath(configType)
	viper.AddConfigPath(configPath)
	viper.SetConfigType(VIPER_CONFIG_TYPE)
	configFileName := configPath + string(os.PathSeparator) + configType + "." + VIPER_CONFIG_TYPE
	if !utils.FileExists(configFileName) {
		utils.WriteFile(configFileName, nil, os.ModeAppend)
	}
	return viper.WriteConfig()
}
func (l *ViperConfig) LoadConfig(key string, config interface{}) error {
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

func GetConfigPath(key string) (path string) {
	if vars.ConfigPath[len(key)-1] == os.PathSeparator {
		path = vars.ConfigPath + key
	} else {
		path = vars.ConfigPath + string(os.PathSeparator) + key
	}
	return
}

func init() {
	jwalterweatherman.SetStdoutThreshold(jwalterweatherman.LevelDebug)
	jwalterweatherman.SetLogThreshold(jwalterweatherman.LevelDebug)
}
