package services

import "fmt"

type DuplicateError struct {
	resource string
}

func NewDuplicateError(duplicateResource string) error {
	return DuplicateError{
		resource: duplicateResource,
	}
}

func (e DuplicateError) Error() string {
	return fmt.Sprintf("%s already exists", e.resource)
}

func (e DuplicateError) Is(target error) bool {
	_, ok := target.(DuplicateError)
	return ok
}

type InternalError struct{}

func NewInternalError() error {
	return InternalError{}
}

func (e InternalError) Error() string {
	return "internal server error"
}

func (e InternalError) Is(target error) bool {
	_, ok := target.(InternalError)
	return ok
}
