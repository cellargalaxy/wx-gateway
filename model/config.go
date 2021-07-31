package model

import (
	"github.com/sirupsen/logrus"
	"time"
)

const (
	ListenAddress = ":8990"
)

type Config struct {
	LogLevel      logrus.Level  `ini:"log_level" json:"log_level"`
	Retry         int           `ini:"retry" json:"retry"`
	Timeout       time.Duration `ini:"timeout" json:"timeout"`
	Sleep         time.Duration `ini:"sleep" json:"sleep"`
	ListenAddress string        `json:"listen_address"`
	Secret        string        `ini:"secret" json:"secret"`
	AppId         string        `ini:"app_id" json:"app_id"`
	AppSecret     string        `ini:"app_secret" json:"app_secret"`
}
