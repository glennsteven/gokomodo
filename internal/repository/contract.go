package repository

import (
	"context"
	"gokomodo_test/internal/entity"
)

type UserRepositories interface {
	FindUser(ctx context.Context, wheres entity.Users) (*entity.Users, error)
	FindId(ctx context.Context, id int) (*entity.Users, error)
}

type UserRoleRepositories interface {
	FindUserRole(ctx context.Context, wheres entity.UserRoles) (*entity.UserRoles, error)
}

type ProductRepositories interface {
	Store(ctx context.Context, payload entity.Products) (*entity.Products, error)
	FindOne(ctx context.Context, id int) (*entity.Products, error)
}

type OrderRepositories interface {
	Store(ctx context.Context, payload entity.Orders) (*entity.Orders, error)
}

type OrderDetailsRepositories interface {
	Store(ctx context.Context, payload entity.OrderDetails) (*entity.OrderDetails, error)
}
