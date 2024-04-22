package accept_order

import (
	"context"
	"gokomodo_test/internal/entity"
	"gokomodo_test/internal/presentations"
	"gokomodo_test/internal/repository"
	"log"
	"net/http"
	"time"
)

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

func (p *acceptOrderUseCase) AcceptOrder(ctx context.Context, payload presentations.PayloadAcceptOrder, authId int) (presentations.Response, error) {
	findDetailOrder, err := p.repoOrderDetail.FindOne(ctx, payload.OrderDetailId)
	if err != nil {
		log.Printf("failed find data order detail %v", err)
		return presentations.Response{
			Code:    http.StatusInternalServerError,
			Message: "Internal Server Error",
		}, err
	}

	if findDetailOrder == nil {
		return presentations.Response{
			Code:    http.StatusNotFound,
			Message: "data order detail not found",
		}, err
	}

	findProduct, err := p.repoProduct.FindOne(ctx, findDetailOrder.ProductId)
	if err != nil {
		log.Printf("failed find data product %v", err)
		return presentations.Response{
			Code:    http.StatusInternalServerError,
			Message: "Internal Server Error",
		}, err
	}

	if findProduct.UserId != authId {
		return presentations.Response{
			Code:    http.StatusUnprocessableEntity,
			Message: "You don't have access for accept this order",
		}, err
	}

	err = p.repoOrderDetail.AcceptOrderBuyer(ctx, entity.OrderDetails{
		Id:        payload.OrderDetailId,
		Status:    payload.Status,
		UpdatedAt: time.Now(),
	})

	if err != nil {
		return presentations.Response{
			Code:    http.StatusInternalServerError,
			Message: "Internal Server Error",
		}, err
	}

	return presentations.Response{
		Code:    http.StatusOK,
		Message: "orders in the packaging process",
	}, err
}
