package simpdb

import "fmt"

type DuplicateRecordError struct {
	record string
}

func NewDuplicateRecordError(record string) error {
	return DuplicateRecordError{record: record}
}

func (e DuplicateRecordError) Error() string {
	return fmt.Sprintf("db already contains record: %s", e.record)
}

func (e DuplicateRecordError) Is(target error) bool {
	_, ok := target.(DuplicateRecordError)
	return ok
}

func (e DuplicateRecordError) GetRecord() string {
	return e.record
}

type UnexpectedError struct {
	error error
}

func NewUnexpectedError(err error) error {
	return UnexpectedError{error: err}
}

func (e UnexpectedError) Error() string {
	return fmt.Sprintf("unexpected db error: %v", e.error)
}

func (e UnexpectedError) Is(target error) bool {
	_, ok := target.(UnexpectedError)
	return ok
}

func (e UnexpectedError) Unwrap() error {
	return e.error
}
