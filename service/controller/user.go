package controller

import (
	"github.com/cellargalaxy/wx-gateway/model"
	"github.com/cellargalaxy/wx-gateway/service/wx"
)

//获取全部用户信息
func ListAllUserInfo(request model.ListAllUserInfoRequest) (*model.ListAllUserInfoResponse, error) {
	list, err := wx.ListAllUserInfo()
	return &model.ListAllUserInfoResponse{List: list}, err
}
