package product

import (
	"gokomodo_test/internal/usecase/seller/product"
)

type productService struct {
	productSeller product.Resolver
}

func NewHandlerSeller(
	productSeller product.Resolver,
) Resolver {
	return &productService{
		productSeller: productSeller,
	}
}
