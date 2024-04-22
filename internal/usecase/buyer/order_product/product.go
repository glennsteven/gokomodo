package order_product

import (
	"gokomodo_test/internal/repository"
)

type orderBuyerUseCase struct {
	repoOrder       repository.OrderRepositories
	repoOrderDetail repository.OrderDetailsRepositories
	repoProduct     repository.ProductRepositories
	repoUser        repository.UserRepositories
}

func NewOrderBuyerUseCase(
	repoOrder repository.OrderRepositories,
	repoOrderDetail repository.OrderDetailsRepositories,
	repoProduct repository.ProductRepositories,
	repoUser repository.UserRepositories,
) Resolver {
	return &orderBuyerUseCase{
		repoOrder:       repoOrder,
		repoOrderDetail: repoOrderDetail,
		repoProduct:     repoProduct,
		repoUser:        repoUser,
	}
}
