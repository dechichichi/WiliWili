package config

import (
	"errors"
	"os"

	"github.com/bytedance/gopkg/util/logger"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

var (
	Server       *server
	Service      *service
	MySQL        *mySQL
	Redis        *redis
	Etcd         *etcd
	runtimeViper = viper.New()
)

const (
	remoteProvider = "etcd3"
	remotePath     = "/config"
	remoteFileName = "config"
	remoteFileType = "yaml"
)

func Init(service string) {
	etcdAddr := os.Getenv("ETCD_ADDR")
	if etcdAddr == "" {
		logger.Error("config.Init: etcd addr is empty")
	}
	logger.Infof("config.Init: etcd addr is %s", etcdAddr)
	Etcd = &etcd{Addr: etcdAddr}
	err := runtimeViper.AddRemoteProvider(remoteProvider, Etcd.Addr, remotePath)
	if err != nil {
		logger.Error("config.Init: add remote provider failed", err)
	}
	runtimeViper.SetConfigType(remoteFileType)
	runtimeViper.SetConfigType(remoteFileType)
	if err := runtimeViper.ReadRemoteConfig(); err != nil {
		var configFileNotFoundError viper.ConfigFileNotFoundError
		if errors.As(err, &configFileNotFoundError) {
			logger.Fatal("config.Init: could not find config files")
		}
		logger.Fatalf("config.Init: read config error: %v", err)
	}
	configMapping(service)
	// 设置持续监听
	runtimeViper.OnConfigChange(func(e fsnotify.Event) {
		// 我们无法确定监听到配置变更时是否已经初始化完毕，所以此处需要做一个判断
		logger.Infof("config: notice config changed: %v\n", e.String())
		configMapping(service) // 重新映射配置
	})
	runtimeViper.WatchConfig()
}

func configMapping(srv string) {
	c := new(config)
	if err := runtimeViper.Unmarshal(c); err != nil {
		logger.Error("config.configMapping: unmarshal config error", err)
	}
	Server = &c.Server
	Service = &c.Service
	MySQL = &c.MySQL
	Redis = &c.Redis
	Etcd = &c.Etcd
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
