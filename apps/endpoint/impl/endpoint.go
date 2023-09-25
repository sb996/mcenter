package impl

import (
	"context"

	"github.com/infraboard/mcube/exception"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/sb996/mcenter/apps/endpoint"
	"github.com/sb996/mcenter/apps/service"
)

func (s *impl) DescribeEndpoint(ctx context.Context, req *endpoint.DescribeEndpointRequest) (
	*endpoint.Endpoint, error) {
	r, err := newDescribeEndpointRequest(req)
	if err != nil {
		return nil, err
	}

	ins := endpoint.NewDefaultEndpoint()
	if err := s.col.FindOne(ctx, r.FindFilter()).Decode(ins); err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, exception.NewNotFound("endpoint %s not found", req)
		}

		return nil, exception.NewInternalServerError("find endpoint %s error, %s", req.Id, err)
	}

	return ins, nil
}

func (s *impl) QueryEndpoints(ctx context.Context, req *endpoint.QueryEndpointRequest) (
	*endpoint.EndpointSet, error) {
	r := newQueryEndpointRequest(req)
	resp, err := s.col.Find(ctx, r.FindFilter(), r.FindOptions())

	if err != nil {
		return nil, exception.NewInternalServerError("find endpoint error, error is %s", err)
	}

	set := endpoint.NewEndpointSet()
	// 循环
	for resp.Next(ctx) {
		app := endpoint.NewDefaultEndpoint()
		if err := resp.Decode(app); err != nil {
			return nil, exception.NewInternalServerError("decode domain error, error is %s", err)
		}

		set.Add(app)
	}

	// count
	count, err := s.col.CountDocuments(ctx, r.FindFilter())
	if err != nil {
		return nil, exception.NewInternalServerError("get device count error, error is %s", err)
	}
	set.Total = count

	return set, nil
}

// Client 信息可以通过参数传递进来(HTTP协议时)
// GRPC客户端时 也可以通过ctx传递进来
func (s *impl) RegistryEndpoint(ctx context.Context, req *endpoint.RegistryRequest) (*endpoint.RegistryResponse, error) {
	// 获取认证后的client 服务相关信息
	svc, err := service.GetServiceFromCtx(ctx)
	if err != nil {
		return nil, err
	}
	if svc.Meta != nil && svc.Meta.Id != "" {
		req.ServiceId = svc.Meta.Id
		req.SetExtension("domain", svc.Spec.Domain)
		req.SetExtension("namespace", svc.Spec.Namespace)
	}

	if err := req.Validate(); err != nil {
		return nil, exception.NewBadRequest(err.Error())
	}

	// 生成该服务的Endpoint
	endpoints := req.Endpoints(req.ServiceId)
	s.log.Debugf("registry endpoints: %s", endpoints)

	// 更新已有的记录
	news := make([]interface{}, 0, len(endpoints))
	for i := range endpoints {
		if err := s.col.FindOneAndReplace(ctx, bson.M{"_id": endpoints[i].Id}, endpoints[i]).Err(); err != nil {
			if err == mongo.ErrNoDocuments {
				news = append(news, endpoints[i])
			} else {
				return nil, err
			}
		}
	}

	// 插入新增记录
	if len(news) > 0 {
		if _, err := s.col.InsertMany(ctx, news); err != nil {
			return nil, exception.NewInternalServerError("inserted a service document error, %s", err)
		}
	}

	return endpoint.NewRegistryResponse("ok"), nil
}

func (s *impl) DeleteEndpoint(ctx context.Context, req *endpoint.DeleteEndpointRequest) (*endpoint.Endpoint, error) {
	result, err := s.col.DeleteOne(ctx, bson.M{"service_id": req.ServiceId})
	if err != nil {
		return nil, exception.NewInternalServerError("delete service(%s) endpoint error, %s", req.ServiceId, err)
	}

	s.log.Infof("delete service %s endpoint success, total count: %d", req.ServiceId, result.DeletedCount)
	return nil, nil
}
