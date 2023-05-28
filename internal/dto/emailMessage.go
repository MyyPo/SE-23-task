package dto

import (
	"fmt"
	"strings"
)

type EmailMessage struct {
	exchangeRate *string
	receivers    *[]string
}

func NewEmailMessage(exchangeRate *string, receivers *[]string) *EmailMessage {
	return &EmailMessage{
		exchangeRate: exchangeRate,
		receivers:    receivers,
	}
}

func (m *EmailMessage) GetMessage() []byte {
	from := "From: myypo@github.com\r\n"
	to := "To: " + strings.Join(*m.receivers, ", ") + "\r\n"
	subject := "Subject: BTC to UAH rate\r\n\r\n"
	body := fmt.Sprintf(
		"Hello, the current BTC to UAH exchange rate is: %s UAH for 1 bitcoin.\r\n",
		*m.exchangeRate,
	)
	msgString := from + to + subject + body
	return []byte(msgString)
}
