package symbiosis

import (
	"fmt"
)

type AuthError struct {
	StatusCode int
	Err        error
}

type GenericError struct {
	Status    int32  `json:"status"`
	ErrorType string `json:"error"`
	Message   string `json:"message"`
	Path      string `json:"path"`
}

type NotFoundError struct {
	StatusCode int
	URL        string
	Method     string
}

func (e *AuthError) Error() string {
	return e.Err.Error()
}

func (e *GenericError) Error() string {
	return fmt.Sprintf("Symbiosis: %v (type=%v, path=%v)", e.Message, e.ErrorType, e.Path)
}

func (e *NotFoundError) Error() string {
	return fmt.Sprintf("Symbiosis: %s %s. 404 not found", e.Method, e.URL)
}
