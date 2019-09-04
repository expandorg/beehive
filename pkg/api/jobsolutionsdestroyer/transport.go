package jobsolutionsdestroyer

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	service "github.com/gemsorg/beehive/pkg/service"
	kithttp "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
)

func MakeHandler(s service.BeehiveService) http.Handler {
	return kithttp.NewServer(
		makeJobSolutionsDestroyerEndpoint(s),
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
	return JobSoltionsRequest{jobID}, nil
}
