package test

import (
	"context"
	"github.com/cellargalaxy/go_common/util"
	"github.com/cellargalaxy/wx-gateway/sdk"
	"testing"
	"time"
)

func TestSendWxTemplateToTag(test *testing.T) {
	util.InitLog("msg_gateway.log")
	ctx := context.Background()
	ctx = util.SetLogId(ctx)
	wxClient, err := sdk.NewMsgClient(time.Second*3, time.Second*3, 3, "http://127.0.0.1:8990", "secret")
	if err != nil {
		test.Error(err)
		test.FailNow()
	}
	response, err := wxClient.SendWxTemplateToTag(ctx, "7ub0o1jXJGfar5Zaj-imwwoisFiH6xW6CsS4pKWjnKc", 109, "", map[string]interface{}{"zhi": 111})
	test.Logf("response: %+v\r\n", util.ToJsonIndentString(response))
	if err != nil {
		test.Error(err)
		test.FailNow()
	}
}

func TestSendTgMsg2ConfigChatId(test *testing.T) {
	util.InitLog("msg_gateway.log")
	ctx := context.Background()
	ctx = util.SetLogId(ctx)
	wxClient, err := sdk.NewMsgClient(time.Second*3, time.Second*3, 3, "http://127.0.0.1:8990", "secret")
	if err != nil {
		test.Error(err)
		test.FailNow()
	}
	response, err := wxClient.SendTgMsg2ConfigChatId(ctx, "")
	test.Logf("response: %+v\r\n", util.ToJsonIndentString(response))
	if err != nil {
		test.Error(err)
		test.FailNow()
	}
}
