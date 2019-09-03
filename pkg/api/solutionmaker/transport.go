package solutionmaker

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gemsorg/beehive/pkg/apierror"
	"github.com/gemsorg/beehive/pkg/honey"
	service "github.com/gemsorg/beehive/pkg/service"
	kithttp "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
)

func MakeHandler(s service.BeehiveService) http.Handler {
	return kithttp.NewServer(
		makeSolutionMakerEndpoint(s),
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
	var ok bool
	potID, ok := vars["task_id"]
	if !ok {
		return nil, fmt.Errorf("missing task_id parameter")
	}

	var sol honey.Solution
	decoder := json.NewDecoder(r.Body)

	err := decoder.Decode(&sol)
	if err != nil {
		return nil, apierror.New(500, err.Error(), err)
	}
	if valid, err := validateRequest(potID, sol); !valid {
		return nil, err
	}
	return sol, nil
}
