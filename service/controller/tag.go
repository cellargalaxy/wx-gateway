package controller

import (
	"github.com/cellargalaxy/wx-gateway/model"
	"github.com/cellargalaxy/wx-gateway/service/wx"
)

//创建标签
func CreateTag(request model.CreateTagRequest) (*model.CreateTagResponse, error) {
	result, err := wx.CreateTag(request.Tag)
	return &model.CreateTagResponse{Result: result}, err
}

//删除标签
func DeleteTag(request model.DeleteTagRequest) (*model.DeleteTagResponse, error) {
	result, err := wx.DeleteTag(request.TagId)
	return &model.DeleteTagResponse{Result: result}, err
}

//获取所有标签
func ListAllTag(request model.ListAllTagRequest) (*model.ListAllTagResponse, error) {
	list, err := wx.ListAllTag()
	return &model.ListAllTagResponse{List: list}, err
}

//为用户加标签
func AddTagToUser(request model.AddTagToUserRequest) (*model.AddTagToUserResponse, error) {
	result, err := wx.AddTagToUser(request.TagId, []string{request.OpenId})
	return &model.AddTagToUserResponse{Result: result}, err
}

//为用户删标签
func DeleteTagFromUser(request model.DeleteTagFromUserRequest) (*model.DeleteTagFromUserResponse, error) {
	result, err := wx.DeleteTagFromUser(request.TagId, []string{request.OpenId})
	return &model.DeleteTagFromUserResponse{Result: result}, err
}
