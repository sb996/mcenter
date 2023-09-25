package impl_test

import (
	"context"

	"github.com/infraboard/mcube/ioc"
	"github.com/sb996/mcenter/apps/user"
	"github.com/sb996/mcenter/test/tools"
)

var (
	impl user.Service
	ctx  = context.Background()
)

func init() {
	tools.DevelopmentSetup()
	impl = ioc.GetController(user.AppName).(user.Service)
}
