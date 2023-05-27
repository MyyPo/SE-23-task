package repositories

import (
	"errors"

	"github.com/myypo/btcinform/internal/dto"
	simpdb "github.com/myypo/btcinform/pkg/simpDB"
)

type SubscriptionRepository interface {
	CreateSubscription(dto.Subscription) error
	GetAllSubscriptions() (*dto.Subscriptions, error)
}

type SubscriptionRepositoryImpl struct {
	dbEngine simpdb.SimpDB
}

func NewSubscriptionRepositoryImpl(dbEngine simpdb.SimpDB) *SubscriptionRepositoryImpl {
	return &SubscriptionRepositoryImpl{
		dbEngine: dbEngine,
	}
}

func (r *SubscriptionRepositoryImpl) CreateSubscription(
	subscription dto.Subscription,
) error {
	err := r.dbEngine.CreateOne(*subscription.GetContact())
	if err != nil {
		if !errors.Is(err, simpdb.DuplicateRecordError{}) {
			return NewUnexpectedRepoError(err)
		}

		return NewAlreadySubscribedError(*subscription.GetContact())
	}
	return nil
}

func (r *SubscriptionRepositoryImpl) GetAllSubscriptions() (*dto.Subscriptions, error) {
	subscriptionsDataModel, err := r.dbEngine.GetAll()
	if err != nil {
		return nil, NewUnexpectedRepoError(err)
	}

	subscriptions := dto.Subscriptions{
		Contacts: subscriptionsDataModel,
	}

	return &subscriptions, nil
}
