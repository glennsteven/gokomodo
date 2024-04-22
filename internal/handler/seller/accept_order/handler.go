package accept_order

import (
	"gokomodo_test/internal/usecase/seller/accept_order"
)

type acceptProductService struct {
	acceptProduct accept_order.Resolver
}

func NewAcceptProductService(
	acceptProduct accept_order.Resolver,
) Resolver {
	return &acceptProductService{
		acceptProduct: acceptProduct,
	}
}
