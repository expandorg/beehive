package jobsolutionsmaker

import (
	"encoding/json"
	"reflect"
	"testing"

	"github.com/gemsorg/beehive/pkg/apierror"
	"github.com/gemsorg/beehive/pkg/honey"
)

func Test_validateRequest(t *testing.T) {
	type args struct {
		req honey.JobSolutions
	}
	tests := []struct {
		name  string
		args  args
		want  bool
		want1 *apierror.APIError
	}{
		{
			"is valid if solutions are not empty",
			args{honey.JobSolutions{honey.Solution{0, 1, 1, json.RawMessage{}}}},
			true,
			nil,
		},
		{
			"is not valid if solutions are empty",
			args{honey.JobSolutions{}},
			false,
			errorResponse(&apierror.MissingParameters{Params: []string{"solutions"}}),
		},
		{
			"is not valid if solutions are empty",
			args{},
			false,
			errorResponse(&apierror.MissingParameters{Params: []string{"solutions"}}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := validateRequest(tt.args.req)
			if got != tt.want {
				t.Errorf("validateRequest() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("validateRequest() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
