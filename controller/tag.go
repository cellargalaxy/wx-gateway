package controller

import (
	"github.com/cellargalaxy/go_common/util"
	"github.com/cellargalaxy/wx-gateway/model"
	"github.com/cellargalaxy/wx-gateway/service/controller"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
)

//创建标签
func createTag(context *gin.Context) {
	var request model.CreateTagRequest
	err := context.BindJSON(&request)
	if err != nil {
		logrus.WithContext(context).WithFields(logrus.Fields{"request": request, "err": err}).Error("创建标签，请求参数解析异常")
		context.JSON(http.StatusOK, util.CreateErrResponse(err.Error()))
		return
	}
	logrus.WithContext(context).WithFields(logrus.Fields{"request": request}).Info("创建标签")
	context.JSON(http.StatusOK, util.CreateResponse(controller.CreateTag(context, request)))
}

//删除标签
func deleteTag(context *gin.Context) {
	var request model.DeleteTagRequest
	err := context.BindJSON(&request)
	if err != nil {
		logrus.WithContext(context).WithFields(logrus.Fields{"request": request, "err": err}).Error("删除标签，请求参数解析异常")
		context.JSON(http.StatusOK, util.CreateErrResponse(err.Error()))
		return
	}
	logrus.WithContext(context).WithFields(logrus.Fields{"request": request}).Info("删除标签")
	context.JSON(http.StatusOK, util.CreateResponse(controller.DeleteTag(context, request)))
}

//获取所有标签
func listAllTag(context *gin.Context) {
	var request model.ListAllTagRequest
	err := context.BindQuery(&request)
	if err != nil {
		logrus.WithContext(context).WithFields(logrus.Fields{"request": request, "err": err}).Error("获取所有标签，请求参数解析异常")
		context.JSON(http.StatusOK, util.CreateErrResponse(err.Error()))
		return
	}
	logrus.WithContext(context).WithFields(logrus.Fields{"request": request}).Info("获取所有标签")
	context.JSON(http.StatusOK, util.CreateResponse(controller.ListAllTag(context, request)))
}

//为用户加标签
func addTagToUser(context *gin.Context) {
	var request model.AddTagToUserRequest
	err := context.BindJSON(&request)
	if err != nil {
		logrus.WithContext(context).WithFields(logrus.Fields{"request": request, "err": err}).Error("为用户加标签，请求参数解析异常")
		context.JSON(http.StatusOK, util.CreateErrResponse(err.Error()))
		return
	}
	logrus.WithContext(context).WithFields(logrus.Fields{"request": request}).Info("为用户加标签")
	context.JSON(http.StatusOK, util.CreateResponse(controller.AddTagToUser(context, request)))
}

//为用户删标签
func deleteTagFromUser(context *gin.Context) {
	var request model.DeleteTagFromUserRequest
	err := context.BindJSON(&request)
	if err != nil {
		logrus.WithContext(context).WithFields(logrus.Fields{"request": request, "err": err}).Error("为用户删标签，请求参数解析异常")
		context.JSON(http.StatusOK, util.CreateErrResponse(err.Error()))
		return
	}
	logrus.WithContext(context).WithFields(logrus.Fields{"request": request}).Info("为用户删标签")
	context.JSON(http.StatusOK, util.CreateResponse(controller.DeleteTagFromUser(context, request)))
}
