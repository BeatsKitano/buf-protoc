package repository

import (
	"context"
)

func (s *Repository) CheckRoleInUse(ctx context.Context, role string) (bool, error) {
	return false, nil
}
