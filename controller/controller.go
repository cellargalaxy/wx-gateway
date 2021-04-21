package controller

import (
	"fmt"
	"github.com/cellargalaxy/wx-gateway/model"
	"github.com/cellargalaxy/wx-gateway/static"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"math/rand"
	"net/http"
	"strconv"
	"strings"
)

var secretKey = "secret"
var secret = strconv.FormatFloat(rand.Float64(), 'E', -1, 64)

func Controller() error {
	engine := gin.Default()
	store := cookie.NewStore([]byte(secret))
	engine.Use(sessions.Sessions("session_id", store))

	engine.Use(staticCache)
	engine.StaticFS("/static", http.FS(static.StaticFile))

	engine.POST("/api/login", login)

	engine.GET("/api/listAllTemplate", validate, listAllTemplate)
	engine.POST("/api/sendTemplateToTag", validate, sendTemplateToTag)

	engine.POST("/api/createTag", validate, createTag)
	engine.POST("/api/deleteTag", validate, deleteTag)
	engine.GET("/api/listAllTag", validate, listAllTag)
	engine.POST("/api/addTagToUser", validate, addTagToUser)
	engine.POST("/api/deleteTagFromUser", validate, deleteTagFromUser)

	engine.GET("/api/listAllUserInfo", validate, listAllUserInfo)

	err := engine.Run(model.ListenAddress)
	if err != nil {
		logrus.WithFields(logrus.Fields{"err": err}).Warn("web服务启动，异常")
		return fmt.Errorf("web服务启动，异常: %+v", err)
	}
	return nil
}

func staticCache(c *gin.Context) {
	if strings.HasPrefix(c.Request.RequestURI, "/static") {
		c.Header("Cache-Control", "max-age=86400")
	}
}

func createErrResponse(message string, err error) map[string]interface{} {
	return gin.H{"code": model.FailCode, "message": fmt.Sprintf("%+v: %+v", message, err.Error()), "data": nil}
}

func createResponse(data interface{}, err error) map[string]interface{} {
	if err == nil {
		return gin.H{"code": model.SuccessCode, "message": nil, "data": data}
	} else {
		return gin.H{"code": model.FailCode, "message": err.Error(), "data": data}
	}
}
