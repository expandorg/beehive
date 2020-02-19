package jobsolutionsmaker

import (
	"context"

	"github.com/expandorg/beehive/pkg/apierror"
	"github.com/expandorg/beehive/pkg/honey"

	service "github.com/expandorg/beehive/pkg/service"

	"github.com/go-kit/kit/endpoint"
)

func makeJobSolutionsMakerEndpoint(svc service.BeehiveService) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req := request.(JobSoltionsRequest)
		saved, err := svc.CreateJobSolutions(req.JobID, req.Solutions)
		if err != nil {
			return nil, errorResponse(err)
		}
		return saved, nil
	}
}

type JobSoltionsRequest struct {
	JobID     string
	Solutions honey.JobSolutions
}

func errorResponse(err error) *apierror.APIError {
	return apierror.New(500, err.Error(), err)
}
