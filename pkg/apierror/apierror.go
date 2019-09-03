package apierror

import (
	"encoding/json"
	"fmt"
)

type APIError struct {
	statusCode   int    `json:"-"`
	ErrorMessage string `json:"error"`
	error        error  `json:"-"`
}

type errorPayload struct {
	ErrorMessage string `json:"error"`
}

func New(statusCode int, error string, err error) *APIError {
	return &APIError{
		statusCode:   statusCode,
		ErrorMessage: error,
		error:        err,
	}
}

func (a *APIError) MarshalJSON() ([]byte, error) {
	return json.Marshal(errorPayload{
		ErrorMessage: a.ErrorMessage,
	})
}

func (a *APIError) StatusCode() int {
	return a.statusCode
}

func (a *APIError) Error() string {
	return fmt.Sprintf("Status: %d, ErrorMessage: [%s]:%s", a.statusCode, a.ErrorMessage, a.error)
}
