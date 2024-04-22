package presentations

type PayloadAcceptOrder struct {
	OrderDetailId int `json:"order_detail_id"`
	Status        int `json:"status"`
}
