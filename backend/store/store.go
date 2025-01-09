package store

import (
	lru "github.com/hashicorp/golang-lru/v2"
)

// Store is the store for the application.
type Store struct {
	db *struct{}

	catalogCache *lru.Cache[int, interface{}]
}

func NewStore() *Store {
	return &Store{}
}
