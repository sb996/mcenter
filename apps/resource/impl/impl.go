package impl

import (
	"github.com/infraboard/mcube/ioc"
	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"
	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/grpc"

	"github.com/sb996/mcenter/apps/endpoint"
	"github.com/sb996/mcenter/apps/resource"
	"github.com/sb996/mcenter/conf"
)

func init() {
	ioc.RegistryController(&impl{})
}

type impl struct {
	col *mongo.Collection
	log logger.Logger
	resource.UnimplementedRPCServer
	ioc.IocObjectImpl

	ep endpoint.Service
}

func (i *impl) Init() error {
	db, err := conf.C().Mongo.GetDB()
	if err != nil {
		return err
	}
	i.col = db.Collection(i.Name())
	i.log = zap.L().Named(i.Name())

	i.ep = ioc.GetController(endpoint.AppName).(endpoint.Service)
	return nil
}

func (i *impl) Name() string {
	return resource.AppName
}

func (i *impl) Registry(server *grpc.Server) {
	resource.RegisterRPCServer(server, i)
}
