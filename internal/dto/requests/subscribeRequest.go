package requests

type SubscribeRequest struct {
	Email string `json:"email" binding:"required,email"`
}
