package feishu

import (
	"context"

	"github.com/infraboard/mcube/logger"
	lark "github.com/larksuite/oapi-sdk-go/v3"
	"github.com/sb996/mcenter/apps/domain"
)

type Feishu struct {
	conf   *domain.FeishuConfig
	client *lark.Client
	log    logger.Logger
}

func (f *Feishu) ListUser(ctx context.Context) {
	f.client.Contact.User.List(ctx, nil)
}
