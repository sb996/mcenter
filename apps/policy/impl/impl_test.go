package impl_test

import (
	"context"

	"github.com/infraboard/mcube/ioc"
	"github.com/sb996/mcenter/apps/policy"
	"github.com/sb996/mcenter/test/tools"
)

var (
	impl policy.Service
	ctx  = context.Background()
)

func init() {
	tools.DevelopmentSetup()
	impl = ioc.GetController(policy.AppName).(policy.Service)
}
