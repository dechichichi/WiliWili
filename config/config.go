package config

var (
	Mysql *mysql
	Redis *redis
)

func Init(service string) {
	ConfigMaping(service)
}

func ConfigMaping(srv string) {
	c := new(config)
	Mysql = &c.Mysql
	Redis = &c.Redis
}
