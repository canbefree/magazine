package core

// db数据支持： 支持redis,以及 mariadb
type DBIFace interface {
}

//generate db query update
// 缓存数据支持 redis
type CacheIFace interface {
}

// FastDB nosql关系数据库 分布式
type FastDB interface {
}
