package utils

import "errors"

func TestAreErrorsEqual(err error, targetErr error) bool {
	if err == nil && targetErr == nil {
		return true
	}
	if targetErr != nil {
		return errors.Is(err, targetErr)
	}
	return false
}
