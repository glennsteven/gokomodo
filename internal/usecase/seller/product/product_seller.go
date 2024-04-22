package product

import (
	"context"
	"gokomodo_test/internal/entity"
	"gokomodo_test/internal/presentations"
	"log"
	"net/http"
	"time"
)

func (p *sellerUseCase) CreateProduct(ctx context.Context, payload presentations.PayloadProduct) (presentations.Response, error) {
	valueProduct := entity.Products{
		UserId:      payload.UserId,
		ProductName: payload.Name,
		Description: payload.Description,
		Price:       payload.Price,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
	saveProducts, err := p.repoProduct.Store(ctx, valueProduct)
	if err != nil {
		log.Printf("failed save new product %v", err)
		return presentations.Response{
			Code:    http.StatusInternalServerError,
			Message: "Internal Server Error",
		}, err
	}

	return presentations.Response{
		Code:    http.StatusCreated,
		Message: "success create product",
		Data:    saveProducts,
	}, nil
}
