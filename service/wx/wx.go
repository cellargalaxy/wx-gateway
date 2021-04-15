package wx

import (
	"crypto/tls"
	"github.com/cellargalaxy/wx-gateway/config"
	"github.com/go-resty/resty/v2"
	"time"
)

var httpClient *resty.Client

func init() {
	httpClient = resty.New().
		SetTimeout(config.Config.Timeout).
		SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	flushAccessToken()
	go func() {
		for {
			time.Sleep(30 * time.Minute)
			flushAccessToken()
		}
	}()
}
