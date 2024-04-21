package repository

import (
	"context"
	"database/sql"
	"errors"
	"gokomodo_test/internal/entity"
	"log"
)

type userRepository struct {
	db *sql.DB
}

func NewUsers(db *sql.DB) UserRepositories {
	return &userRepository{db: db}
}

func (u *userRepository) FindUser(ctx context.Context, wheres entity.Users) (*entity.Users, error) {
	var (
		result entity.Users
		err    error
	)

	q := `SELECT id, email, full_name, address, password FROM users WHERE email = ?`
	rows, err := u.db.QueryContext(ctx, q, &wheres.Email)
	if err != nil {
		log.Printf("got error when find username %v", err)
		return nil, err
	}

	defer rows.Close()
	if rows.Next() {
		err = rows.Scan(&result.Id, &result.Email, &result.FullName, &result.Address, &result.Password)
		if err != nil {
			log.Printf("got error scan value %v", err)
			return nil, err
		}
		return &result, nil
	} else {
		return &result, errors.New("user is not found")
	}
}

func (u *userRepository) FindId(ctx context.Context, id int) (*entity.Users, error) {
	var (
		result entity.Users
		err    error
	)

	q := `SELECT 
			id,
			email,
            full_name,
            address
		FROM users WHERE id = ?`

	rows, err := u.db.QueryContext(ctx, q, id)
	if err != nil {
		log.Printf("got error when find user %v", err)
		return nil, err
	}

	defer rows.Close()
	if rows.Next() {
		err = rows.Scan(&result.Id, &result.Email, &result.FullName, &result.Address)
		if err != nil {
			log.Printf("got error scan value %v", err)
			return nil, err
		}
		return &result, nil
	} else {
		return &result, errors.New("user is not found")
	}
}
