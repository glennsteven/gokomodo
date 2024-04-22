package order_product

import (
	"context"
	"gokomodo_test/internal/presentations"
)

type Resolver interface {
	OrderProduct(ctx context.Context, payload presentations.PayloadOrder, authId int) (presentations.Response, error)
}
