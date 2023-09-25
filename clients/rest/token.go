package rest

import (
	"context"

	"github.com/infraboard/mcube/client/rest"
	"github.com/sb996/mcenter/apps/token"
)

type TokenService interface {
	// 校验Token
	ValidateToken(context.Context, *token.ValidateTokenRequest) (*token.Token, error)
}

type tokenImpl struct {
	client *rest.RESTClient
}

func (i *tokenImpl) ValidateToken(ctx context.Context, req *token.ValidateTokenRequest) (*token.Token, error) {
	ins := token.NewDefaultToken()

	err := i.client.
		Get("token").
		Header(token.VALIDATE_TOKEN_HEADER_KEY, req.AccessToken).
		Do(ctx).
		Into(ins)
	if err != nil {
		return nil, err
	}

	return ins, nil
}
