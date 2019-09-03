package jobsolutionsmaker

import (
	"github.com/gemsorg/beehive/pkg/apierror"
	"github.com/gemsorg/beehive/pkg/honey"
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
