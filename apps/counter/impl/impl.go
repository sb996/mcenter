package impl

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/sb996/mcenter/apps/counter"
	"github.com/sb996/mcenter/conf"
	"github.com/infraboard/mcube/ioc"
)

func init() {
	ioc.RegistryController(&service{})
}

type service struct {
	col *mongo.Collection
	ioc.IocObjectImpl
}

func (s *service) Init() error {
	db, err := conf.C().Mongo.GetDB()
	if err != nil {
		return err
	}

	dc := db.Collection("counter")
	s.col = dc
	return nil
}

func (s *service) GetNextSequenceValue(sequenceName string) (*counter.Count, error) {
	result := s.col.FindOneAndUpdate(
		context.Background(),
		bson.M{"_id": sequenceName},
		bson.M{"$inc": bson.M{"value": 1}},
		options.FindOneAndUpdate().SetUpsert(true),
	)

	count := counter.NewCount()
	err := result.Decode(count)
	if err != nil && err != mongo.ErrNoDocuments {
		return nil, fmt.Errorf("counter decode error, %s", err)
	}

	count.Value++
	return count, nil
}

func (s *service) Name() string {
	return counter.AppName
}
