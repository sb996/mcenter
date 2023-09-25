package impl

import (
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/infraboard/mcube/ioc"
	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"
	"google.golang.org/grpc"

	"github.com/sb996/mcenter/apps/domain"
	"github.com/sb996/mcenter/apps/notify"
	"github.com/sb996/mcenter/apps/user"
	"github.com/sb996/mcenter/conf"
)

func init() {
	ioc.RegistryController(&service{})
}

type service struct {
	col *mongo.Collection
	notify.UnimplementedRPCServer
	ioc.IocObjectImpl

	user   user.Service
	log    logger.Logger
	domain domain.Service
}

func (s *service) Init() error {
	db, err := conf.C().Mongo.GetDB()
	if err != nil {
		return err
	}

	s.col = db.Collection(s.Name())
	s.log = zap.L().Named(s.Name())
	s.user = ioc.GetController(user.AppName).(user.Service)
	s.domain = ioc.GetController(domain.AppName).(domain.Service)
	return nil
}

func (s *service) Name() string {
	return notify.AppName
}

func (s *service) Registry(server *grpc.Server) {
	notify.RegisterRPCServer(server, s)
}
