package mariadb

import (
	"cc/utils"
	"io/ioutil"
	"strings"

	"gorm.io/gorm"
)

var RunSqls []string

var createTableGun = `
test 

`

func init() {
	RunSqls = append(RunSqls, createTableGun)
}

func RunSQLScript(db *gorm.DB, script string, skipDrop bool) {
	content, err := ioutil.ReadFile(script)
	utils.FatalIfError(err)
	sqls := strings.Split(string(content), ";")
	for _, sql := range sqls {
		s := strings.TrimSpace(sql)
		if s == "" || (skipDrop && strings.Contains(s, "drop")) {
			continue
		}
		db.Exec(s)
	}
}
