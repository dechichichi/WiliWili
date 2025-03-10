package config

import (
	"os"

	"github.com/bytedance/gopkg/util/logger"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	_ "github.com/spf13/viper/remote"
)

var (
	Service      *service
	Mysql        *mysql
	Redis        *redis
	Etcd         *etcd
	Server       *server
	runtimeViper = viper.New()
)

const (
	remoteProvider = "etcd3"
	remotePath     = "/config" // 修改为实际存储配置的路径
	remoteFileType = "yaml"
)

func Init(service string) {
	etcdAddr := os.Getenv("ETCD_ADDR")
	if etcdAddr == "" {
		logger.Fatalf("config.Init: etcd addr is empty")
	}
	logger.Infof("config.Init: etcd addr: %v", etcdAddr)

	Etcd = &etcd{Addr: etcdAddr}

	err := runtimeViper.AddRemoteProvider(remoteProvider, Etcd.Addr, remotePath)
	if err != nil {
		logger.Fatalf("config.Init: add remote provider error: %v", err)
	}

	runtimeViper.SetConfigType(remoteFileType) // 确保配置类型为 YAML

	if err := runtimeViper.ReadRemoteConfig(); err != nil {
		logger.Fatalf("config.Init: read config error: %v", err)
	}

	logger.Infof("Loaded config: %+v", runtimeViper.AllSettings()) // 打印加载的配置

	configMapping(service)

	runtimeViper.OnConfigChange(func(e fsnotify.Event) {
		logger.Infof("config: notice config changed: %v", e.String())
		configMapping(service)
	})
	runtimeViper.WatchConfig()
}

func configMapping(srv string) {
	c := new(config)
	if err := runtimeViper.Unmarshal(&c); err != nil {
		// 由于这个函数会在配置重载时被再次触发，所以需要判断日志记录方式
		logger.Fatalf("config.configMapping: config: unmarshal error: %v", err)
	}
	Mysql = &c.Mysql
	Redis = &c.Redis
	Service = &c.Service
	Server = &c.Server
	Service = getService(srv)
}

func getService(name string) *service {
	addrList := runtimeViper.GetStringSlice("services." + name + ".addr")
	return &service{
		Name:     runtimeViper.GetString("services." + name + ".name"),
		AddrList: addrList,
		LB:       runtimeViper.GetBool("services." + name + ".load-balance"),
	}
}
