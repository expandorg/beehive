package solutionmaker

import (
	"encoding/json"
	"reflect"
	"testing"

	"github.com/expandorg/beehive/pkg/apierror"
	"github.com/expandorg/beehive/pkg/honey"
)

func Test_validateRequest(t *testing.T) {
	type args struct {
		req honey.Solution
	}
	tests := []struct {
		name  string
		args  args
		want  bool
		want1 *apierror.APIError
	}{
		{
			"is valid if task_id and data are not empty",
			args{honey.Solution{0, 1, 1, json.RawMessage{}}},
			true,
			nil,
		},
		{
			"is not valid if task_id is empty",
			args{honey.Solution{0, 0, 1, []byte{}}},
			false,
			errorResponse(&apierror.MissingParameters{Params: []string{"task_id"}}),
		},
		{
			"is not valid if job_id is empty",
			args{honey.Solution{0, 1, 0, []byte{}}},
			false,
			errorResponse(&apierror.MissingParameters{Params: []string{"job_id"}}),
		},
		{
			"is not valid if task_id, job_id and data are empty",
			args{honey.Solution{0, 0, 0, nil}},
			false,
			errorResponse(&apierror.MissingParameters{Params: []string{"task_id", "job_id", "data"}}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := validateRequest("1", tt.args.req)
			if got != tt.want {
				t.Errorf("validateRequest() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("validateRequest() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
