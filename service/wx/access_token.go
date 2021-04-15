package wx

import (
	"encoding/json"
	"fmt"
	"github.com/cellargalaxy/wx-gateway/config"
	"github.com/sirupsen/logrus"
	"net/http"
)

var accessToken string

func GetAccessToken() string {
	return accessToken
}

func flushAccessToken() {
	for i := 0; i < config.Config.Retry; i++ {
		jsonString, err := requestAccessToken()
		if err != nil {
			continue
		}
		token, err := analysisAccessToken(jsonString)
		if err != nil {
			continue
		}
		accessToken = token
		return
	}
}

//获取accessToken
func analysisAccessToken(jsonString string) (string, error) {
	type Response struct {
		ErrCode     int    `json:"errcode"`
		ErrMsg      string `json:"errmsg"`
		AccessToken string `json:"access_token"`
		ExpiresIn   int    `json:"expires_in"`
	}
	var response Response
	err := json.Unmarshal([]byte(jsonString), &response)
	if err != nil {
		logrus.WithFields(logrus.Fields{"err": err, "jsonString": jsonString}).Error("获取accessToken，解析响应异常")
		return "", fmt.Errorf("获取accessToken，解析响应异常")
	}
	if response.ErrCode != 0 {
		logrus.WithFields(logrus.Fields{"err": err, "jsonString": jsonString}).Error("获取accessToken，失败")
		return "", fmt.Errorf("获取accessToken，失败")
	}
	if response.AccessToken == "" {
		logrus.WithFields(logrus.Fields{"err": err, "jsonString": jsonString}).Error("获取accessToken，accessToken为空")
		return "", fmt.Errorf("获取accessToken，accessToken为空")
	}
	logrus.WithFields(logrus.Fields{"accessToken": len(response.AccessToken)}).Info("accessToken长度")
	return response.AccessToken, nil
}

//获取accessToken
func requestAccessToken() (string, error) {
	response, err := httpClient.R().
		SetQueryParam("appid", config.Config.AppId).
		SetQueryParam("secret", config.Config.AppSecret).
		SetQueryParam("grant_type", "client_credential").
		Get("https://api.weixin.qq.com/cgi-bin/token")

	if err != nil {
		logrus.WithFields(logrus.Fields{"err": err}).Error("获取accessToken，请求异常")
		return "", fmt.Errorf("获取accessToken，请求异常")
	}
	if response == nil {
		logrus.WithFields(logrus.Fields{"err": err}).Error("获取accessToken，响应为空")
		return "", fmt.Errorf("获取accessToken，响应为空")
	}
	statusCode := response.StatusCode()
	body := response.String()
	logrus.WithFields(logrus.Fields{"statusCode": statusCode, "body": len(body)}).Info("获取accessToken，响应")
	if statusCode != http.StatusOK {
		logrus.WithFields(logrus.Fields{"StatusCode": statusCode}).Error("获取accessToken，响应码失败")
		return "", fmt.Errorf("获取accessToken，响应码失败: %+v", statusCode)
	}
	return body, nil
}
