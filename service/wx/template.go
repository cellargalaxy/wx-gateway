package wx

import (
	"encoding/json"
	"fmt"
	"github.com/cellargalaxy/wx-gateway/config"
	"github.com/cellargalaxy/wx-gateway/model"
	"github.com/sirupsen/logrus"
	"net/http"
)

//给标签用户发送模板消息
func SendTemplateToTag(templateId string, tagId int, url string, dataMap map[string]interface{}) ([]string, error) {
	data := map[string]model.TemplateData{}
	for key, value := range dataMap {
		data[key] = model.TemplateData{Value: value}
	}
	logrus.WithFields(logrus.Fields{"data": data}).Info("给标签用户发送模板消息")
	openIds, err := ListOpenIdByTagId(tagId)
	if err != nil {
		return nil, err
	}
	var failOpenIds []string
	for i := range openIds {
		success, _ := SendTemplate(openIds[i], templateId, url, data)
		if !success {
			failOpenIds = append(failOpenIds, openIds[i])
		}
	}
	return failOpenIds, nil
}

//获取全部模板
func ListAllTemplate() ([]model.Template, error) {
	var jsonString string
	var object []model.Template
	var err error
	for i := 0; i < config.Config.Retry; i++ {
		jsonString, err = requestListAllTemplate()
		if err == nil {
			object, err = analysisListAllTemplate(jsonString)
			if err == nil {
				return object, err
			}
		}
		flushAccessToken()
	}
	return object, err
}

//获取所有模板
func analysisListAllTemplate(jsonString string) ([]model.Template, error) {
	type Response struct {
		ErrCode      int              `json:"errcode"`
		ErrMsg       string           `json:"errmsg"`
		TemplateList []model.Template `json:"template_list"`
	}
	var response Response
	err := json.Unmarshal([]byte(jsonString), &response)
	if err != nil {
		logrus.WithFields(logrus.Fields{"err": err, "jsonString": jsonString}).Error("获取所有模板，解析响应异常")
		return nil, fmt.Errorf("获取所有模板，解析响应异常")
	}
	if response.ErrCode != 0 {
		logrus.WithFields(logrus.Fields{"jsonString": jsonString}).Error("获取所有模板，失败")
		return nil, fmt.Errorf("获取所有模板，失败")
	}
	return response.TemplateList, nil
}

//获取所有模板
func requestListAllTemplate() (string, error) {
	response, err := httpClient.R().
		SetQueryParam("access_token", GetAccessToken()).
		Get("https://api.weixin.qq.com/cgi-bin/template/get_all_private_template")

	if err != nil {
		logrus.WithFields(logrus.Fields{"err": err}).Error("获取所有模板，请求异常")
		return "", fmt.Errorf("获取所有模板，请求异常")
	}
	if response == nil {
		logrus.WithFields(logrus.Fields{"err": err}).Error("获取所有模板，响应为空")
		return "", fmt.Errorf("获取所有模板，响应为空")
	}
	statusCode := response.StatusCode()
	body := response.String()
	logrus.WithFields(logrus.Fields{"statusCode": statusCode, "body": len(body)}).Info("获取所有模板，响应")
	if statusCode != http.StatusOK {
		logrus.WithFields(logrus.Fields{"StatusCode": statusCode}).Error("获取所有模板，响应码失败")
		return "", fmt.Errorf("获取所有模板，响应码失败: %+v", statusCode)
	}
	return body, nil
}

//发送模板信息
func SendTemplate(openId string, templateId string, url string, data map[string]model.TemplateData) (bool, error) {
	var jsonString string
	var object bool
	var err error
	for i := 0; i < config.Config.Retry; i++ {
		jsonString, err = requestSendTemplate(openId, templateId, url, data)
		if err == nil {
			object, err = analysisSendTemplate(jsonString)
			if err == nil {
				return object, err
			}
		}
		flushAccessToken()
	}
	return object, err
}

//发送模板信息
func analysisSendTemplate(jsonString string) (bool, error) {
	type Response struct {
		ErrCode int    `json:"errcode"`
		ErrMsg  string `json:"errmsg"`
	}
	var response Response
	err := json.Unmarshal([]byte(jsonString), &response)
	if err != nil {
		logrus.WithFields(logrus.Fields{"err": err, "jsonString": jsonString}).Error("发送模板信息，解析响应异常")
		return false, fmt.Errorf("发送模板信息，解析响应异常")
	}
	if response.ErrCode != 0 {
		logrus.WithFields(logrus.Fields{"jsonString": jsonString}).Error("发送模板信息，失败")
		return false, fmt.Errorf("发送模板信息，失败")
	}
	return true, nil
}

//发送模板信息
func requestSendTemplate(openId string, templateId string, url string, data map[string]model.TemplateData) (string, error) {
	response, err := httpClient.R().
		SetHeader("Content-Type", "application/json;CHARSET=utf-8").
		SetQueryParam("access_token", GetAccessToken()).
		SetBody(map[string]interface{}{
			"touser":      openId,
			"template_id": templateId,
			"url":         url,
			"data":        data,
		}).
		Post("https://api.weixin.qq.com/cgi-bin/message/template/send")

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
	logrus.WithFields(logrus.Fields{"statusCode": statusCode, "body": len(body)}).Info("发送模板信息，响应")
	if statusCode != http.StatusOK {
		logrus.WithFields(logrus.Fields{"StatusCode": statusCode}).Error("发送模板信息，响应码失败")
		return "", fmt.Errorf("发送模板信息，响应码失败: %+v", statusCode)
	}
	return body, nil
}
