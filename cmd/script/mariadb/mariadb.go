package mariadb

var RunSqls []string

var createTableGun = `
test 

`

func init() {
	RunSqls = append(RunSqls, createTableGun)
}
