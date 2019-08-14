package service

import (
	"github.com/gemsorg/beehive/pkg/authorization"
	"github.com/gemsorg/beehive/pkg/datastore"
)

type BeehiveService interface {
	Healthy() bool
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
