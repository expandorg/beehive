package jobsolutionsmaker

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/expandorg/beehive/pkg/apierror"
	"github.com/expandorg/beehive/pkg/honey"
	service "github.com/expandorg/beehive/pkg/service"
	kithttp "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
)

func MakeHandler(s service.BeehiveService) http.Handler {
	return kithttp.NewServer(
		makeJobSolutionsMakerEndpoint(s),
		decodeResponseMakerRequest,
		encodeResponse,
	)
}

func encodeResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	return json.NewEncoder(w).Encode(response)
}

func decodeResponseMakerRequest(_ context.Context, r *http.Request) (interface{}, error) {
	vars := mux.Vars(r)
	jobID, ok := vars["job_id"]
	if !ok {
		return nil, fmt.Errorf("missing job_id parameter")
	}

	var sols honey.JobSolutions
	decoder := json.NewDecoder(r.Body)

	err := decoder.Decode(&sols)
	if err != nil {
		return nil, apierror.New(500, err.Error(), err)
	}

	if valid, err := validateRequest(sols); !valid {
		return nil, err
	}

	return JobSoltionsRequest{jobID, sols}, nil
}
