package github

import (
	"context"

	"github.com/google/go-github/v45/github"
	"github.com/sb996/mcenter/apps/service"
)

func (s *Github) ListOrganizations(ctx context.Context) (*service.ServiceSet, error) {
	s.client.Organizations.List(ctx, "", &github.ListOptions{})
	// s.client.Repositories.List()
	// r, resp, err := s.client.Repositories.List(ctx, "", &github.RepositoryListOptions{})
	// s.client.Repositories.CreateHook()
	return nil, nil
}

func (s *Github) ListProjects(ctx context.Context) (*service.ServiceSet, error) {
	s.client.Organizations.ListProjects(ctx, "", &github.ProjectListOptions{})
	return nil, nil
}
