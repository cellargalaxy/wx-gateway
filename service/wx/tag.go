package wx

import (
	"encoding/json"
	"fmt"
	"github.com/cellargalaxy/wx-gateway/config"
	"github.com/cellargalaxy/wx-gateway/model"
	"github.com/sirupsen/logrus"
	"net/http"
)

//为用户删标签
func DeleteTagFromUser(tagId int, openIds []string) (bool, error) {
	var jsonString string
	var object bool
	var err error
	for i := 0; i < config.Config.Retry; i++ {
		jsonString, err = requestDeleteTagFromUser(tagId, openIds)
		if err == nil {
			object, err = analysisDeleteTagFromUser(jsonString)
			if err == nil {
				return object, err
			}
		}
		flushAccessToken()
	}
	return object, err
}

//为用户删标签
func analysisDeleteTagFromUser(jsonString string) (bool, error) {
	type Response struct {
		ErrCode int    `json:"errcode"`
		ErrMsg  string `json:"errmsg"`
	}
	var response Response
	err := json.Unmarshal([]byte(jsonString), &response)
	if err != nil {
		logrus.WithFields(logrus.Fields{"err": err, "jsonString": jsonString}).Error("为用户删标签，解析响应异常")
		return false, fmt.Errorf("为用户删标签，解析响应异常")
	}
	if response.ErrCode != 0 {
		logrus.WithFields(logrus.Fields{"jsonString": jsonString}).Error("为用户删标签，失败")
		return false, fmt.Errorf("为用户删标签，失败")
	}
	return true, nil
}

//为用户删标签
func requestDeleteTagFromUser(tagId int, openIds []string) (string, error) {
	response, err := httpClient.R().
		SetHeader("Content-Type", "application/json;CHARSET=utf-8").
		SetQueryParam("access_token", GetAccessToken()).
		SetBody(map[string]interface{}{
			"tagid":       tagId,
			"openid_list": openIds,
		}).
		Post("https://api.weixin.qq.com/cgi-bin/tags/members/batchuntagging")

	if err != nil {
		logrus.WithFields(logrus.Fields{"err": err}).Error("为用户删标签，请求异常")
		return "", fmt.Errorf("为用户删标签，请求异常")
	}
	if response == nil {
		logrus.WithFields(logrus.Fields{"err": err}).Error("为用户删标签，响应为空")
		return "", fmt.Errorf("为用户删标签，响应为空")
	}
	statusCode := response.StatusCode()
	body := response.String()
	logrus.WithFields(logrus.Fields{"statusCode": statusCode, "body": len(body)}).Info("为用户删标签，响应")
	if statusCode != http.StatusOK {
		logrus.WithFields(logrus.Fields{"StatusCode": statusCode}).Error("为用户删标签，响应码失败")
		return "", fmt.Errorf("为用户删标签，响应码失败: %+v", statusCode)
	}
	return body, nil
}

//为用户加标签
func AddTagToUser(tagId int, openIds []string) (bool, error) {
	var jsonString string
	var object bool
	var err error
	for i := 0; i < config.Config.Retry; i++ {
		jsonString, err = requestAddTagToUser(tagId, openIds)
		if err == nil {
			object, err = analysisAddTagToUser(jsonString)
			if err == nil {
				return object, err
			}
		}
		flushAccessToken()
	}
	return object, err
}

//给用户加标签
func analysisAddTagToUser(jsonString string) (bool, error) {
	type Response struct {
		ErrCode int    `json:"errcode"`
		ErrMsg  string `json:"errmsg"`
	}
	var response Response
	err := json.Unmarshal([]byte(jsonString), &response)
	if err != nil {
		logrus.WithFields(logrus.Fields{"err": err, "jsonString": jsonString}).Error("给用户加标签，解析响应异常")
		return false, fmt.Errorf("给用户加标签，解析响应异常")
	}
	if response.ErrCode != 0 {
		logrus.WithFields(logrus.Fields{"jsonString": jsonString}).Error("给用户加标签，失败")
		return false, fmt.Errorf("给用户加标签，失败")
	}
	return true, nil
}

