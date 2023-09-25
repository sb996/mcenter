package impl_test

import (
	"context"

	"github.com/infraboard/mcube/ioc"
	"github.com/sb996/mcenter/apps/token"
	"github.com/sb996/mcenter/test/tools"
)

var (
	impl token.Service
	ctx  = context.Background()
)

func init() {
	tools.DevelopmentSetup()
	impl = ioc.GetController(token.AppName).(token.Service)
}
