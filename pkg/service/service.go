package service

import (
	"github.com/expandorg/beehive/pkg/authorization"
	"github.com/expandorg/beehive/pkg/datastore"
	"github.com/expandorg/beehive/pkg/honey"
)

type BeehiveService interface {
	Healthy() bool
	CreateSolution(hp honey.Solution) (honey.Solution, error)
	CreateJobSolutions(jobID string, hp honey.JobSolutions) (honey.JobSolutions, error)
	DeleteJobSolutions(jobID string) (bool, error)
}

type service struct {
	store      datastore.Storage
	authorizor authorization.Authorizer
}

func New(s datastore.Storage, a authorization.Authorizer) *service {
	return &service{
		store:      s,
		authorizor: a,
	}
}

func (s *service) Healthy() bool {
	return true
}

func (s *service) CreateSolution(sol honey.Solution) (honey.Solution, error) {
	return s.store.CreateSolution(sol)
}

func (s *service) CreateJobSolutions(jobID string, sols honey.JobSolutions) (honey.JobSolutions, error) {
	return s.store.CreateJobSolutions(jobID, sols)
}

func (s *service) DeleteJobSolutions(jobID string) (bool, error) {
	return s.store.DeleteJobSolutions(jobID)
}
