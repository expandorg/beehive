package jobsolutionsmaker

import (
	"errors"
	"reflect"
	"testing"

	"github.com/expandorg/beehive/pkg/honey"
	"github.com/stretchr/testify/assert"

	"github.com/expandorg/beehive/pkg/apierror"
	"github.com/expandorg/beehive/pkg/mock"
	service "github.com/expandorg/beehive/pkg/service"
	"github.com/golang/mock/gomock"
)

func Test_makeJobSolutionsMakerEndpoint(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	s := service.NewMockBeehiveService(ctrl)
	cxt := mock.MockContext{}
	type args struct {
		svc service.BeehiveService
	}
	// No error
	sol := honey.Solution{1, 1, 1, []byte{}}
	sols := honey.JobSolutions{sol}
	req := JobSoltionsRequest{"1", sols}
	s.EXPECT().
		CreateJobSolutions(req.JobID, req.Solutions).
		Return(sols, nil).
		Times(1)

	resp, _ := makeJobSolutionsMakerEndpoint(s)(cxt, req)
	assert.Equal(t, sols, resp)
	// Error
	err := errors.New("error creating honeypots")
	s.EXPECT().
		CreateJobSolutions(req.JobID, req.Solutions).
		Return(honey.JobSolutions{}, err).
		Times(1)
	resp, e := makeJobSolutionsMakerEndpoint(s)(cxt, req)
	assert.Equal(t, nil, resp)
	assert.Equal(t, errorResponse(err), e)
}

func Test_errorResponse(t *testing.T) {
	msg := "error message"
	err := errors.New(msg)
	apiErr := apierror.New(500, msg, err)
	type args struct {
		err error
	}
	tests := []struct {
		name string
		args args
		want *apierror.APIError
	}{
		{
			"it returns an API error",
			args{err},
			apiErr,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := errorResponse(tt.args.err); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("errorResponse() = %v, want %v", got, tt.want)
			}
		})
	}
}
