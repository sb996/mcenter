package api

import (
	restfulspec "github.com/emicklei/go-restful-openapi/v2"
	"github.com/emicklei/go-restful/v3"
	"github.com/infraboard/mcube/http/label"
	"github.com/infraboard/mcube/http/restful/response"
	"github.com/infraboard/mcube/ioc"
	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"

	"github.com/sb996/mcenter/apps/policy"
	"github.com/sb996/mcenter/apps/token"
)

func init() {
	ioc.RegistryApi(&policyHandler{})
}

type policyHandler struct {
	service policy.Service
	log     logger.Logger
	ioc.IocObjectImpl
}

func (h *policyHandler) Init() error {
	h.log = zap.L().Named(policy.AppName)
	h.service = ioc.GetController(policy.AppName).(policy.Service)
	return nil
}

func (h *policyHandler) Name() string {
	return "policy"
}

func (h *policyHandler) Version() string {
	return "v1"
}

func (h *policyHandler) Registry(ws *restful.WebService) {
	tags := []string{"策略管理"}

	ws.Route(ws.POST("/").To(h.CreatePolicy).
		Doc("创建策略").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Metadata(label.Resource, h.Name()).
		Metadata(label.Action, label.Create.Value()).
		Metadata(label.Auth, label.Enable).
		Metadata(label.Permission, label.Enable).
		Reads(policy.CreatePolicyRequest{}).
		Writes(policy.Policy{}))

	ws.Route(ws.GET("/").To(h.QueryPolicy).
		Doc("查询策略列表").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Metadata(label.Resource, h.Name()).
		Metadata(label.Action, label.List.Value()).
		Metadata(label.Auth, label.Enable).
		Metadata(label.Permission, label.Disable).
		Writes(policy.PolicySet{}).
		Returns(200, "OK", policy.PolicySet{}).
		Returns(404, "Not Found", nil))

	ws.Route(ws.GET("/{id}").To(h.DescribePolicy).
		Doc("查询策略详情").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Metadata(label.Resource, h.Name()).
		Metadata(label.Action, label.Get.Value()).
		Metadata(label.Auth, label.Enable).
		Metadata(label.Permission, label.Disable).
		Writes(policy.Policy{}).
		Returns(200, "OK", policy.Policy{}).
		Returns(404, "Not Found", nil))

	ws.Route(ws.POST("/{id}").To(h.CreatePolicy).
		Doc("删除策略").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Metadata(label.Resource, h.Name()).
		Metadata(label.Action, label.Delete.Value()).
		Metadata(label.Auth, label.Enable).
		Metadata(label.Permission, label.Enable).
		Reads(policy.DeletePolicyRequest{}).
		Writes(policy.Policy{}))
}

func (h *policyHandler) CreatePolicy(r *restful.Request, w *restful.Response) {
	req := policy.NewCreatePolicyRequest()
	if err := r.ReadEntity(req); err != nil {
		response.Failed(w, err)
		return
	}

	req.UpdateFromToken(token.GetTokenFromRequest(r))
	set, err := h.service.CreatePolicy(r.Request.Context(), req)
	if err != nil {
		response.Failed(w, err)
		return
	}

	response.Success(w, set)
}

func (h *policyHandler) QueryPolicy(r *restful.Request, w *restful.Response) {
	req := policy.NewQueryPolicyRequestFromHTTP(r)

	set, err := h.service.QueryPolicy(r.Request.Context(), req)
	if err != nil {
		response.Failed(w, err)
		return
	}

	response.Success(w, set)
}

func (h *policyHandler) DescribePolicy(r *restful.Request, w *restful.Response) {
	req := policy.NewDescriptPolicyRequest(r.PathParameter("id"))
	ins, err := h.service.DescribePolicy(r.Request.Context(), req)
	if err != nil {
		response.Failed(w, err)
		return
	}

	response.Success(w, ins)
}

func (h *policyHandler) DeletePolicy(r *restful.Request, w *restful.Response) {
	req := policy.NewDeletePolicyRequestWithID(r.PathParameter("id"))

	req.Scope = token.GetTokenFromRequest(r).GenScope()
	set, err := h.service.DeletePolicy(r.Request.Context(), req)
	if err != nil {
		response.Failed(w, err)
		return
	}
	response.Success(w, set)
}
