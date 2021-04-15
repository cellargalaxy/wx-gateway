package test

import (
	"github.com/cellargalaxy/wx-gateway/model"
	"github.com/cellargalaxy/wx-gateway/service/controller"
	"github.com/cellargalaxy/wx-gateway/util"
	"testing"
)

func TestCreateTag(test *testing.T) {
	request := model.CreateTagRequest{
		Tag: "",
	}
	response, err := controller.CreateTag(request)
	test.Logf("response: %+v\r\n", util.ToJsonIndent(response))
	if err != nil {
		test.Error(err)
		test.FailNow()
	}
}

func TestDeleteTag(test *testing.T) {
	request := model.DeleteTagRequest{
		TagId: 0,
	}
	response, err := controller.DeleteTag(request)
	test.Logf("response: %+v\r\n", util.ToJsonIndent(response))
	if err != nil {
		test.Error(err)
		test.FailNow()
	}
}

func TestListAllTag(test *testing.T) {
	request := model.ListAllTagRequest{}
	response, err := controller.ListAllTag(request)
	test.Logf("response: %+v\r\n", util.ToJsonIndent(response))
	if err != nil {
		test.Error(err)
		test.FailNow()
	}
}

func TestAddTagToUser(test *testing.T) {
	request := model.AddTagToUserRequest{
		TagId:  0,
		OpenId: "",
	}
	response, err := controller.AddTagToUser(request)
	test.Logf("response: %+v\r\n", util.ToJsonIndent(response))
	if err != nil {
		test.Error(err)
		test.FailNow()
	}
}

func TestDeleteTagFromUser(test *testing.T) {
	request := model.DeleteTagFromUserRequest{
		TagId:  0,
		OpenId: "",
	}
	response, err := controller.DeleteTagFromUser(request)
	test.Logf("response: %+v\r\n", util.ToJsonIndent(response))
	if err != nil {
		test.Error(err)
		test.FailNow()
	}
}
