package dto

type Subscriptions struct {
	Contacts *[]string
}

func NewSubscriptions(contacts *[]string) *Subscriptions {
	return &Subscriptions{
		Contacts: contacts,
	}
}
