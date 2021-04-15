package wx

import (
	"encoding/json"
	"fmt"
	"github.com/cellargalaxy/wx-gateway/config"
	"github.com/cellargalaxy/wx-gateway/model"
	"github.com/sirupsen/logrus"
	"net/http"
)

//获取全部用户信息
func ListAllUserInfo() ([]model.UserInfo, error) {
	openIds, err := ListAllOpenId()
	if err != nil {
		return nil, err
	}
	return ListUserInfo(openIds)
}

//获取全部openId
func ListAllOpenId() ([]string, error) {
	var jsonString string
	var object []string
	var err error
	for i := 0; i < config.Config.Retry; i++ {
		jsonString, err = requestListAllOpenId()
		if err == nil {
			object, err = analysisListAllOpenId(jsonString)
			if err == nil {
				return object, err
			}
		}
		flushAccessToken()
	}
	return object, err
}

//获取全部openId
func analysisListAllOpenId(jsonString string) ([]string, error) {
	type Data struct {
		OpenIds []string `json:"openid"`
	}
	type Response struct {
		ErrCode int    `json:"errcode"`
		ErrMsg  string `json:"errmsg"`
		Data    Data   `json:"data"`
	}
	var response Response
	err := json.Unmarshal([]byte(jsonString), &response)
	if err != nil {
		logrus.WithFields(logrus.Fields{"err": err, "jsonString": jsonString}).Error("获取全部openId，解析响应异常")
		return nil, fmt.Errorf("获取全部openId，解析响应异常")
	}
	if response.ErrCode != 0 {
		logrus.WithFields(logrus.Fields{"jsonString": jsonString}).Error("获取全部openId，失败")
		return nil, fmt.Errorf("获取全部openId，失败")
	}
	return response.Data.OpenIds, nil
}

//获取全部openId
func requestListAllOpenId() (string, error) {
	response, err := httpClient.R().
		SetQueryParam("access_token", GetAccessToken()).
		Get("https://api.weixin.qq.com/cgi-bin/user/get")

	if err != nil {
		logrus.WithFields(logrus.Fields{"err": err}).Error("获取全部openId，请求异常")
		return "", fmt.Errorf("获取全部openId，请求异常")
	}
	if response == nil {
		logrus.WithFields(logrus.Fields{"err": err}).Error("获取全部openId，响应为空")
		return "", fmt.Errorf("获取全部openId，响应为空")
	}
	statusCode := response.StatusCode()
	body := response.String()
	logrus.WithFields(logrus.Fields{"statusCode": statusCode, "body": len(body)}).Info("获取全部openId，响应")
	if statusCode != http.StatusOK {
		logrus.WithFields(logrus.Fields{"StatusCode": statusCode}).Error("获取全部openId，响应码失败")
		return "", fmt.Errorf("获取全部openId，响应码失败: %+v", statusCode)
	}
	return body, nil
}

//获取用户信息
func ListUserInfo(openIds []string) ([]model.UserInfo, error) {
	var jsonString string
	var object []model.UserInfo
	var err error
	for i := 0; i < config.Config.Retry; i++ {
		jsonString, err = requestListUserInfo(openIds)
		if err == nil {
			object, err = analysisListUserInfo(jsonString)
			if err == nil {
				return object, err
			}
		}
		flushAccessToken()
	}
	return object, err
}

//获取用户信息
func analysisListUserInfo(jsonString string) ([]model.UserInfo, error) {
	type Response struct {
		ErrCode      int              `json:"errcode"`
		ErrMsg       string           `json:"errmsg"`
		UserInfoList []model.UserInfo `json:"user_info_list"`
	}
	var response Response
	err := json.Unmarshal([]byte(jsonString), &response)
	if err != nil {
		logrus.WithFields(logrus.Fields{"err": err, "jsonString": jsonString}).Error("获取用户信息，解析响应异常")
		return nil, fmt.Errorf("获取用户信息，解析响应异常")
	}
	if response.ErrCode != 0 {
		logrus.WithFields(logrus.Fields{"jsonString": jsonString}).Error("获取用户信息，失败")
		return nil, fmt.Errorf("获取用户信息，失败")
	}
	return response.UserInfoList, nil
}

