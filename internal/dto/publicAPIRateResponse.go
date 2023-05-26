package dto

type PublicAPIRateResponse struct {
	BTC struct {
		UAH int `json:"uah"`
	} `json:"bitcoin"`
}
