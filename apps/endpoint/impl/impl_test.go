package impl_test

import (
	"context"

	"github.com/sb996/mcenter/apps/endpoint"
	"github.com/sb996/mcenter/test/tools"
	"github.com/infraboard/mcube/ioc"
)

var (
	impl endpoint.Service
	ctx  = context.Background()
)

func init() {
	tools.DevelopmentSetup()
	impl = ioc.GetController(endpoint.AppName).(endpoint.Service)
}
