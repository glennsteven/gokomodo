package order_product

import (
	"context"
	"gokomodo_test/internal/presentations"
)

type Resolver interface {
	OrderProduct(ctx context.Context, payload presentations.PayloadOrder, authId int) (presentations.Response, error)
	ListOfProducts(ctx context.Context, payload presentations.Paging) (presentations.ResponseListingProduct, error)
	ListOrder(ctx context.Context, payload presentations.Paging, authId int) (presentations.ResponseListingOrderBuyer, error)
}
