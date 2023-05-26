package services

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/myypo/btcinform/internal/constants"
	"github.com/myypo/btcinform/internal/dto"
	"github.com/myypo/btcinform/internal/dto/requests"
	"github.com/myypo/btcinform/internal/dto/responses"
)

type ExchangeRateService interface {
	GetExchangeRate(request requests.GetRateRequest) (*responses.GetRateResponse, error)
}

type ExchangeRateServiceImpl struct{}

func NewExchangeRateServiceImpl() *ExchangeRateServiceImpl {
	return &ExchangeRateServiceImpl{}
}

func (s *ExchangeRateServiceImpl) GetExchangeRate(
	request requests.GetRateRequest,
) (*responses.GetRateResponse, error) {
	publicAPIResponse, err := http.Get(constants.ExchangeRateURL)
	if err != nil {
		return nil, err
	}
	defer publicAPIResponse.Body.Close()

	body, err := ioutil.ReadAll(publicAPIResponse.Body)
	if err != nil {
		return nil, err
	}

	var publicAPIData dto.PublicAPIRateResponse
	err = json.Unmarshal(body, &publicAPIData)
	if err != nil {
		return nil, err
	}

	return &responses.GetRateResponse{
		Rate: publicAPIData.BTC.UAH,
	}, nil
}
