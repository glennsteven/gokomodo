package accept_order

import "gokomodo_test/internal/repository"

type acceptOrderUseCase struct {
	repoOrderDetail repository.OrderDetailsRepositories
	repoProduct     repository.ProductRepositories
}

func NewAcceptOrderUseCase(
	repoOrderDetail repository.OrderDetailsRepositories,
	repoProduct repository.ProductRepositories,
) Resolver {
	return &acceptOrderUseCase{
		repoOrderDetail: repoOrderDetail,
		repoProduct:     repoProduct,
	}
}
