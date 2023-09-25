package impl_test

import (
	"context"

	"github.com/infraboard/mcube/ioc"
	"github.com/sb996/mcenter/apps/domain"
	"github.com/sb996/mcenter/test/tools"
)

var (
	impl domain.Service
	ctx  = context.Background()
)

func init() {
	tools.DevelopmentSetup()
	impl = ioc.GetController(domain.AppName).(domain.Service)
}
