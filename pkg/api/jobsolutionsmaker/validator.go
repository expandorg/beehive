package jobsolutionsmaker

import (
	"github.com/expandorg/beehive/pkg/apierror"
	"github.com/expandorg/beehive/pkg/honey"
)

func validateRequest(req honey.JobSolutions) (bool, *apierror.APIError) {
	var missingParams []string

	if req == nil || len(req) == 0 {
		missingParams = append(missingParams, "solutions")
	}
	if len(missingParams) > 0 {
		return false, errorResponse(&apierror.MissingParameters{Params: missingParams})
	}
	return true, nil
}
