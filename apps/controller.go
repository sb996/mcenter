package apps

import (
	// 注册所有内部服务模块, 无须对外暴露的服务, 用于内部依赖
	_ "github.com/sb996/mcenter/apps/counter/impl"
	_ "github.com/sb996/mcenter/apps/ip2region/impl"
	_ "github.com/sb996/mcenter/apps/storage/impl"

	// 注册所有GRPC服务模块, 暴露给框架GRPC服务器加载, 注意 导入有先后顺序
	_ "github.com/sb996/mcenter/apps/domain/impl"
	_ "github.com/sb996/mcenter/apps/endpoint/impl"
	_ "github.com/sb996/mcenter/apps/instance/impl"
	_ "github.com/sb996/mcenter/apps/label/impl"
	_ "github.com/sb996/mcenter/apps/namespace/impl"
	_ "github.com/sb996/mcenter/apps/notify/impl"
	_ "github.com/sb996/mcenter/apps/policy/impl"
	_ "github.com/sb996/mcenter/apps/resource/impl"
	_ "github.com/sb996/mcenter/apps/role/impl"
	_ "github.com/sb996/mcenter/apps/service/impl"
	_ "github.com/sb996/mcenter/apps/token/impl"
	_ "github.com/sb996/mcenter/apps/user/impl"
)
