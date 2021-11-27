package vars

var ConfigKeyGlobalSetting = "setting"

// 全局配置
type Setting struct {
	DefaultMagazineSize int64 `json:"default_magazine_size,omitempty"` //默认弹夹数量
}

type RedisConfig struct {
	Host string `json:"host,omitempty"`
	Port string `json:"port,omitempty"`
	Auth string `json:"auth,omitempty"`
}

type SqliteConfig struct {
	DBPath string `json:"db_path,omitempty"`
}

type MysqlConfig struct {
	DSN string `json:"dsn,omitempty"`
}

var ConfigMap map[string]interface{} = map[string]interface{}{
	"setting": &Setting{
		DefaultMagazineSize: 100,
	},
	"redis_config": &RedisConfig{},
	"sqlite":       &SqliteConfig{},
	"mysql_config": &MysqlConfig{},
}
