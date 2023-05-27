package repositories

import "fmt"

type AlreadySubscribedError struct {
	contact string
}

func NewAlreadySubscribedError(contact string) error {
	return AlreadySubscribedError{contact: contact}
}

func (e AlreadySubscribedError) Error() string {
	return fmt.Sprintf("there is already a subscription for: %s", e.contact)
}

func (e AlreadySubscribedError) Is(target error) bool {
	_, ok := target.(AlreadySubscribedError)
	return ok
}

type UnexpectedRepoError struct {
	dbError error
}

func NewUnexpectedRepoError(dbError error) error {
	return UnexpectedRepoError{dbError: dbError}
}

func (e UnexpectedRepoError) Error() string {
	return fmt.Sprintf("unexpected repo error: %v", e.dbError)
}

func (e UnexpectedRepoError) Is(target error) bool {
	_, ok := target.(UnexpectedRepoError)
	return ok
}

func (e UnexpectedRepoError) Unwrap() error {
	return e.dbError
}
