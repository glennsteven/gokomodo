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

type ListingOrderDetails struct {
	Id            int       `db:"id" json:"id"`
	SellerId      int       `db:"seller_id" json:"seller_id"`
	Seller        string    `db:"seller" json:"seller"`
	SellerAddress string    `db:"seller_address" json:"seller_address"`
	BuyerId       int       `db:"buyer_id" json:"buyer_id"`
	Buyer         string    `db:"buyer" json:"buyer"`
	BuyerAddress  string    `db:"buyer_address" json:"buyer_address"`
	Quantity      int       `db:"quantity" json:"quantity"`
	TotalPrice    float64   `db:"total_price" json:"total_price"`
	Status        int       `db:"status" json:"status"`
	ProductName   string    `db:"product_name" json:"product_name"`
	Description   string    `db:"description" json:"description"`
	Price         float64   `db:"price" json:"price"`
	CreatedAt     time.Time `db:"created_at" json:"created_at"`
	UpdatedAt     time.Time `db:"updated_at" json:"updated_at"`
}
