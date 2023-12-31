package impl_test

import (
	"testing"

	"github.com/sb996/mcenter/apps/domain"
	"github.com/sb996/mcenter/apps/namespace"
	"github.com/sb996/mcenter/apps/policy"
)

func TestCheckPermissionOk(t *testing.T) {
	req := policy.NewCheckPermissionRequest()
	req.Domain = domain.DEFAULT_DOMAIN
	req.Namespace = namespace.DEFAULT_NAMESPACE
	req.Username = "test02"

	// 检查test用户在默认空间下是否有访问mpaas服务的构建配置功能
	req.ServiceId = "mcenter-api"
	req.Path = "GET./mcenter/api/v1/service"
	r, err := impl.CheckPermission(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(r.ToJson())
}

func TestCheckPermissionDeny(t *testing.T) {
	req := policy.NewCheckPermissionRequest()
	req.Domain = domain.DEFAULT_DOMAIN
	req.Namespace = namespace.DEFAULT_NAMESPACE
	req.Username = "test"

	// 检查是否有创建Pipeline权限
	req.ServiceId = "cd08fc9c"
	req.Path = "POST./mpaas/api/v1/pipelines"
	_, err := impl.CheckPermission(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
}
