package private_token_test

import (
	"context"
	"testing"

	"github.com/sb996/mcenter/apps/token"
	"github.com/sb996/mcenter/apps/token/provider"
	"github.com/sb996/mcenter/test/tools"
)

var (
	impl provider.TokenIssuer
	ctx  = context.Background()
)

func TestIssueToken(t *testing.T) {
	req := token.NewPrivateTokenIssueTokenRequest("TrXmcSBvVssEgdVPGW948oiR", "测试")
	tk, err := impl.IssueToken(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(tk.Json())
}

func init() {
	tools.DevelopmentSetup()
	impl = provider.GetTokenIssuer(token.GRANT_TYPE_PRIVATE_TOKEN)
}
