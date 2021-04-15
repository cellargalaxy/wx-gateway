package model

import (
	"github.com/sirupsen/logrus"
	"time"
)

type Config struct {
	SuccessCode   int           `json:"success_code"`
	FailCode      int           `json:"fail_code"`
	LogLevel      logrus.Level  `ini:"log_level" json:"log_level"`
	Retry         int           `ini:"retry" json:"retry"`
	Timeout       time.Duration `ini:"timeout" json:"timeout"`
	Sleep         time.Duration `ini:"sleep" json:"sleep"`
	ListenAddress string        `ini:"listen_address" json:"listen_address"`
	Token         string        `ini:"token" json:"token"`
	AppId         string        `ini:"app_id" json:"app_id"`
	AppSecret     string        `ini:"app_secret" json:"app_secret"`
}
