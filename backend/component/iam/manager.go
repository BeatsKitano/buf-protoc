package iam

import (
	"context"
	_ "embed"

	"buf-protoc/backend/domain"
	"buf-protoc/backend/store"
)

//go:embed acl.yaml
var aclYaml []byte

type acl struct {
	Roles []struct {
		Name        string   `yaml:"name"`
		Title       string   `yaml:"title"`
		Permissions []string `yaml:"permissions"`
	} `yaml:"roles"`
}

type Manager struct {
	groupMembers    map[string]map[string]bool
	PredefinedRoles []*domain.Role
	store           *store.Store
}

func NewManager(store *store.Store) (*Manager, error) {
	m := &Manager{
		store: store,
	}
	return m, nil
}

// Check if the user has permission on the resource hierarchy.
// CEL on the binding is not considered.
// When multiple projects are specified, the user should have permission on every projects.
func (m *Manager) CheckPermission(ctx context.Context, p Permission, userID int) (bool, error) {
	return false, nil
}

// GetPermissions returns all permissions for the given role.
// Role format is roles/{role}.
func (m *Manager) GetPermissions(role string) (map[Permission]bool, error) {
	return nil, nil
}
