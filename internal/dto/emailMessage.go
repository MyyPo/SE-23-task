package dto

import "fmt"

type EmailMessage struct {
	exchangeRate int
}

func NewEmailMessage(exchangeRate int) *EmailMessage {
	return &EmailMessage{
		exchangeRate: exchangeRate,
	}
}

func (m *EmailMessage) GetMessage() []byte {
	msgStr := fmt.Sprintf(
		"Hello, the current BTC to UAH exchange rate is: %d UAH for 1 bitcoin",
		m.exchangeRate,
	)
	return []byte(msgStr)
}
