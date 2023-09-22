package impl

import (
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/infraboard/mcube/ioc"
	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"
	"google.golang.org/grpc"

	"github.com/infraboard/mcenter/apps/endpoint"
	"github.com/infraboard/mcenter/apps/namespace"
	"github.com/infraboard/mcenter/apps/policy"
	"github.com/infraboard/mcenter/apps/role"
	"github.com/infraboard/mcenter/apps/user"
	"github.com/infraboard/mcenter/conf"
)

func init() {
	ioc.RegistryController(&impl{})
}

type impl struct {
	col *mongo.Collection
	log logger.Logger
	policy.UnimplementedRPCServer
	ioc.IocObjectImpl

	user      user.Service
	role      role.Service
	namespace namespace.Service
	endpoint  endpoint.Service
}

func (i *impl) Init() error {
	db, err := conf.C().Mongo.GetDB()
	if err != nil {
		return err
	}
	i.col = db.Collection(i.Name())
	i.log = zap.L().Named(i.Name())

	i.user = ioc.GetController(user.AppName).(user.Service)
	i.role = ioc.GetController(role.AppName).(role.Service)
	i.namespace = ioc.GetController(namespace.AppName).(namespace.Service)
	i.endpoint = ioc.GetController(endpoint.AppName).(endpoint.Service)
	return nil
}

func (i *impl) Name() string {
	return policy.AppName
}

func (i *impl) Registry(server *grpc.Server) {
	policy.RegisterRPCServer(server, i)
}
