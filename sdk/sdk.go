package sdk

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"github.com/cellargalaxy/wx-gateway/model"
	"github.com/go-resty/resty/v2"
	"github.com/sirupsen/logrus"
	"net/http"
	"time"
)

type wxClient struct {
	address    string
	retry      int
	httpClient *resty.Client
}

func NewWxClient(timeout, sleep time.Duration, retry int, address, token string) (*wxClient, error) {
	if address == "" {
		return nil, fmt.Errorf("address为空")
	}
	if token == "" {
		return nil, fmt.Errorf("token为空")
	}
	httpClient := createHttpClient(timeout, sleep, retry, token)
	return &wxClient{address: address, retry: retry, httpClient: httpClient}, nil
}

func createHttpClient(timeout, sleep time.Duration, retry int, token string) *resty.Client {
	httpClient := resty.New().
		SetTimeout(timeout).
		SetRetryCount(retry).
		SetRetryWaitTime(sleep).
		SetRetryMaxWaitTime(5*time.Minute).
		AddRetryCondition(func(response *resty.Response, err error) bool {
			var statusCode int
			if response != nil {
				statusCode = response.StatusCode()
			}
			retry := statusCode != http.StatusOK || err != nil
			if retry {
				logrus.WithFields(logrus.Fields{"statusCode": statusCode, "err": err}).Warn("HTTP请求异常，进行重试")
			}
			return retry
		}).
		SetRetryAfter(func(client *resty.Client, resp *resty.Response) (time.Duration, error) {
			var attempt int
			if resp != nil && resp.Request != nil {
				attempt = resp.Request.Attempt
			}
			if attempt > retry {
				logrus.WithFields(logrus.Fields{"attempt": attempt}).Error("HTTP请求异常，超过最大重试次数")
				return 0, fmt.Errorf("HTTP请求异常，超过最大重试次数")
			}
			duration := sleep
			for i := 0; i < attempt-1; i++ {
				duration *= 10
			}
			logrus.WithFields(logrus.Fields{"attempt": attempt, "duration": duration}).Warn("HTTP请求异常，休眠重试")
			return duration, nil
		}).
		SetHeader(model.TokenHeadKey, token).
		SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	return httpClient
}

func (this wxClient) SendTemplateToTag(templateId string, tagId int, url string, data map[string]interface{}) (bool, error) {
	var jsonString string
	var object bool
	var err error
	for i := 0; i < this.retry; i++ {
		jsonString, err = this.requestSendTemplateToTag(templateId, tagId, url, data)
		if err == nil {
			object, err = this.analysisSendTemplateToTag(jsonString)
			if err == nil {
				return object, err
			}
		}
	}
	return object, err
}

//发送模板信息
func (this wxClient) analysisSendTemplateToTag(jsonString string) (bool, error) {
	type Response struct {
		Code    int                             `json:"code"`
		Message string                          `json:"message"`
		Data    model.SendTemplateToTagResponse `json:"data"`
	}
	var response Response
	err := json.Unmarshal([]byte(jsonString), &response)
	if err != nil {
		logrus.WithFields(logrus.Fields{"err": err, "jsonString": jsonString}).Error("发送模板信息，解析响应异常")
		return false, fmt.Errorf("发送模板信息，解析响应异常")
	}
	if response.Code != 1 {
		logrus.WithFields(logrus.Fields{"jsonString": jsonString}).Error("发送模板信息，失败")
		return false, fmt.Errorf("发送模板信息，失败: %+v", jsonString)
	}
	return true, nil
}

//发送模板信息
func (this wxClient) requestSendTemplateToTag(templateId string, tagId int, url string, data map[string]interface{}) (string, error) {
	response, err := this.httpClient.R().
		SetHeader("Content-Type", "application/json;CHARSET=utf-8").
		SetBody(map[string]interface{}{
			"template_id": templateId,
			"tag_id":      tagId,
			"url":         url,
			"data":        data,
		}).
		Post(this.address + "/api/sendTemplateToTag")

	if err != nil {
		logrus.WithFields(logrus.Fields{"err": err}).Error("发送模板信息，请求异常")
		return "", fmt.Errorf("发送模板信息，请求异常")
	}
	if response == nil {
		logrus.WithFields(logrus.Fields{"err": err}).Error("发送模板信息，响应为空")
		return "", fmt.Errorf("发送模板信息，响应为空")
	}
	statusCode := response.StatusCode()
	body := response.String()
	logrus.WithFields(logrus.Fields{"statusCode": statusCode, "body": body}).Info("发送模板信息，响应")
	if statusCode != http.StatusOK {
		logrus.WithFields(logrus.Fields{"StatusCode": statusCode}).Error("发送模板信息，响应码失败")
		return "", fmt.Errorf("发送模板信息，响应码失败: %+v", statusCode)
	}
	return body, nil
}
