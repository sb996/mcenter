package impl_test

import (
	"testing"

	"github.com/sb996/mcenter/apps/domain"
	"github.com/sb996/mcenter/apps/namespace"
)

func TestCreateNamespace(t *testing.T) {
	req := namespace.NewCreateNamespaceRequest()
	req.Domain = domain.DEFAULT_DOMAIN
	req.Name = namespace.DEFAULT_NAMESPACE
	req.Owner = "admin"
	r, err := impl.CreateNamespace(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(r)
}

func TestQueryNamespace(t *testing.T) {
	req := namespace.NewQueryNamespaceRequest()
	req.Domain = domain.DEFAULT_DOMAIN
	r, err := impl.QueryNamespace(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(r)
}

func TestDescribeNamespaceById(t *testing.T) {
	req := namespace.NewDescriptNamespaceRequestById("5fecb49ae23804e7")
	r, err := impl.DescribeNamespace(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(r)
}

func TestDescribeNamespace(t *testing.T) {
	req := namespace.NewDescriptNamespaceRequestByName(domain.DEFAULT_DOMAIN, namespace.DEFAULT_NAMESPACE)
	r, err := impl.DescribeNamespace(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(r)
}
