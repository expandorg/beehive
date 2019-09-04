package jobsolutionsdestroyer

import (
	"context"

	"github.com/gemsorg/beehive/pkg/apierror"

	service "github.com/gemsorg/beehive/pkg/service"

	"github.com/go-kit/kit/endpoint"
)

func makeJobSolutionsDestroyerEndpoint(svc service.BeehiveService) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req := request.(JobSoltionsRequest)
		saved, err := svc.DeleteJobSolutions(req.JobID)
		if err != nil {
			return nil, errorResponse(err)
		}
		return saved, nil
	}
}

type JobSoltionsRequest struct {
	JobID string
}

func errorResponse(err error) *apierror.APIError {
	return apierror.New(500, err.Error(), err)
}
