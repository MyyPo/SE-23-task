package dto

type Subscription interface {
	GetContact() *string
}

type SubscriptionImpl struct {
	Email *string
}

func NewSubscriptionImpl(email *string) *SubscriptionImpl {
	return &SubscriptionImpl{
		Email: email,
	}
}

func (s *SubscriptionImpl) GetContact() *string {
	return s.Email
}
