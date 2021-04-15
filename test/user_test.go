package test

import (
	"github.com/cellargalaxy/wx-gateway/model"
	"github.com/cellargalaxy/wx-gateway/service/controller"
	"github.com/cellargalaxy/wx-gateway/util"
	"testing"
)

func TestListAllUserInfo(test *testing.T) {
	request := model.ListAllUserInfoRequest{}
	response, err := controller.ListAllUserInfo(request)
	test.Logf("response: %+v\r\n", util.ToJsonIndent(response))
	if err != nil {
		test.Error(err)
		test.FailNow()
	}
}
