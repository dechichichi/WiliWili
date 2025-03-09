package config

type config struct {
	Mysql mysql
	Redis redis
}
type mysql struct {
	Addr     string
	Database string
	Username string
	Password string
	Charset  string
}

type redis struct {
	Addr     string
	Password string
}
