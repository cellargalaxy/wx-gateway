package controller

import (
	"github.com/cellargalaxy/go_common/util"
	"github.com/cellargalaxy/wx-gateway/model"
	"github.com/cellargalaxy/wx-gateway/service/controller"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
)

//获取全部模板
func listAllTemplate(context *gin.Context) {
	var request model.ListAllTemplateRequest
	err := context.BindQuery(&request)
	if err != nil {
		logrus.WithContext(context).WithFields(logrus.Fields{"request": request, "err": err}).Error("获取全部模板，请求参数解析异常")
		context.JSON(http.StatusOK, util.CreateErrResponse(err.Error()))
		return
	}
	logrus.WithContext(context).WithFields(logrus.Fields{"request": request}).Info("获取全部模板")
	context.JSON(http.StatusOK, util.CreateResponse(controller.ListAllTemplate(context, request)))
}

//给标签用户发送模板消息
func sendTemplateToTag(context *gin.Context) {
	var request model.SendTemplateToTagRequest
	err := context.BindJSON(&request)
	if err != nil {
		logrus.WithContext(context).WithFields(logrus.Fields{"request": request, "err": err}).Error("给标签用户发送模板消息，请求参数解析异常")
		context.JSON(http.StatusOK, util.CreateErrResponse(err.Error()))
		return
	}
	logrus.WithContext(context).WithFields(logrus.Fields{"request": request}).Info("给标签用户发送模板消息")
	context.JSON(http.StatusOK, util.CreateResponse(controller.SendTemplateToTag(context, request)))
}
