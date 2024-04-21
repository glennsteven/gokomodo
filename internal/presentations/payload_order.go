package presentations

type PayloadOrder struct {
	ProductId                int    `json:"product_id"`
	Quantity                 int    `json:"quantity"`
	DeliveryDestinationOrder string `json:"delivery_destination_order"`
}
