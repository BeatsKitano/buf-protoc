package repository

import (
	"buf-protoc/backend/domain"
	"context"
)

func (s *Repository) SaveCatalog(ctx context.Context, catalog *domain.Catalog) error {
	return nil
}

func (s *Repository) DeleteCatalog(ctx context.Context, id int) error {
	return nil
}

func (s *Repository) UpdateCatalog(ctx context.Context, catalog *domain.Catalog) error {
	return nil
}

func (s *Repository) FindCatalogById(ctx context.Context, filter *domain.CatalogFilter) (*domain.Catalog, error) {
	return nil, nil
}

func (s *Repository) FindCatalogs(ctx context.Context) ([]*domain.Catalog, error) {
	return nil, nil
}
