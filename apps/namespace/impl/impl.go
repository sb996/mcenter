package impl

import (
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/infraboard/mcube/ioc"
	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"
	"google.golang.org/grpc"

	"github.com/sb996/mcenter/apps/counter"
	"github.com/sb996/mcenter/apps/namespace"
	"github.com/sb996/mcenter/apps/policy"
	"github.com/sb996/mcenter/apps/role"
	"github.com/sb996/mcenter/conf"
)

func init() {
	ioc.RegistryController(&impl{})
}

type impl struct {
	col *mongo.Collection
	log logger.Logger
	namespace.UnimplementedRPCServer
	ioc.IocObjectImpl

	counter counter.Service
	role    role.Service
	policy  policy.Service
}

func (i *impl) Init() error {
	db, err := conf.C().Mongo.GetDB()
	if err != nil {
		return err
	}
	i.col = db.Collection(i.Name())
	i.log = zap.L().Named(i.Name())

	i.role = ioc.GetController(role.AppName).(role.Service)
	i.policy = ioc.GetController(policy.AppName).(policy.Service)
	i.counter = ioc.GetController(counter.AppName).(counter.Service)
	return nil
}

func (i *impl) Name() string {
	return namespace.AppName
}

func (i *impl) Registry(server *grpc.Server) {
	namespace.RegisterRPCServer(server, i)
}
