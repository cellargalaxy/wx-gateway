package controller

import (
	"errors"
	"fmt"
	"github.com/cellargalaxy/wx-gateway/config"
	"github.com/cellargalaxy/wx-gateway/model"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
)

//登录检查
func validate(context *gin.Context) {
	if isLogin(context) {
		context.Next()
		return
	}
	context.Abort()
	context.JSON(http.StatusUnauthorized, createErrResponse("please login", errors.New("please login")))
}

//登录
func login(context *gin.Context) {
	var request model.LoginRequest
	err := context.BindJSON(&request)
	if err != nil {
		logrus.WithFields(logrus.Fields{"request": request, "err": err}).Error("登录，请求参数解析异常")
		createErrResponse("登录，请求参数解析异常", err)
		return
	}
	if request.Token != config.Config.Token {
		logrus.WithFields(logrus.Fields{"request": request}).Error("登录，失败")
		context.JSON(http.StatusOK, createResponse("illegal token", errors.New("illegal token")))
		return
	}
	err = setLogin(context)
	if err != nil {
		context.JSON(http.StatusOK, createResponse("login fail", err))
		return
	}
	context.JSON(http.StatusOK, createResponse("login success", nil))
}

func setLogin(context *gin.Context) error {
	session := sessions.Default(context)
	session.Set(secretKey, secret)
	err := session.Save()
	if err != nil {
		logrus.WithFields(logrus.Fields{"err": err}).Error("登录，异常")
		return fmt.Errorf("登录，异常: %+v", err)
	}
	return err
}

func isLogin(context *gin.Context) bool {
	request := context.Request
	if request != nil && request.Header != nil && request.Header.Get(model.TokenHeadKey) == config.Config.Token {
		return true
	}
	session := sessions.Default(context)
	sessionSecret := session.Get(secretKey)
	return sessionSecret == secret
}
