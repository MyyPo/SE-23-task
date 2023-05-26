package responses

type SubscribeResponse struct {
	Status string `json:"status"`
}

func NewSubscribeResponse(status string) *SubscribeResponse {
	return &SubscribeResponse{
		Status: status,
	}
}
