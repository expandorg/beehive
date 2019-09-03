package solutionmaker

import (
	"strconv"

	"github.com/gemsorg/beehive/pkg/apierror"
	"github.com/gemsorg/beehive/pkg/honey"
)

func validateRequest(potID string, req honey.Solution) (bool, *apierror.APIError) {
	var missingParams []string
	requestID, err := strconv.ParseUint(potID, 10, 64)
	if err != nil || req.TaskID == 0 || req.TaskID != requestID {
		missingParams = append(missingParams, "task_id")
	}
	if req.JobID == 0 {
		missingParams = append(missingParams, "job_id")
	}
	if req.Data == nil {
		missingParams = append(missingParams, "data")
	}
	if len(missingParams) > 0 {
		return false, errorResponse(&apierror.MissingParameters{Params: missingParams})
	}
	return true, nil
}
