package services

import (
	"testing"

	"github.com/myypo/btcinform/internal/dto/requests"
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
			NewExchangeRateServiceImpl(),
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
