package repo

// quick model genrate
// $host 需替换
// $pwd  需替换
// go:generate gen --sqltype=mysql --connstr root:$pwd@($host:3306)/magazine --database magazine --json --gorm  --generate-dao --generate-proj --dao=repo_dao --model=repo_model --module=github.com/canbefree/magazine/internel/repo
