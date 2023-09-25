package impl

import (
	"context"

	"github.com/infraboard/mcube/http/request"
	"github.com/sb996/mcenter/apps/endpoint"
	"github.com/sb996/mcenter/apps/resource"
)

const (
	// MaxQueryEndpoints todo
	MaxQueryEndpoints = 1000
)

func (s *impl) QueryResources(ctx context.Context, req *resource.QueryResourceRequest) (
	*resource.ResourceSet, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}

	rs := resource.NewResourceSet()
	queryE := endpoint.NewQueryEndpointRequest(request.NewPageRequest(MaxQueryEndpoints, 1))
	queryE.PermissionEnable = req.PermissionEnable
	queryE.Resources = req.Resources
	queryE.ServiceIds = req.ServiceIds
	eps, err := s.ep.QueryEndpoints(ctx, queryE)
	if err != nil {
		return nil, err
	}
	if eps.Total > MaxQueryEndpoints {
		s.log.Warnf("service %s total endpoints > %d", req.ServiceIds, eps.Total)
	}

	rs.AddEndpointSet(eps)
	return rs, nil
}
