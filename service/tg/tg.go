package tg

import (
	"context"
	"crypto/tls"
	"github.com/cellargalaxy/wx-gateway/config"
	"github.com/go-resty/resty/v2"
)

var httpClient *resty.Client

func init() {
	httpClient = resty.New().
		SetTimeout(config.Config.Timeout).
		SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
}

//给配置chatId发送tg信息
func SendTgMsg2ConfigChatId(ctx context.Context, text string) (bool, error) {
	return SendMsg(ctx, config.Config.TgChatId, text)
}
