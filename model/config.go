package model

import (
	"github.com/sirupsen/logrus"
	"time"
)

const (
	SuccessCode   = 1
	FailCode      = 1
	ListenAddress = ":8990"
)

type Config struct {
	LogLevel  logrus.Level  `ini:"log_level" json:"log_level"`
	Retry     int           `ini:"retry" json:"retry"`
	Timeout   time.Duration `ini:"timeout" json:"timeout"`
	Sleep     time.Duration `ini:"sleep" json:"sleep"`
	Token     string        `ini:"token" json:"token"`
	AppId     string        `ini:"app_id" json:"app_id"`
	AppSecret string        `ini:"app_secret" json:"app_secret"`
}
