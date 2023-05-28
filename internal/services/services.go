package services

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/smtp"

	"github.com/myypo/btcinform/internal/config"
	"github.com/myypo/btcinform/internal/constants"
	"github.com/myypo/btcinform/internal/dto"
	"github.com/myypo/btcinform/internal/dto/requests"
	"github.com/myypo/btcinform/internal/dto/responses"
	"github.com/myypo/btcinform/internal/repositories"
	"github.com/myypo/btcinform/internal/utils"
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

	rate := utils.FormatNumberWithCommas(publicAPIResponse.BTC.UAH)

	return &responses.GetRateResponse{
		Rate: rate,
	}, nil
}

func (s *ExchangeRateServiceImpl) Subscribe(
	request requests.SubscribeRequest,
) (*responses.SubscribeResponse, error) {
	subscription := dto.NewSubscriptionImpl(&request.Email)
	err := s.repo.CreateSubscription(subscription)
	if err != nil {
		if !errors.Is(err, repositories.AlreadySubscribedError{}) {
			log.Println("repository Subscribe InternalError: ", err)
			return nil, NewInternalError()
		}
		log.Println("repository Subscribe DuplicateError for: ", subscription.GetContact())
		return nil, NewDuplicateError(request.Email)
	}
	return responses.NewSubscribeResponse(constants.StatusSuccess), nil
}

func (s *ExchangeRateServiceImpl) SendEmails(
	request requests.SendEmailsRequest,
) (*responses.SendEmailsResponse, error) {
	subscribers, err := s.repo.GetAllSubscriptions()
	if err != nil {
		log.Println("repository GetAllSubscriptions InternalError: ", err)
		return nil, NewInternalError()
	}
	log.Println("fetched contacts: ", *subscribers.Contacts)

	exchangeRate, err := s.callExchangeRateAPI()
	if err != nil {
		return nil, err
	}
	formattedExchangeRate := utils.FormatNumberWithCommas(exchangeRate.BTC.UAH)

	err = smtp.SendMail(
		fmt.Sprintf("%s:%s", *s.config.GetSMTPHost(), *s.config.GetSMTPPort()),
		s.smtpAuth,
		*s.config.GetEmailUsername(),
		*subscribers.Contacts,
		dto.NewEmailMessage(&formattedExchangeRate, subscribers.Contacts).GetMessage(),
	)

	if err != nil {
		log.Println("service SendEmails InternalError when sending: ", err)
		return nil, NewInternalError()
	}

	return responses.NewSendEmailsResponse(constants.StatusSuccess), nil
}

func (s *ExchangeRateServiceImpl) callExchangeRateAPI() (*dto.PublicAPIRateResponse, error) {
	publicAPIResponse, err := http.Get(constants.ExchangeRateURL)
	if err != nil {
		log.Println("service callExchangeRateAPI InternalError when getting rate: ", err)
		return nil, NewInternalError()
	}
	defer publicAPIResponse.Body.Close()

	body, err := io.ReadAll(publicAPIResponse.Body)
	if err != nil {
		log.Println("service callExchangeRateAPI InternalError when reading body: ", err)
		return nil, NewInternalError()
	}

	var publicAPIData dto.PublicAPIRateResponse
	err = json.Unmarshal(body, &publicAPIData)
	if err != nil {
		log.Println("service callExchangeRateAPI InternalError when Unmarshalling: ", err)
		return nil, NewInternalError()
	}

	return &publicAPIData, nil
}
