package startup

import (
	"cc/utils"
	"cc/utils/config"
	"cc/utils/config/toml"
	"cc/vars"
)

func init() {
	vars.RootPath = utils.GetCurrfilePath()
	vars.ConfigPath = vars.RootPath + "config/"

	// 初始化configManage
	tomlLoader := toml.NewTomlLoader()
	vars.ConfigManage = config.NewConfigManage(tomlLoader)

	// TODO  代码应该保留一份默认配置
	for k, v := range vars.ConfigMap {
		vars.ConfigManage.Loader.SaveConfig(k, v)
	}

	for k, v := range vars.ConfigMap {
		vars.ConfigManage.Loader.LoadConfig(k, v)
	}
}
