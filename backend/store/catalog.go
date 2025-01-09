package store

import (
	"buf-protoc/backend/domain"
	"context"
)

type FindCatalogOption struct {
	Id         *int
	Name       *string
	PathPrefix *string
	ParentID   *int
}

func (s *Store) SaveCatalog(ctx context.Context, catalog *domain.Catalog) error {
	return nil
}

func (s *Store) DeleteCatalog(ctx context.Context, id int) error {
	return nil
}

func (s *Store) UpdateCatalog(ctx context.Context, catalog *domain.Catalog) error {
	return nil
}

func (s *Store) FindCatalogById(ctx context.Context, option *FindCatalogOption) (*domain.Catalog, error) {
	return nil, nil
}

func (s *Store) FindCatalogs(ctx context.Context) ([]*domain.Catalog, error) {
	return nil, nil
}
