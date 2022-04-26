package symbiosis

import (
	"fmt"
)

type AuthError struct {
	StatusCode int
	Err        error
}

type GenericError struct {
	Status    int32
	ErrorType string `json:"error"`
	Message   string
	Path      string
}

type NotFoundError struct {
	StatusCode int
}

func (e *AuthError) Error() string {
	return e.Err.Error()
}

func (e *GenericError) Error() string {
	return fmt.Sprintf("Symbiosis: %v (type=%v, path=%v)", e.Message, e.ErrorType, e.Path)
}

func (e *NotFoundError) Error() string {
	return fmt.Sprintf("Symbiosis: %v (type=%v, path=%v) not found")
}
