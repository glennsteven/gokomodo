package entity

import "time"

type Products struct {
	Id          int       `db:"id" json:"id"`
	UserId      int       `db:"user_id" json:"user_id"`
	ProductName string    `db:"product_name" json:"product_name"`
	Description string    `db:"description" json:"description"`
	Price       float64   `db:"price" json:"price"`
	CreatedAt   time.Time `db:"created_at" json:"created_at"`
	UpdatedAt   time.Time `db:"updated_at" json:"updated_at"`
}

type ListingProduct struct {
	Id          int       `db:"id" json:"id"`
	UserId      int       `db:"user_id" json:"user_id"`
	FullName    string    `db:"full_name" json:"full_name"`
	Address     string    `db:"address" json:"address"`
	ProductName string    `db:"product_name" json:"product_name"`
	Description string    `db:"description" json:"description"`
	Price       float64   `db:"price" json:"price"`
	CreatedAt   time.Time `db:"created_at" json:"created_at"`
	UpdatedAt   time.Time `db:"updated_at" json:"updated_at"`
}
