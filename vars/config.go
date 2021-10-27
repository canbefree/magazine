package vars

import (
	"cc/tools/utils/config"
	"cc/tools/utils/config/toml"
)

var ConfigKeyGlobalSetting = "setting"

// 全局配置
type Setting struct {
	DefaultMagazineSize int64 `json:"default_magazine_size,omitempty"` //默认弹夹数量
}

var configMap map[string]interface{}

// LoadConfig 加载所有配置文件
func LoadConfig() {
	configManage := config.NewConfigManage(toml.NewTomlLoader())
	//
	configMap = make(map[string]interface{})

	// 记载全局配置

}

//
func init() {
	LoadConfig()
}
