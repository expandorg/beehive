package datastore

import (
	"strconv"

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
	solutions := honey.JobSolutions{}
	for _, sol := range sols {
		id, _ := strconv.ParseUint(jobID, 10, 64)
		sol.JobID = id
	}

	tx := b.DB.MustBegin()
	results, err := tx.NamedExec("INSERT INTO solutions (task_id, job_id, data) VALUES (:task_id, :job_id, :data)", sols)

	if err != nil {
		tx.Rollback()
		return solutions, err
	}

	ra, err := results.RowsAffected()
	if err != nil || ra != int64(len(sols)) {
		tx.Rollback()
		return solutions, err
	}

	err = tx.Select(&solutions, "SELECT * FROM solutions WHERE job_id = ?", jobID)
	if err != nil {
		tx.Rollback()
		return solutions, err
	}

	err = tx.Commit()
	if err != nil {
		return solutions, err
	}

	return solutions, nil
}

func (b *BeehiveStore) DeleteJobSolutions(jobID string) (bool, error) {
	_, err := b.DB.Exec("DELETE FROM solutions WHERE job_id=?", jobID)
	if err != nil {
		return false, err
	}
	return true, nil
}
