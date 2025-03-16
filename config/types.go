package config

type service struct {
	Name     string
	AddrList []string
	LB       bool `mapstructure:"load-balance"`
}
type server struct {
	Secret      string `mapstructure:"private-key"`
	PublicKey   string `mapstructure:"public-key"`
	Version     string
	Name        string
	LogLevel    string `mapstructure:"log-level"`
	IntranetUrl string `mapstructure:"intranet-url"`
}

type config struct {
	Mysql   mysql
	Redis   redis
	Service service
	Etcd    etcd
	Minio   minio
	Server  server
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
type etcd struct {
	Addr string
}

type minio struct{
	Addr string
	AccessKey string 
	SecretKey string 
}