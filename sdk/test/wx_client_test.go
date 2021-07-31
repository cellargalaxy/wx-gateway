package test

import (
	"context"
	"github.com/cellargalaxy/go_common/util"
	"github.com/cellargalaxy/wx-gateway/sdk"
	"testing"
	"time"
)

func TestWxClient(test *testing.T) {
	ctx := context.Background()
	wxClient, err := sdk.NewWxClient(time.Second*3, time.Second*3, 3, "http://127.0.0.1:8990", "secret")
	if err != nil {
		test.Error(err)
		test.FailNow()
	}
	response, err := wxClient.SendTemplateToTag(ctx, "7ub0o1jXJGfar5Zaj-imwwoisFiH6xW6CsS4pKWjnKc", 109, "", map[string]interface{}{"zhi": 111})
	test.Logf("response: %+v\r\n", util.ToJsonIndent(response))
	if err != nil {
		test.Error(err)
		test.FailNow()
	}
}
