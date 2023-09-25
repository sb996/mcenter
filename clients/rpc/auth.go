package rpc

import (
	"context"
	"os"

	"github.com/sb996/mcenter/apps/service"
)

// 客户端携带的凭证
func NewAuthentication(clientId, clientSecret string) *Authentication {
	return &Authentication{
		clientID:     clientId,
		clientSecret: clientSecret,
	}
}

func NewAuthenticationFromEnv() *Authentication {
	return NewAuthentication(
		os.Getenv("MCENTER_CLINET_ID"),
		os.Getenv("MCENTER_CLIENT_SECRET"),
	)
}

// Authentication todo
type Authentication struct {
	clientID     string
	clientSecret string
}

// SetClientCredentials todo
func (a *Authentication) SetClientCredentials(clientID, clientSecret string) {
	a.clientID = clientID
	a.clientSecret = clientSecret
}

// GetRequestMetadata todo
func (a *Authentication) GetRequestMetadata(context.Context, ...string) (
	map[string]string, error,
) {
	return map[string]string{
		service.ClientHeaderKey: a.clientID,
		service.ClientSecretKey: a.clientSecret,
	}, nil
}

// RequireTransportSecurity todo
func (a *Authentication) RequireTransportSecurity() bool {
	return false
}
