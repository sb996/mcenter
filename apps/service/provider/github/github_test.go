package github_test

import (
	"context"
	"testing"

	"github.com/infraboard/mcube/logger/zap"
	"github.com/sb996/mcenter/apps/service/provider/github"
)

var (
	client *github.Github
	ctx    = context.Background()
)

func TestAuthCodeURL(t *testing.T) {
	t.Log(client.AuthCodeURL())
}

func TestExchange(t *testing.T) {
	err := client.Exchange(ctx, "ecb383090de8db2a3828")
	if err != nil {
		t.Fatal(err)
	}
}

func init() {
	zap.DevelopmentSetup()
	conf, err := github.LoadConfigFromEnv()
	if err != nil {
		panic(err)
	}
	client = github.NewGithub(conf)
}
