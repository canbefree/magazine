package config

//
type Example struct {
}

type Configer interface {
}

//  ConfigMange is 配置管理
type ConfigManage struct {
	Loader Loader
}

// NewConfigManage create new manage
func NewConfigManage(loader Loader) *ConfigManage {
	return &ConfigManage{Loader: loader}
}

// Loader is 配置解析,支持 toml 等
type Loader interface {
	SaveConfig(key string, data interface{}) error
	LoadConfig(key string, config interface{}) (interface{}, error)
}
