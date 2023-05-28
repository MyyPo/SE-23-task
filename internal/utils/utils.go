package utils

import (
	"errors"

	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

func TestAreErrorsEqual(err error, targetErr error) bool {
	if err == nil && targetErr == nil {
		return true
	}
	if targetErr != nil {
		return errors.Is(err, targetErr)
	}
	return false
}

func FormatNumberWithCommas(number int) string {
	printer := message.NewPrinter(language.English)
	return printer.Sprint(number)
}
