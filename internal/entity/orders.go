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
