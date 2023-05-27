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
