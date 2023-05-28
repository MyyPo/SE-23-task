package services

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/smtp"

	"github.com/myypo/btcinform/internal/config"
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
	config   config.Config
	repo     repositories.SubscriptionRepository
	smtpAuth smtp.Auth
}

func NewExchangeRateServiceImpl(
	config config.Config,
	repo repositories.SubscriptionRepository,
) *ExchangeRateServiceImpl {
	smtpAuth := smtp.PlainAuth(
		"",
		*config.GetEmailUsername(),
		*config.GetEmailPassword(),
		*config.GetSMTPHost(),
	)

	return &ExchangeRateServiceImpl{
		config:   config,
		repo:     repo,
		smtpAuth: smtpAuth,
	}
}

func (s *ExchangeRateServiceImpl) GetExchangeRate(
	request requests.GetRateRequest,
) (*responses.GetRateResponse, error) {
	publicAPIResponse, err := s.callExchangeRateAPI()
	if err != nil {
		return nil, err
	}

	return &responses.GetRateResponse{
		Rate: publicAPIResponse.BTC.UAH,
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
	subscribers, err := s.repo.GetAllSubscriptions()
	if err != nil {
		return nil, NewInternalError()
	}

	exchangeRate, err := s.callExchangeRateAPI()
	if err != nil {
		return nil, err
	}

	err = smtp.SendMail(
		fmt.Sprintf("%s:%s", *s.config.GetSMTPHost(), *s.config.GetSMTPPort()),
		s.smtpAuth,
		*s.config.GetEmailUsername(),
		*subscribers.Contacts,
		dto.NewEmailMessage(exchangeRate.BTC.UAH).GetMessage(),
	)
	if err != nil {
		return nil, NewInternalError()
	}

	return responses.NewSendEmailsResponse(constants.StatusSuccess), nil
}

func (s *ExchangeRateServiceImpl) callExchangeRateAPI() (*dto.PublicAPIRateResponse, error) {
	publicAPIResponse, err := http.Get(constants.ExchangeRateURL)
	if err != nil {
		return nil, NewInternalError()
	}
	defer publicAPIResponse.Body.Close()

	body, err := ioutil.ReadAll(publicAPIResponse.Body)
	if err != nil {
		return nil, NewInternalError()
	}

	var publicAPIData dto.PublicAPIRateResponse
	err = json.Unmarshal(body, &publicAPIData)
	if err != nil {
		return nil, NewInternalError()
	}

	return &publicAPIData, nil
}
