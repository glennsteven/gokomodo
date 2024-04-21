package repository

import (
	"context"
	"gokomodo_test/internal/entity"
)

type UserRepositories interface {
	FindUser(ctx context.Context, wheres entity.Users) (*entity.Users, error)
}
