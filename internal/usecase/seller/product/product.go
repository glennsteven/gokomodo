package product

import "gokomodo_test/internal/repository"

type sellerUseCase struct {
	repoProduct repository.ProductRepositories
}

func NewProductUseCase(
	repoProduct repository.ProductRepositories,
) Resolver {
	return &sellerUseCase{
		repoProduct: repoProduct,
	}
}
