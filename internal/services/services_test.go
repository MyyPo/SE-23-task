package services

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/myypo/btcinform/internal/constants"
	"github.com/myypo/btcinform/internal/dto"
	"github.com/myypo/btcinform/internal/dto/requests"
	"github.com/myypo/btcinform/internal/dto/responses"
	"github.com/myypo/btcinform/internal/repositories"
	"github.com/myypo/btcinform/internal/utils"
)

func TestExchangeRateServiceImpl_GetExchangeRate(t *testing.T) {
	type args struct {
		request requests.GetRateRequest
	}

	tests := []struct {
		name    string
		s       *ExchangeRateServiceImpl
		args    args
		wantErr bool
	}{
		{
			"Valid rate request",
			NewExchangeRateServiceImpl(&testConfig{}, &subscriptionRepositoryMock{}),
			args{request: requests.GetRateRequest{}},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.s.GetExchangeRate(tt.args.request)
			if (err != nil) != tt.wantErr {
				t.Errorf(
					"ExchangeRateServiceImpl.GetExchangeRate() error = %v, wantErr %v",
					err,
					tt.wantErr,
				)
				return
			}
			if got.Rate == 0 {
				t.Errorf("ExchangeRateServiceImpl.GetExchangeRate() = didn't return rate")
			}
			t.Logf("ExchangeRateServiceImpl.GetExchangeRate() returned %d", got.Rate)
		})
	}
}

func TestExchangeRateServiceImpl_Subscribe(t *testing.T) {
	tests := []struct {
		name    string
		s       *ExchangeRateServiceImpl
		want    *responses.SubscribeResponse
		wantErr error
	}{
		{
			name: "successful subscription",
			s: NewExchangeRateServiceImpl(&testConfig{}, &subscriptionRepositoryMock{
				createSubscriptionFunc: func() error {
					return nil
				},
			}),
			want:    responses.NewSubscribeResponse(constants.StatusSuccess),
			wantErr: nil,
		},
		{
			name: "attempt to subscribe already subscribed email",
			s: NewExchangeRateServiceImpl(&testConfig{}, &subscriptionRepositoryMock{
				createSubscriptionFunc: func() error {
					return repositories.NewAlreadySubscribedError(testEmail)
				},
			}),
			want:    nil,
			wantErr: DuplicateError{},
		},
		{
			name: "encounter unexpected error",
			s: NewExchangeRateServiceImpl(&testConfig{}, &subscriptionRepositoryMock{
				createSubscriptionFunc: func() error {
					return repositories.NewUnexpectedRepoError(fmt.Errorf("disaster"))
				},
			}),
			want:    nil,
			wantErr: InternalError{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.s.Subscribe(requests.SubscribeRequest{})
			if !utils.TestAreErrorsEqual(err, tt.wantErr) {
				t.Errorf(
					"ExchangeRateServiceImpl.Subscribe() error = %v, wantErr %v",
					err,
					tt.wantErr,
				)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ExchangeRateServiceImpl.Subscribe() = %v, want %v", got, tt.want)
			}
		})
	}
}

const (
	testEmail = "test@gmail.com"
)

type subscriptionRepositoryMock struct {
	createSubscriptionFunc  func() error
	getAllSubscriptionsFunc func() (*dto.Subscriptions, error)
}

func (r *subscriptionRepositoryMock) CreateSubscription(dto.Subscription) error {
	return r.createSubscriptionFunc()
}

func (r *subscriptionRepositoryMock) GetAllSubscriptions() (*dto.Subscriptions, error) {
	return r.getAllSubscriptionsFunc()
}

type testConfig struct{}

func (c *testConfig) GetDBPath() *string {
	dbPath := "dbPath"
	return &dbPath
}

func (c *testConfig) GetSMTPHost() *string {
	host := "host"
	return &host
}

func (c *testConfig) GetSMTPPort() *string {
	port := "port"
	return &port
}

func (c *testConfig) GetEmailUsername() *string {
	username := "username"
	return &username
}

func (c *testConfig) GetEmailPassword() *string {
	password := "password"
	return &password
}
