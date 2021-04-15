package test

import (
	"github.com/cellargalaxy/wx-gateway/model"
	"github.com/cellargalaxy/wx-gateway/service/controller"
	"github.com/cellargalaxy/wx-gateway/util"
	"testing"
)

func TestListAllTemplate(test *testing.T) {
	request := model.ListAllTemplateRequest{}
	response, err := controller.ListAllTemplate(request)
	test.Logf("response: %+v\r\n", util.ToJsonIndent(response))
	if err != nil {
		test.Error(err)
		test.FailNow()
	}
}

func TestSendTemplateToTag(test *testing.T) {
	request := model.SendTemplateToTagRequest{
		TemplateId: "7ub0o1jXJGfar5Zaj-imwwoisFiH6xW6CsS4pKWjnKc",
		TagId:      108,
		Url:        "https://baidu.com",
		Data:       map[string]interface{}{},
	}
	response, err := controller.SendTemplateToTag(request)
	test.Logf("response: %+v\r\n", util.ToJsonIndent(response))
	if err != nil {
		test.Error(err)
		test.FailNow()
	}
}
