package datastore

import (
	"github.com/jmoiron/sqlx"
)

type Storage interface {
}

type BeehiveStore struct {
	DB *sqlx.DB
}

func NewDatastore(db *sqlx.DB) *BeehiveStore {
	return &BeehiveStore{
		DB: db,
	}
}
