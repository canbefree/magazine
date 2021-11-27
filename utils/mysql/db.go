package mysql

import (
	"github.com/canbefree/magazine/utils"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Conect_test() {
	_, err := gorm.Open(mysql.Open("root:123456@(host:3366)/temp"))
	utils.PanicIfErr(err)
}