//获取用户信息
func requestListUserInfo(openIds []string) (string, error) {
	var userList []map[string]interface{}
	for i := range openIds {
		userList = append(userList, map[string]interface{}{"openid": openIds[i], "lang": "zh_CN"})
	}
	logrus.WithFields(logrus.Fields{"userList": userList}).Info("获取用户信息")

	response, err := httpClient.R().
		SetHeader("Content-Type", "application/json;CHARSET=utf-8").
		SetQueryParam("access_token", GetAccessToken()).
		SetBody(map[string]interface{}{
			"user_list": userList,
		}).
		Post("https://api.weixin.qq.com/cgi-bin/user/info/batchget")

	if err != nil {
		logrus.WithFields(logrus.Fields{"err": err}).Error("获取用户信息，请求异常")
		return "", fmt.Errorf("获取用户信息，请求异常")
	}
	if response == nil {
		logrus.WithFields(logrus.Fields{"err": err}).Error("获取用户信息，响应为空")
		return "", fmt.Errorf("获取用户信息，响应为空")
	}
	statusCode := response.StatusCode()
	body := response.String()
	logrus.WithFields(logrus.Fields{"statusCode": statusCode, "body": len(body)}).Info("获取用户信息，响应")
	if statusCode != http.StatusOK {
		logrus.WithFields(logrus.Fields{"StatusCode": statusCode}).Error("获取用户信息，响应码失败")
		return "", fmt.Errorf("获取用户信息，响应码失败: %+v", statusCode)
	}
	return body, nil
}

//获取标签下的openId
func ListOpenIdByTagId(tagId int) ([]string, error) {
	var jsonString string
	var object []string
	var err error
	for i := 0; i < config.Config.Retry; i++ {
		jsonString, err = requestListOpenIdByTagId(tagId)
		if err == nil {
			object, err = analysisListOpenIdByTagId(jsonString)
			if err == nil {
				return object, err
			}
		}
		flushAccessToken()
	}
	return object, err
}

//获取标签下的openId
func analysisListOpenIdByTagId(jsonString string) ([]string, error) {
	type Data struct {
		OpenIds []string `json:"openid"`
	}
	type Response struct {
		ErrCode int    `json:"errcode"`
		ErrMsg  string `json:"errmsg"`
		Data    Data   `json:"data"`
	}
	var response Response
	err := json.Unmarshal([]byte(jsonString), &response)
	if err != nil {
		logrus.WithFields(logrus.Fields{"err": err, "jsonString": jsonString}).Error("获取标签下的openId，解析响应异常")
		return nil, fmt.Errorf("获取标签下的openId，解析响应异常")
	}
	if response.ErrCode != 0 {
		logrus.WithFields(logrus.Fields{"jsonString": jsonString}).Error("获取标签下的openId，失败")
		return nil, fmt.Errorf("获取标签下的openId，失败")
	}
	return response.Data.OpenIds, nil
}

//获取标签下的openId
func requestListOpenIdByTagId(tagId int) (string, error) {
	response, err := httpClient.R().
		SetHeader("Content-Type", "application/json;CHARSET=utf-8").
		SetQueryParam("access_token", GetAccessToken()).
		SetBody(map[string]interface{}{
			"tagid":       tagId,
			"next_openid": "",
		}).
		Post("https://api.weixin.qq.com/cgi-bin/user/tag/get")

	if err != nil {
		logrus.WithFields(logrus.Fields{"err": err}).Error("获取标签下的openId，请求异常")
		return "", fmt.Errorf("获取标签下的openId，请求异常")
	}
	if response == nil {
		logrus.WithFields(logrus.Fields{"err": err}).Error("获取标签下的openId，响应为空")
		return "", fmt.Errorf("获取标签下的openId，响应为空")
	}
	statusCode := response.StatusCode()
	body := response.String()
	logrus.WithFields(logrus.Fields{"statusCode": statusCode, "body": len(body)}).Info("获取标签下的openId，响应")
	if statusCode != http.StatusOK {
		logrus.WithFields(logrus.Fields{"StatusCode": statusCode}).Error("获取标签下的openId，响应码失败")
		return "", fmt.Errorf("获取标签下的openId，响应码失败: %+v", statusCode)
	}
	return body, nil
}