//给用户加标签
func requestAddTagToUser(tagId int, openIds []string) (string, error) {
	response, err := httpClient.R().
		SetHeader("Content-Type", "application/json;CHARSET=utf-8").
		SetQueryParam("access_token", GetAccessToken()).
		SetBody(map[string]interface{}{
			"tagid":       tagId,
			"openid_list": openIds,
		}).
		Post("https://api.weixin.qq.com/cgi-bin/tags/members/batchtagging")

	if err != nil {
		logrus.WithFields(logrus.Fields{"err": err}).Error("给用户加标签，请求异常")
		return "", fmt.Errorf("给用户加标签，请求异常")
	}
	if response == nil {
		logrus.WithFields(logrus.Fields{"err": err}).Error("给用户加标签，响应为空")
		return "", fmt.Errorf("给用户加标签，响应为空")
	}
	statusCode := response.StatusCode()
	body := response.String()
	logrus.WithFields(logrus.Fields{"statusCode": statusCode, "body": len(body)}).Info("给用户加标签，响应")
	if statusCode != http.StatusOK {
		logrus.WithFields(logrus.Fields{"StatusCode": statusCode}).Error("给用户加标签，响应码失败")
		return "", fmt.Errorf("给用户加标签，响应码失败: %+v", statusCode)
	}
	return body, nil
}

//删除标签
func DeleteTag(tagId int) (bool, error) {
	var jsonString string
	var object bool
	var err error
	for i := 0; i < config.Config.Retry; i++ {
		jsonString, err = requestDeleteTag(tagId)
		if err == nil {
			object, err = analysisDeleteTag(jsonString)
			if err == nil {
				return object, err
			}
		}
		flushAccessToken()
	}
	return object, err
}

//删除标签
func analysisDeleteTag(jsonString string) (bool, error) {
	type Response struct {
		ErrCode int    `json:"errcode"`
		ErrMsg  string `json:"errmsg"`
	}
	var response Response
	err := json.Unmarshal([]byte(jsonString), &response)
	if err != nil {
		logrus.WithFields(logrus.Fields{"err": err, "jsonString": jsonString}).Error("删除标签，解析响应异常")
		return false, fmt.Errorf("删除标签，解析响应异常")
	}
	if response.ErrCode != 0 {
		logrus.WithFields(logrus.Fields{"jsonString": jsonString}).Error("删除标签，失败")
		return false, fmt.Errorf("删除标签，失败")
	}
	return true, nil
}

//删除标签
func requestDeleteTag(tagId int) (string, error) {
	response, err := httpClient.R().
		SetHeader("Content-Type", "application/json;CHARSET=utf-8").
		SetQueryParam("access_token", GetAccessToken()).
		SetBody(map[string]interface{}{
			"tag": map[string]interface{}{"id": tagId},
		}).
		Post("https://api.weixin.qq.com/cgi-bin/tags/delete")

	if err != nil {
		logrus.WithFields(logrus.Fields{"err": err}).Error("删除标签，请求异常")
		return "", fmt.Errorf("删除标签，请求异常")
	}
	if response == nil {
		logrus.WithFields(logrus.Fields{"err": err}).Error("删除标签，响应为空")
		return "", fmt.Errorf("删除标签，响应为空")
	}
	statusCode := response.StatusCode()
	body := response.String()
	logrus.WithFields(logrus.Fields{"statusCode": statusCode, "body": len(body)}).Info("删除标签，响应")
	if statusCode != http.StatusOK {
		logrus.WithFields(logrus.Fields{"StatusCode": statusCode}).Error("删除标签，响应码失败")
		return "", fmt.Errorf("删除标签，响应码失败: %+v", statusCode)
	}
	return body, nil
}

//获取所有标签
func ListAllTag() ([]model.Tag, error) {
	var jsonString string
	var object []model.Tag
	var err error
	for i := 0; i < config.Config.Retry; i++ {
		jsonString, err = requestListAllTag()
		if err == nil {
			object, err = analysisListAllTag(jsonString)
			if err == nil {
				return object, err
			}
		}
		flushAccessToken()
	}
	return object, err
}

