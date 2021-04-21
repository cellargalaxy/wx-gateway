package config

import (
	"fmt"
	"github.com/cellargalaxy/wx-gateway/model"
	"github.com/cellargalaxy/wx-gateway/util"
	"github.com/go-ini/ini"
	"github.com/sirupsen/logrus"
	"time"
)

const (
	configFilePath = "resources/config.ini"
)

var Config = model.Config{
	LogLevel: logrus.InfoLevel,
	Retry:    3,
	Timeout:  3 * time.Second,
	Sleep:    3 * time.Second,
}

func init() {
	exist, _ := util.ExistAndIsFile(configFilePath)
	if exist {
		cfg, err := ini.Load(configFilePath)
		if err != nil {
			panic(err)
		}
		err = cfg.MapTo(&Config)
		if err != nil {
			panic(err)
		}
	}
	checkAndResetConfig()
	if Config.Token == "" {
		panic(fmt.Errorf("token配置为空"))
	}
	if Config.AppId == "" {
		panic(fmt.Errorf("appId配置为空"))
	}
	if Config.AppSecret == "" {
		panic(fmt.Errorf("appSecret配置为空"))
	}
}

func checkAndResetConfig() {
	if Config.LogLevel <= 0 || Config.LogLevel > logrus.TraceLevel {
		Config.LogLevel = logrus.InfoLevel
	}
	if Config.Timeout < 0 {
		Config.Timeout = 3 * time.Second
	}
	if Config.Sleep < 0 {
		Config.Sleep = 3 * time.Second
	}
}
