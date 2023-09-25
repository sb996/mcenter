package impl_test

import (
	"context"

	// 注册所有服务
	"github.com/infraboard/mcube/ioc"
	"github.com/sb996/mcenter/apps/notify"
	"github.com/sb996/mcenter/test/tools"
)

var (
	impl notify.Service
	ctx  = context.Background()
)

func init() {
	tools.DevelopmentSetup()
	impl = ioc.GetController(notify.AppName).(notify.Service)
}
