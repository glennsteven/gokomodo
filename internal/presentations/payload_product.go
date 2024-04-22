package presentations

type PayloadProduct struct {
	UserId      int     `json:"user_id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
}
