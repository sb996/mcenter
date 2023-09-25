package impl_test

import (
	"context"

	"github.com/infraboard/mcube/ioc"

	// 注册所有服务
	"github.com/sb996/mcenter/apps/service"
	"github.com/sb996/mcenter/test/tools"
)

var (
	impl service.MetaService
	ctx  = context.Background()
)

func init() {
	tools.DevelopmentSetup()
	impl = ioc.GetController(service.AppName).(service.MetaService)
}
