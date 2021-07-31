package wx

import (
	"context"
	"crypto/tls"
	"github.com/cellargalaxy/go_common/util"
	"github.com/cellargalaxy/wx-gateway/config"
	"github.com/go-resty/resty/v2"
	"time"
)

var httpClient *resty.Client

func init() {
	cxt := context.Background()
	cxt = util.SetLogId(cxt)
	httpClient = resty.New().
		SetTimeout(config.Config.Timeout).
		SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	flushAccessToken(cxt)
	go func() {
		for {
			time.Sleep(30 * time.Minute)
			flushAccessToken(cxt)
		}
	}()
}
