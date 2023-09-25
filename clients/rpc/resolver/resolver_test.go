package resolver_test

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/sb996/mcenter/clients/rpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/sb996/mcenter/clients/rpc/resolver"
)

func TestResolver(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	// 连接到服务
	conn, err := grpc.DialContext(
		ctx,
		// Dial to "mcenter://mpaas"
		fmt.Sprintf("%s://%s", resolver.Scheme, "mpaas"),
		// 认证
		grpc.WithPerRPCCredentials(rpc.NewAuthenticationFromEnv()),
		// 不开启TLS
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		// gprc 支持的负载均衡策略: https://github.com/grpc/grpc/blob/master/doc/load-balancing.md
		grpc.WithDefaultServiceConfig(`{"loadBalancingConfig": [{"round_robin":{}}]}`),
		// 直到建立连接
		grpc.WithBlock(),
	)

	if err != nil {
		t.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
}
