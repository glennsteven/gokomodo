package presentations

type PayloadLogin struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
