package entity

import "time"

type OrderDetails struct {
	Id         int       `db:"id" json:"id"`
	ProductId  int       `db:"product_id" json:"product_id"`
	OrderId    int       `db:"order_id" json:"order_id"`
	Quantity   int       `db:"quantity" json:"quantity"`
	TotalPrice float64   `db:"total_price" json:"total_price"`
	Status     int       `db:"status" json:"status"`
	CreatedAt  time.Time `db:"created_at" json:"created_at"`
	UpdatedAt  time.Time `db:"updated_at" json:"updated_at"`
}
