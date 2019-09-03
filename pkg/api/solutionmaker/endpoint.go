package solutionmaker

import (
	"context"

	"github.com/gemsorg/beehive/pkg/apierror"
	"github.com/gemsorg/beehive/pkg/honey"

	service "github.com/gemsorg/beehive/pkg/service"

	"github.com/go-kit/kit/endpoint"
)

func makeSolutionMakerEndpoint(svc service.BeehiveService) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req := request.(honey.Solution)
		saved, err := svc.CreateSolution(req)
		if err != nil {
			return nil, errorResponse(err)
		}
		return saved, nil
	}
}

func errorResponse(err error) *apierror.APIError {
	return apierror.New(500, err.Error(), err)
}
