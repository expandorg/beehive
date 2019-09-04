package datastore

import (
	"fmt"
	"strings"

	"github.com/gemsorg/beehive/pkg/honey"
	"github.com/jmoiron/sqlx"
)

type Storage interface {
	CreateSolution(sol honey.Solution) (honey.Solution, error)
	CreateJobSolutions(jobID string, sols honey.JobSolutions) (honey.JobSolutions, error)
	DeleteJobSolutions(jobID string) (bool, error)
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

func (b *BeehiveStore) CreateJobSolutions(jobID string, sols honey.JobSolutions) (honey.JobSolutions, error) {
	vals := []string{}
	for _, sol := range sols {

		vals = append(vals, fmt.Sprintf("(%d, %s, %q)", sol.TaskID, jobID, sol.Data))
	}

	tx, err := b.DB.Begin()
	attrQuery := "INSERT INTO solutions (task_id, job_id, data) VALUES" + strings.Join(vals, ",")

	_, err = tx.Exec(attrQuery)
	if err != nil {
		tx.Rollback()
		return honey.JobSolutions{}, err
	}

	err = tx.Commit()
	if err != nil {
		return honey.JobSolutions{}, err
	}
	return sols, nil
}

func (b *BeehiveStore) DeleteJobSolutions(jobID string) (bool, error) {
	_, err := b.DB.Exec("DELETE FROM solutions WHERE job_id=?", jobID)
	if err != nil {
		return false, err
	}
	return true, nil
}
