package repositories

import "fmt"

type AlreadySubscribedError struct {
	contact string
}

func NewAlreadySubscribedError(contact string) *AlreadySubscribedError {
	return &AlreadySubscribedError{contact: contact}
}

func (e AlreadySubscribedError) Error() string {
	return fmt.Sprintf("there is already a subscription for: %s", e.contact)
}

func (e AlreadySubscribedError) GetContact() string {
	return e.contact
}

type UnexpectedRepoError struct {
	dbError error
}

func NewUnexpectedRepoError(dbError error) *UnexpectedRepoError {
	return &UnexpectedRepoError{dbError: dbError}
}

func (e UnexpectedRepoError) Error() string {
	return fmt.Sprintf("unexpected repo error: %v", e.dbError)
}
