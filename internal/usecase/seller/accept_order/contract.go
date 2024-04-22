package accept_order

import (
	"context"
	"gokomodo_test/internal/presentations"
)

type Resolver interface {
	AcceptOrder(ctx context.Context, payload presentations.PayloadAcceptOrder, authId int) (presentations.Response, error)
}
