package repository

import (
	lru "github.com/hashicorp/golang-lru/v2"
)

// Repository is the repository for the application.
type Repository struct {
	db *struct{}

	catalogCache *lru.Cache[int, interface{}]
}

func NewRepository() *Repository {
	return &Repository{}
}
