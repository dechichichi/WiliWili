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
	Kafka   kafka
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

type minio struct {
	Addr        string
	AccessKey   string
	AccessKeyID string
	SecretKey   string
}

type kafka struct {
	Broker          string `yaml:"broker"`
	Topic           string `yaml:"topic"`
	ConsumerGroup   string `yaml:"consumer_group"`
	MaxConnections  int    `yaml:"max_connections"`
	MaxQPS          int    `yaml:"max_qps"`
	AutoOffsetReset string `yaml:"auto_offset_reset"`
}
