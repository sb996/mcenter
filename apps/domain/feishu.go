package domain

import "github.com/infraboard/mcube/tools/sense"

// NewDefaultConfig represents the default LDAP config.
func NewDefaultFeishuConfig() *FeishuConfig {
	return &FeishuConfig{}
}

func (s *FeishuConfig) DeSense() {
	s.AppSecret = sense.DeSense(s.AppSecret)
}
