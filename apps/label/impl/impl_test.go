package impl_test

import (
	"context"

	"github.com/sb996/mcenter/apps/label"
	"github.com/sb996/mcenter/test/tools"
	"github.com/infraboard/mcube/ioc"
)

var (
	impl label.Service
	ctx  = context.Background()
)

func init() {
	tools.DevelopmentSetup()
	impl = ioc.GetController(label.AppName).(label.Service)
}
