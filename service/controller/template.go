package controller

import (
	"github.com/cellargalaxy/wx-gateway/model"
	"github.com/cellargalaxy/wx-gateway/service/wx"
)

//获取全部模板
func ListAllTemplate(request model.ListAllTemplateRequest) (*model.ListAllTemplateResponse, error) {
	list, err := wx.ListAllTemplate()
	return &model.ListAllTemplateResponse{List: list}, err
}

//给标签用户发送模板消息
func SendTemplateToTag(request model.SendTemplateToTagRequest) (*model.SendTemplateToTagResponse, error) {
	failOpenIds, err := wx.SendTemplateToTag(request.TemplateId, request.TagId, request.Url, request.Data)
	return &model.SendTemplateToTagResponse{FailOpenIds: failOpenIds}, err
}
