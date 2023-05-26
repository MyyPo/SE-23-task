package responses

type SendEmailsResponse struct {
	Status string `json:"status"`
}

func NewSendEmailsResponse(status string) *SendEmailsResponse {
	return &SendEmailsResponse{
		Status: status,
	}
}
