package role

import (
	"github.com/infraboard/mcube/http/label"
	resource "github.com/infraboard/mcube/pb/resource"
	"github.com/sb996/mcenter/apps/domain"
	"github.com/sb996/mcenter/apps/namespace"
)

const (
	// AdminRoleName 管理员角色名称, 系统初始化时创建
	ADMIN_ROLE_NAME = "admin"
	// VisitorRoleName 访客角色名称, 系统初始化时创建
	VISITOR_ROLE_NAME = "visitor"
)

const (
	// RoleMaxPermission 一个角色允许添加的最大权限数
	RoleMaxPermission = 500
)

func CreateAdminRoleRequest(createBy string) *CreateRoleRequest {
	req := NewCreateRoleRequest()
	req.CreateBy = createBy
	req.Domain = domain.DEFAULT_DOMAIN
	req.Namespace = namespace.SYSTEM_NAMESPACE
	req.Visiable = resource.VISIABLE_GLOBAL
	req.Name = ADMIN_ROLE_NAME
	req.Description = "管理员"
	req.Specs = []*PermissionSpec{
		{
			Desc:         "管理员具备读写权限",
			Effect:       EffectType_ALLOW,
			ServiceId:    "*",
			ResourceName: "*",
			MatchAll:     true,
		},
	}
	return req
}

func CreateVisitorRoleRequest(createBy string) *CreateRoleRequest {
	req := NewCreateRoleRequest()
	req.CreateBy = createBy
	req.Domain = domain.DEFAULT_DOMAIN
	req.Namespace = namespace.SYSTEM_NAMESPACE
	req.Visiable = resource.VISIABLE_GLOBAL
	req.Name = VISITOR_ROLE_NAME
	req.Description = "访客"
	req.Specs = []*PermissionSpec{
		{
			Desc:         "访客具备可读权限",
			Effect:       EffectType_ALLOW,
			ServiceId:    "*",
			ResourceName: "*",
			LabelValues:  []string{label.List.Value(), label.Get.Value()},
		},
	}
	return req
}

func NamespaceManagerPermssion() *Permission {
	p := NewDeaultPermission()
	p.Spec = &PermissionSpec{
		Desc:         "拥有资源的所有操作权限",
		Effect:       EffectType_ALLOW,
		ServiceId:    "*",
		ResourceName: "*",
		MatchAll:     true,
	}
	return p
}
