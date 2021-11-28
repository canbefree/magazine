package startup

import (
	"log"
	"os"
	"path"

	"github.com/canbefree/magazine/utils"
	"github.com/canbefree/magazine/utils/config"
	"github.com/canbefree/magazine/utils/config/toml"
	"github.com/canbefree/magazine/vars"
	jww "github.com/spf13/jwalterweatherman"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

const (
	ENV_MAGAZINE_PATH = "MAGAZINE_PATH"
)

func init() {
	jww.SetStdoutThreshold(jww.LevelTrace)
	jww.SetFlags(log.Llongfile)
	envRootPath := os.Getenv(ENV_MAGAZINE_PATH)
	if envRootPath != "" {
		vars.RootPath = envRootPath
	} else {
		vars.RootPath = utils.GetCurrfilePath()
	}

	// TODO 目前为了方便调试，将目录设置为编译目录 ./config
	vars.ConfigPath = path.Dir(utils.GetCallerCodeDir(1)) + "/.config"
	jww.TRACE.Printf("vars.ConfigPath init :%v", vars.ConfigPath)
	// 初始化configManage
	tomlLoader := toml.NewTomlLoader()
	vars.ConfigManage = config.NewConfigManage(tomlLoader)
	jww.TRACE.Printf("%v", vars.ConfigMap)
	for k, v := range vars.ConfigMap {
		vars.ConfigManage.Loader.LoadConfig(k, v)
	}

	// 初始化DB
	var err error
	dsn := vars.ConfigMap["mysql_config"].(*vars.MysqlConfigSetting).DSN
	vars.DB, err = gorm.Open(mysql.Open(dsn))
	utils.PanicIfErr(err)
}
