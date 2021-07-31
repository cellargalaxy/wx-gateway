package controller

import (
	"github.com/cellargalaxy/go_common/util"
	"github.com/cellargalaxy/wx-gateway/model"
	"github.com/cellargalaxy/wx-gateway/service/controller"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
)

//获取全部用户信息
func listAllUserInfo(context *gin.Context) {
	var request model.ListAllUserInfoRequest
	err := context.BindQuery(&request)
	if err != nil {
		logrus.WithContext(context).WithFields(logrus.Fields{"request": request, "err": err}).Error("获取全部用户信息，请求参数解析异常")
		context.JSON(http.StatusOK, util.CreateErrResponse(err.Error()))
		return
	}
	logrus.WithContext(context).WithFields(logrus.Fields{"request": request}).Info("获取全部用户信息")
	context.JSON(http.StatusOK, util.CreateResponse(controller.ListAllUserInfo(context, request)))
}
