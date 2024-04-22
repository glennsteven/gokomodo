package entity

import "time"

type Orders struct {
	Id                         int       `db:"id" json:"id"`
	UserId                     int       `db:"user_id" json:"user_id"`
	DeliverySourceAddress      string    `db:"delivery_source_address" json:"delivery_source_address"`
	DeliveryDestinationAddress string    `db:"delivery_destination_address" json:"delivery_destination_address"`
	CreatedAt                  time.Time `db:"created_at" json:"created_at"`
	UpdatedAt                  time.Time `db:"updated_at" json:"updated_at"`
}

type ListingOrderBuyer struct {
	Id                         int       `db:"id" json:"id"`
	ProductName                string    `db:"product_name" json:"product_name"`
	Description                string    `db:"description" json:"description"`
	IdSeller                   int       `db:"id_seller" json:"id_seller"`
	Seller                     string    `db:"seller" json:"seller"`
	IdBuyer                    string    `db:"id_buyer" json:"id_buyer"`
	Buyer                      string    `db:"buyer" json:"buyer"`
	Email                      string    `db:"email" json:"email"`
	DeliverySourceAddress      string    `db:"delivery_source_address" json:"delivery_source_address"`
	DeliveryDestinationAddress string    `db:"delivery_destination_address" json:"delivery_destination_address"`
	Quantity                   int       `db:"quantity" json:"quantity"`
	TotalPrice                 float64   `db:"total_price" json:"total_price"`
	Status                     int       `db:"status" json:"status"`
	CreatedAt                  time.Time `db:"created_at" json:"created_at"`
}
