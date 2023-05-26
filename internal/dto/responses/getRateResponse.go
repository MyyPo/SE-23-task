package responses

type GetRateResponse struct {
	Rate int `json:"rate"`
}

func NewGetRateResponse(rate int) *GetRateResponse {
	return &GetRateResponse{
		Rate: rate,
	}
}
