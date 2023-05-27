package repositories

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/myypo/btcinform/internal/dto"
	"github.com/myypo/btcinform/internal/utils"
	simpdb "github.com/myypo/btcinform/pkg/simpDB"
)

func TestSubscriptionRepositoryImpl_CreateSubscription(t *testing.T) {
	type args struct {
		subscription dto.Subscription
	}
	tests := []struct {
		name    string
		r       *SubscriptionRepositoryImpl
		args    args
		wantErr error
	}{
		{
			name: "create valid subscription",
			r: NewSubscriptionRepositoryImpl(
				&simpDBMock{
					createOneFunc: func() error {
						return nil
					},
				},
			),
			args: args{
				subscription: dto.NewSubscriptionImpl(&testSubContact),
			},
			wantErr: nil,
		},
		{
			name: "try to create invalid subscription (for the same contact)",
			r: NewSubscriptionRepositoryImpl(
				&simpDBMock{
					createOneFunc: func() error {
						return simpdb.NewDuplicateRecordError(testSubContact)
					},
				},
			),
			args: args{
				subscription: dto.NewSubscriptionImpl(&testSubContact),
			},
			wantErr: AlreadySubscribedError{},
		},
		{
			name: "handle unexpected error from db",
			r: NewSubscriptionRepositoryImpl(
				&simpDBMock{
					createOneFunc: func() error {
						return fmt.Errorf("unexpected error from db")
					},
				},
			),
			args: args{
				subscription: dto.NewSubscriptionImpl(&testSubContact),
			},
			wantErr: UnexpectedRepoError{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.r.CreateSubscription(tt.args.subscription); !utils.TestAreErrorsEqual(
				err,
				tt.wantErr,
			) {
				t.Errorf(
					"SubscriptionRepositoryImpl.CreateSubscription() error = %v, wantErr %v",
					err,
					tt.wantErr,
				)
			}
		})
	}
}

func TestSubscriptionRepositoryImpl_GetAllSubscriptions(t *testing.T) {
	tests := []struct {
		name    string
		r       *SubscriptionRepositoryImpl
		want    *dto.Subscriptions
		wantErr error
	}{
		{
			name: "get subscriptions without errors",
			r: NewSubscriptionRepositoryImpl(&simpDBMock{
				getAllFunc: func() (*[]string, error) {
					return &[]string{testSubContact, "something@outlook.com", "test@test.com"}, nil
				},
			}),
			want: &dto.Subscriptions{
				Contacts: &[]string{testSubContact, "something@outlook.com", "test@test.com"},
			},
			wantErr: nil,
		},
		{
			name: "handle unexpected error from db",
			r: NewSubscriptionRepositoryImpl(&simpDBMock{
				getAllFunc: func() (*[]string, error) {
					return nil, simpdb.NewUnexpectedError(fmt.Errorf("something went wrong"))
				},
			}),
			want:    nil,
			wantErr: UnexpectedRepoError{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.r.GetAllSubscriptions()
			if !utils.TestAreErrorsEqual(err, tt.wantErr) {
				t.Errorf(
					"SubscriptionRepositoryImpl.GetAllSubscriptions() error = %v, wantErr %v",
					err,
					tt.wantErr,
				)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf(
					"SubscriptionRepositoryImpl.GetAllSubscriptions() = %v, want %v",
					got,
					tt.want,
				)
			}
		})
	}
}

var testSubContact string = "email@gmail.com"

type simpDBMock struct {
	createOneFunc func() error
	getAllFunc    func() (*[]string, error)
}

func (m *simpDBMock) CreateOne(newRecord string) error {
	return m.createOneFunc()
}

func (m *simpDBMock) GetAll() (*[]string, error) {
	return m.getAllFunc()
}
