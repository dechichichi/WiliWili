package config

import (
	"errors"
	"os"

	"github.com/bytedance/gopkg/util/logger"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
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
	remotePath     = "/config"
	remoteFileName = "config"
	remoteFileType = "yaml"
)

func Init(service string) {
	// 从环境变量中获取 etcd 地址
	etcdAddr := os.Getenv("ETCD_ADDR")
	if etcdAddr == "" {
		logger.Fatalf("config.Init: etcd addr is empty")
	}
	logger.Infof("config.Init: etcd addr: %v", etcdAddr)
	Etcd = &etcd{Addr: etcdAddr}
	// 配置存储在 etcd 中
	err := runtimeViper.AddRemoteProvider(remoteProvider, Etcd.Addr, remotePath)
	if err != nil {
		logger.Fatalf("config.Init: add remote provider error: %v", err)
	}
	runtimeViper.SetConfigName(remoteFileName)
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
