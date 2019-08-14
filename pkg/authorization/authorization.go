package authorization

import (
	"github.com/gemsorg/beehive/pkg/authentication"
)

type Authorizer interface {
	SetAuthData(data authentication.AuthData)
}

type authorizor struct {
	authData authentication.AuthData
}

func NewAuthorizer() Authorizer {
	return &authorizor{
		authentication.AuthData{},
	}
}

func (a *authorizor) SetAuthData(data authentication.AuthData) {
	a.authData = data
}