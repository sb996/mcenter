package impl_test

import (
	"context"

	"github.com/infraboard/mcube/ioc"
	"github.com/sb996/mcenter/apps/role"
	"github.com/sb996/mcenter/test/tools"
)

var (
	impl role.Service
	ctx  = context.Background()
)

func init() {
	tools.DevelopmentSetup()
	impl = ioc.GetController(role.AppName).(role.Service)
}
