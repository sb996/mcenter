package impl_test

import (
	"context"

	"github.com/infraboard/mcenter/apps/policy"
	"github.com/infraboard/mcenter/test/tools"
	"github.com/infraboard/mcube/ioc"
)

var (
	impl policy.Service
	ctx  = context.Background()
)

func init() {
	tools.DevelopmentSetup()
	impl = ioc.GetController(policy.AppName).(policy.Service)
}
