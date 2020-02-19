package jobsolutionsdestroyer

import (
	"errors"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/expandorg/beehive/pkg/apierror"
	"github.com/expandorg/beehive/pkg/mock"
	service "github.com/expandorg/beehive/pkg/service"
	"github.com/golang/mock/gomock"
)

func Test_makeJobSolutionsDestroyerEndpoint(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	s := service.NewMockBeehiveService(ctrl)
	cxt := mock.MockContext{}
	type args struct {
		svc service.BeehiveService
	}
	// No error
	req := JobSoltionsRequest{"1"}
	s.EXPECT().
		DeleteJobSolutions(req.JobID).
		Return(true, nil).
		Times(1)

	resp, _ := makeJobSolutionsDestroyerEndpoint(s)(cxt, req)
	assert.Equal(t, true, resp)
	// Error
	err := errors.New("error creating honeypots")
	s.EXPECT().
		DeleteJobSolutions(req.JobID).
		Return(false, err).
		Times(1)
	resp, e := makeJobSolutionsDestroyerEndpoint(s)(cxt, req)
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
