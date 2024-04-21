package repository

import (
	"context"
	"database/sql"
	"errors"
	"gokomodo_test/internal/entity"
	"log"
)

type userRoleRepository struct {
	db *sql.DB
}

func NewUserRoles(db *sql.DB) UserRoleRepositories {
	return &userRoleRepository{db: db}
}

func (u *userRoleRepository) FindUserRole(ctx context.Context, wheres entity.UserRoles) (*entity.UserRoles, error) {
	var (
		result entity.UserRoles
		err    error
	)

	q := `SELECT user_id, role_id FROM user_roles WHERE user_id = ? AND role_id = ? `
	rows, err := u.db.QueryContext(ctx, q, &wheres.UserId, &wheres.RoleId)
	if err != nil {
		log.Printf("got error when find user role %v", err)
		return nil, err
	}

	defer rows.Close()
	if rows.Next() {
		err = rows.Scan(&result.UserId, &result.RoleId)
		if err != nil {
			log.Printf("got error scan value %v", err)
			return nil, err
		}
		return &result, nil
	} else {
		return &result, errors.New("user role is not found")
	}
}
