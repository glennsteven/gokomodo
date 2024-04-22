package order_product

import (
	"gokomodo_test/internal/usecase/buyer/order_product"
)

type orderProductService struct {
	orderProductUseCase order_product.Resolver
}

func NewHandlerOrderProductService(
	orderProductUseCase order_product.Resolver,
) Resolver {
	return &orderProductService{
		orderProductUseCase: orderProductUseCase,
	}
}
