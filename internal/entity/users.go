package entity

import "time"

type Users struct {
	ID        int64     `db:"id" json:"id"`
	RoleEnum  int       `db:"role_enum" json:"role_enum"`
	Email     string    `db:"email" json:"email"`
	Name      string    `db:"name" json:"name"`
	Password  string    `db:"password" json:"password"`
	Address   string    `db:"address" json:"address"`
	CreatedAt time.Time `db:"created_at" json:"created_at"`
	UpdatedAt time.Time `db:"updated_at" json:"updated_at"`
}
