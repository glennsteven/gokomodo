package entity

import "time"

type Roles struct {
	Id           int       `db:"id" json:"id"`
	Roles        string    `db:"roles" json:"roles"`
	RoleParentId int       `db:"role_parent_id" json:"role_parent_id"`
	CreatedAt    time.Time `db:"created_at" json:"created_at"`
	UpdatedAt    time.Time `db:"updated_at" json:"updated_at"`
}
