package auth

import (
	"context"
	"gokomodo_test/internal/presentations"
)

type Resolver interface {
	Login(ctx context.Context, payload presentations.PayloadLogin) (presentations.Response, error)
}
