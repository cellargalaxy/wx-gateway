package util

import (
	"context"
	"errors"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"io"
	"os"
	"time"
)

func InitLog(logPath string) {
	if logPath != "" {
		file, err := os.OpenFile(logPath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0777)
		if err != nil {
			panic(err)
		}
		multiWriter := io.MultiWriter(os.Stdout, file)
		logrus.SetOutput(multiWriter)
	}
	//logrus.SetFormatter(&logrus.JSONFormatter{})
	//logrus.SetReportCaller(true)
}

type GormLog struct {
	ShowSql    bool
	IgnoreErrs []error
}

func (object GormLog) LogMode(logger.LogLevel) logger.Interface {
	return object
}

func (object GormLog) Info(ctx context.Context, s string, args ...interface{}) {
	logrus.WithContext(ctx).Infof(s, args)
}

func (object GormLog) Warn(ctx context.Context, s string, args ...interface{}) {
	logrus.WithContext(ctx).Warnf(s, args)
}

func (object GormLog) Error(ctx context.Context, s string, args ...interface{}) {
	logrus.WithContext(ctx).Errorf(s, args)
}

func (object GormLog) Trace(ctx context.Context, begin time.Time, fc func() (string, int64), err error) {
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		ignore := false
		for i := range object.IgnoreErrs {
			if errors.Is(err, object.IgnoreErrs[i]) {
				ignore = true
				break
			}
		}
		if !ignore {
			elapsed := time.Since(begin)
			sql, _ := fc()
			fields := logrus.Fields{"err": err, "elapsed": elapsed, "sql": sql}
			logrus.WithContext(ctx).WithFields(fields).Error()
			return
		}
	}
	if object.ShowSql {
		elapsed := time.Since(begin)
		sql, _ := fc()
		fields := logrus.Fields{"elapsed": elapsed, "sql": sql}
		logrus.WithContext(ctx).WithFields(fields).Info()
		return
	}
}