//获取所有标签
func analysisListAllTag(jsonString string) ([]model.Tag, error) {
	type Response struct {
		ErrCode int         `json:"errcode"`
		ErrMsg  string      `json:"errmsg"`
		Tags    []model.Tag `json:"tags"`
	}
	var response Response
	err := json.Unmarshal([]byte(jsonString), &response)
	if err != nil {
		logrus.WithFields(logrus.Fields{"err": err, "jsonString": jsonString}).Error("获取所有标签，解析响应异常")
		return nil, fmt.Errorf("获取所有标签，解析响应异常")
	}
	if response.ErrCode != 0 {
		logrus.WithFields(logrus.Fields{"jsonString": jsonString}).Error("获取所有标签，失败")
		return nil, fmt.Errorf("获取所有标签，失败")
	}
	return response.Tags, nil
}

//获取所有标签
func requestListAllTag() (string, error) {
	response, err := httpClient.R().
		SetQueryParam("access_token", GetAccessToken()).
		Get("https://api.weixin.qq.com/cgi-bin/tags/get")

	if err != nil {
		logrus.WithFields(logrus.Fields{"err": err}).Error("获取所有标签，请求异常")
		return "", fmt.Errorf("获取所有标签，请求异常")
	}
	if response == nil {
		logrus.WithFields(logrus.Fields{"err": err}).Error("获取所有标签，响应为空")
		return "", fmt.Errorf("获取所有标签，响应为空")
	}
	statusCode := response.StatusCode()
	body := response.String()
	logrus.WithFields(logrus.Fields{"statusCode": statusCode, "body": len(body)}).Info("获取所有标签，响应")
	if statusCode != http.StatusOK {
		logrus.WithFields(logrus.Fields{"StatusCode": statusCode}).Error("获取所有标签，响应码失败")
		return "", fmt.Errorf("获取所有标签，响应码失败: %+v", statusCode)
	}
	return body, nil
}

//创建标签
func CreateTag(tag string) (bool, error) {
	var jsonString string
	var object bool
	var err error
	for i := 0; i < config.Config.Retry; i++ {
		jsonString, err = requestCreateTag(tag)
		if err == nil {
			object, err = analysisCreateTag(jsonString)
			if err == nil {
				return object, err
			}
		}
		flushAccessToken()
	}
	return object, err
}

//创建标签
func analysisCreateTag(jsonString string) (bool, error) {
	type Response struct {
		ErrCode int    `json:"errcode"`
		ErrMsg  string `json:"errmsg"`
	}
	var response Response
	err := json.Unmarshal([]byte(jsonString), &response)
	if err != nil {
		logrus.WithFields(logrus.Fields{"err": err, "jsonString": jsonString}).Error("创建标签，解析响应异常")
		return false, fmt.Errorf("创建标签，解析响应异常")
	}
	if response.ErrCode != 0 {
		logrus.WithFields(logrus.Fields{"jsonString": jsonString}).Error("创建标签，失败")
		return false, fmt.Errorf("创建标签，失败")
	}
	return true, nil
}

//创建标签
func requestCreateTag(tag string) (string, error) {
	response, err := httpClient.R().
		SetHeader("Content-Type", "application/json;CHARSET=utf-8").
		SetQueryParam("access_token", GetAccessToken()).
		SetBody(map[string]interface{}{
			"tag": map[string]string{"name": tag},
		}).
		Post("https://api.weixin.qq.com/cgi-bin/tags/create")

	if err != nil {
		logrus.WithFields(logrus.Fields{"err": err}).Error("创建标签，请求异常")
		return "", fmt.Errorf("创建标签，请求异常")
	}
	if response == nil {
		logrus.WithFields(logrus.Fields{"err": err}).Error("创建标签，响应为空")
		return "", fmt.Errorf("创建标签，响应为空")
	}
	statusCode := response.StatusCode()
	body := response.String()
	logrus.WithFields(logrus.Fields{"statusCode": statusCode, "body": len(body)}).Info("创建标签，响应")
	if statusCode != http.StatusOK {
		logrus.WithFields(logrus.Fields{"StatusCode": statusCode}).Error("创建标签，响应码失败")
		return "", fmt.Errorf("创建标签，响应码失败: %+v", statusCode)
	}
	return body, nil
}
