package entity

import "time"

type UserRoles struct {
	Id        int       `db:"id" json:"id"`
	UserId    int       `db:"user_id" json:"user_id"`
	RoleId    int       `db:"role_id" json:"role_id"`
	CreatedAt time.Time `db:"created_at" json:"created_at"`
	UpdatedAt time.Time `db:"updated_at" json:"updated_at"`
}
