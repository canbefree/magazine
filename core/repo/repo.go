package repo

// quick model genrate
// $host 需替换
// $pwd  需替换
//go:generate gen --sqltype=mysql --connstr root:$pwd@($host:3306)/temp --database temp --json --gorm --out ./example --guregu --rest
