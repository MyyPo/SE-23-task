package services

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"

	"github.com/myypo/btcinform/internal/constants"
	"github.com/myypo/btcinform/internal/dto"
	"github.com/myypo/btcinform/internal/dto/requests"
	"github.com/myypo/btcinform/internal/dto/responses"
	"github.com/myypo/btcinform/internal/repositories"
)

type ExchangeRateService interface {
	GetExchangeRate(request requests.GetRateRequest) (*responses.GetRateResponse, error)
	Subscribe(request requests.SubscribeRequest) (*responses.SubscribeResponse, error)
	SendEmails(request requests.SendEmailsRequest) (*responses.SendEmailsResponse, error)
}

type ExchangeRateServiceImpl struct {
	repo repositories.SubscriptionRepository
}

func NewExchangeRateServiceImpl(repo repositories.SubscriptionRepository) *ExchangeRateServiceImpl {
	return &ExchangeRateServiceImpl{
		repo: repo,
	}
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

func (s *ExchangeRateServiceImpl) Subscribe(
	request requests.SubscribeRequest,
) (*responses.SubscribeResponse, error) {
	subscription := dto.NewSubscriptionImpl(&request.Email)
	err := s.repo.CreateSubscription(subscription)
	if err != nil {
		if !errors.Is(err, repositories.AlreadySubscribedError{}) {
			return nil, NewInternalError()
		}
		return nil, NewDuplicateError(request.Email)
	}
	return responses.NewSubscribeResponse(constants.StatusSuccess), nil
}

func (s *ExchangeRateServiceImpl) SendEmails(
	request requests.SendEmailsRequest,
) (*responses.SendEmailsResponse, error) {
	return nil, nil
}
