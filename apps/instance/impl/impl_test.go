package impl_test

import (
	"context"

	"github.com/sb996/mcenter/apps/instance"
	"github.com/sb996/mcenter/test/tools"
	"github.com/infraboard/mcube/ioc"
)

var (
	impl instance.Service
	ctx  = context.Background()
)

func init() {
	tools.DevelopmentSetup()
	impl = ioc.GetController(instance.AppName).(instance.Service)
}
