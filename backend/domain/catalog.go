package domain

import "context"

type Catalog struct {
	Id       int    `json:"id,omitempty"`
	Name     string `json:"name,omitempty"`
	Path     string `json:"path,omitempty"`
	ParentID int    `json:"parent_id,omitempty"`
}

type CatalogFilter struct {
	Id         *int
	Name       *string
	PathPrefix *string
	ParentID   *int
}

type CatalogRepository interface {
	FindById(ctx context.Context, id uint64) (*Catalog, error)
	Save(ctx context.Context, user *Catalog) error
}
