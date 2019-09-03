package service

import (
	"github.com/gemsorg/beehive/pkg/authorization"
	"github.com/gemsorg/beehive/pkg/datastore"
	"github.com/gemsorg/beehive/pkg/honey"
)

type BeehiveService interface {
	Healthy() bool
	CreateSolution(hp honey.Solution) (honey.Solution, error)
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
