package product

import (
	"context"
	"gokomodo_test/internal/presentations"
)

type Resolver interface {
	CreateProduct(ctx context.Context, payload presentations.PayloadProduct) (presentations.Response, error)
}
