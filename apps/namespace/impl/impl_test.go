package impl_test

import (
	"context"

	"github.com/sb996/mcenter/apps/namespace"
	"github.com/sb996/mcenter/test/tools"
	"github.com/infraboard/mcube/ioc"
)

var (
	impl namespace.Service
	ctx  = context.Background()
)

func init() {
	tools.DevelopmentSetup()
	impl = ioc.GetController(namespace.AppName).(namespace.Service)
}
