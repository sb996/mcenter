package apps

import (
	// 注册所有HTTP服务模块, 暴露给框架HTTP服务器加载
	_ "github.com/sb996/mcenter/apps/domain/api"
	_ "github.com/sb996/mcenter/apps/endpoint/api"
	_ "github.com/sb996/mcenter/apps/instance/api"
	_ "github.com/sb996/mcenter/apps/label/api"
	_ "github.com/sb996/mcenter/apps/namespace/api"
	_ "github.com/sb996/mcenter/apps/policy/api"
	_ "github.com/sb996/mcenter/apps/resource/api"
	_ "github.com/sb996/mcenter/apps/role/api"
	_ "github.com/sb996/mcenter/apps/service/api"
	_ "github.com/sb996/mcenter/apps/token/api"
	_ "github.com/sb996/mcenter/apps/user/api"
)
