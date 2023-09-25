package feishu_test

import (
	"context"
	"os"
	"testing"

	"github.com/caarlos0/env/v6"
	"github.com/sb996/mcenter/apps/domain"
	"github.com/sb996/mcenter/apps/notify/provider/im"
	"github.com/sb996/mcenter/apps/notify/provider/im/feishu"
)

var (
	notifyer im.ImNotifyer
	ctx      = context.Background()
)

func TestSendMessage(t *testing.T) {
	req := im.NewSendMessageRequest("验证码", "验证码测试", os.Getenv("FEISHU_USER_ID"))
	if err := notifyer.SendMessage(ctx, req); err != nil {
		t.Fatal(err)
	}
}

func init() {
	conf := domain.NewDefaultFeishuConfig()
	if err := env.Parse(conf); err != nil {
		panic(err)
	}
	notifyer = feishu.NewFeishuNotifyer(conf)
}
