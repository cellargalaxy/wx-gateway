package controller

import (
	"github.com/cellargalaxy/go_common/util"
	"github.com/cellargalaxy/wx-gateway/model"
	"github.com/cellargalaxy/wx-gateway/service/controller"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
)

//给配置chatId发送tg信息
func sendTgMsg2ConfigChatId(context *gin.Context) {
	var request model.SendTgMsg2ConfigChatIdRequest
	err := context.BindJSON(&request)
	if err != nil {
		logrus.WithContext(context).WithFields(logrus.Fields{"request": request, "err": err}).Error("给配置chatId发送tg信息，请求参数解析异常")
		context.JSON(http.StatusOK, util.CreateErrResponse(err.Error()))
		return
	}
	logrus.WithContext(context).WithFields(logrus.Fields{"request": request}).Info("给配置chatId发送tg信息")
	context.JSON(http.StatusOK, util.CreateResponse(controller.SendTgMsg2ConfigChatId(context, request)))
}
