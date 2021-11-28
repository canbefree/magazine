package mariadb

import (
	"io/ioutil"
	"path"
	"strings"

	"github.com/canbefree/magazine/utils"
)

// GetSqls  Convert files in the same directory to strings
// you can use -d to assign directory
func GetSqlsFiles() (sqls []string) {
	files := utils.ListDirRecurrence(utils.GetCallerCodeDir(1), func(s string) bool {
		if path.Ext(s) == ".sql" {
			return true
		}
		return false
	})
	for _, v := range files {
		sql, err := ioutil.ReadFile(v)
		utils.PanicIfErr(err)
		isql := strings.Split(string(sql), ";")
		for _, v := range isql {
			sqls = append(sqls, strings.Trim(v, " \n"))
		}
	}
	return sqls
}
