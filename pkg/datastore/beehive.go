package datastore

import (
	"github.com/gemsorg/beehive/pkg/honey"
	"github.com/jmoiron/sqlx"
)

type Storage interface {
	CreateSolution(sol honey.Solution) (honey.Solution, error)
}

type BeehiveStore struct {
	DB *sqlx.DB
}

func NewDatastore(db *sqlx.DB) *BeehiveStore {
	return &BeehiveStore{
		DB: db,
	}
}

func (b *BeehiveStore) CreateSolution(sol honey.Solution) (honey.Solution, error) {
	result, err := b.DB.Exec(
		"INSERT INTO solutions (task_id, job_id, data) VALUES (?, ?, ?)",
		sol.TaskID, sol.JobID, sol.Data,
	)
	if err != nil {
		return honey.Solution{}, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return honey.Solution{}, err
	}
	sol.ID = uint64(id)
	return sol, nil
}
