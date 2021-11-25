package startup

import (
	"os"

	"github.com/canbefree/magazine/utils"
	"github.com/canbefree/magazine/utils/config"
	"github.com/canbefree/magazine/utils/config/toml"
	"github.com/canbefree/magazine/vars"
)

const (
	ENV_MAGAZINE_PATH = "MAGAZINE_PATH"
)

func init() {
	//
	vars.RootPath = utils.GetCurrfilePath()
	envRootPath := os.Getenv(ENV_MAGAZINE_PATH)
	if envRootPath != "" {
		vars.ConfigPath = envRootPath
	} else {
		vars.ConfigPath = vars.RootPath + "config/"
	}

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
