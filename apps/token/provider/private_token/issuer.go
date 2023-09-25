package private_token

import (
	"context"

	"github.com/infraboard/mcube/exception"
	"github.com/infraboard/mcube/ioc"
	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"
	"github.com/sb996/mcenter/apps/token"
	"github.com/sb996/mcenter/apps/token/provider"
	"github.com/sb996/mcenter/apps/user"
)

type issuer struct {
	token token.Service
	user  user.Service

	log logger.Logger
}

func (i *issuer) Init() error {
	i.token = ioc.GetController(token.AppName).(token.Service)
	i.user = ioc.GetController(user.AppName).(user.Service)
	i.log = zap.L().Named("issuer.privat_token")
	return nil
}

func (i *issuer) GrantType() token.GRANT_TYPE {
	return token.GRANT_TYPE_PRIVATE_TOKEN
}

func (i *issuer) validate(ctx context.Context, accessToken string) (*user.User, error) {
	if accessToken == "" {
		return nil, exception.NewUnauthorized("access token required")
	}

	// 判断凭证合法性
	tk, err := i.token.ValidateToken(ctx, token.NewValidateTokenRequest(accessToken))
	if err != nil {
		return nil, err
	}

	u, err := i.user.DescribeUser(ctx, user.NewDescriptUserRequestById(tk.UserId))
	if err != nil {
		return nil, err
	}

	return u, nil
}

func (i *issuer) IssueToken(ctx context.Context, req *token.IssueTokenRequest) (*token.Token, error) {
	u, err := i.validate(ctx, req.AccessToken)
	if err != nil {
		return nil, err
	}

	// 3. 颁发Token
	newTk := token.NewToken(req)
	newTk.Domain = u.Spec.Domain
	newTk.Username = u.Spec.Username
	newTk.UserType = u.Spec.Type
	newTk.UserId = u.Meta.Id

	return newTk, nil
}

func (i *issuer) IssueCode(ctx context.Context, req *token.IssueCodeRequest) (*token.Code, error) {
	_, err := i.validate(ctx, req.AccessToken)
	if err != nil {
		return nil, err
	}

	// 颁发Token
	c, err := token.NewCode(req)
	if err != nil {
		return nil, err
	}

	return c, nil
}

func init() {
	provider.Registe(&issuer{})
}
