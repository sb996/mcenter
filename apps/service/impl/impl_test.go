package impl_test

import (
	"context"

	"github.com/infraboard/mcube/ioc"

	// 注册所有服务
	"github.com/infraboard/mcenter/apps/service"
	"github.com/infraboard/mcenter/test/tools"
)

var (
	impl service.MetaService
	ctx  = context.Background()
)

func init() {
	tools.DevelopmentSetup()
	impl = ioc.GetController(service.AppName).(service.MetaService)
}
