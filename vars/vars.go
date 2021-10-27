package vars

import (
	"cc/tools/utils"
	"os"
)

var (
	// 项目根目录
	RootPath string
	//
	VarConstPath string
	// 配置目录
	ConfigPath string
)

func init() {

	VarConstPath = utils.GetCurrfilePath()

	RootPath = utils.GetParentfilePath(VarConstPath) + string(os.PathSeparator)

	ConfigPath = RootPath + "config/"
}
