package responses

type GetRateResponse struct {
	Rate string `json:"rate"`
}

func NewGetRateResponse(rate string) *GetRateResponse {
	return &GetRateResponse{
		Rate: rate,
	}
}
