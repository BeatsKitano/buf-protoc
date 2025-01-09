package store

import (
	"context"
)

func (s *Store) CheckRoleInUse(ctx context.Context, role string) (bool, error) {
	return false, nil
}
