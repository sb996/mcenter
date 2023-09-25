package rpc

import (
	"net/url"
	"time"

	"github.com/sb996/mcenter/apps/instance"
)

// NewDefaultConfig todo
func NewDefaultConfig() *Config {
	return &Config{
		Address:       "localhost:18010",
		TimeoutSecond: 10,
		Resolver:      NewDefaultResolver(),
	}
}

// Config 客户端配置
type Config struct {
	Address      string `json:"address" toml:"address" yaml:"address" env:"MCENTER_GRPC_ADDRESS"`
	ClientID     string `json:"client_id" toml:"client_id" yaml:"client_id" env:"MCENTER_CLINET_ID"`
	ClientSecret string `json:"client_secret" toml:"client_secret" yaml:"client_secret" env:"MCENTER_CLIENT_SECRET"`
	// 默认值10秒
	TimeoutSecond uint      `json:"timeout_second" toml:"timeout_second" yaml:"timeout_second" env:"MCENTER_GRPC_TIMEOUT_SECOND"`
	Resolver      *Resolver `json:"resolver" toml:"resolver" yaml:"resolver"`
}

func (c *Config) Timeout() time.Duration {
	return time.Second * time.Duration(c.TimeoutSecond)
}

func (c *Config) WithCredentials(clientId, clientSecret string) {
	c.ClientID = clientId
	c.ClientSecret = clientSecret
}

func (c *Config) Credentials() *Authentication {
	return NewAuthentication(c.ClientID, c.ClientSecret)
}

func NewDefaultResolver() *Resolver {
	return &Resolver{
		Region:      instance.DEFAULT_REGION,
		Environment: instance.DEFAULT_ENV,
		Cluster:     instance.DEFAULT_CLUSTER,
		Group:       instance.DEFAULT_GROUP,
	}
}

type Resolver struct {
	// 实例所属地域, 默认default
	Region string `json:"region" toml:"region" yaml:"region" env:"MCENTER_INSTANCE_REGION" validate:"required"`
	// 实例所属环境, 默认default
	Environment string `json:"environment" toml:"environment" yaml:"environment" env:"MCENTER_INSTANCE_ENV" validate:"required"`
	// 实例所属集群,默认default
	Cluster string `json:"cluster" toml:"cluster" yaml:"cluster" env:"MCENTER_INSTANCE_CLUSTER" validate:"required"`
	// 实例所属分组,默认default
	Group string `json:"group" toml:"group" yaml:"group" env:"MCENTER_INSTANCE_GROUP" validate:"required"`
	// 实例标签, 可以根据标签快速过滤实例, 格式k=v,k=v
	Labels string `json:"labels" toml:"labels" yaml:"labels" env:"MCENTER_INSTANCE_LABELS"`
}

func (r *Resolver) ToQueryString() string {
	m := make(url.Values)
	m.Add("region", r.Region)
	m.Add("environment", r.Environment)
	m.Add("cluster", r.Cluster)
	m.Add("group", r.Group)
	m.Add("labels", r.Labels)
	return m.Encode()
}
