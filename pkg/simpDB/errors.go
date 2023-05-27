package simpdb

import "fmt"

type DuplicateRecordError struct {
	record string
}

func NewDuplicateRecordError(record string) *DuplicateRecordError {
	return &DuplicateRecordError{record: record}
}

func (e DuplicateRecordError) Error() string {
	return fmt.Sprintf("db already contains record: %s", e.record)
}

func (e DuplicateRecordError) GetRecord() string {
	return e.record
}

type UnexpectedError struct {
	error error
}

func NewUnexpectedError(err error) *UnexpectedError {
	return &UnexpectedError{error: err}
}

func (e UnexpectedError) Error() string {
	return fmt.Sprintf("unexpected db error: %v", e.error)
}
