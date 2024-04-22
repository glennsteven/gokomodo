package order_product

import (
	"context"
	"gokomodo_test/internal/consts"
	"gokomodo_test/internal/entity"
	"gokomodo_test/internal/presentations"
	"log"
	"net/http"
	"time"
)

func (o *orderBuyerUseCase) OrderProduct(ctx context.Context, payload presentations.PayloadOrder, authId int) (presentations.Response, error) {
	findProduct, err := o.repoProduct.FindOne(ctx, payload.ProductId)
	if err != nil {
		log.Printf("failed find data product detail %v", err)
		return presentations.Response{
			Code:    http.StatusInternalServerError,
			Message: "Internal Server Error",
		}, err
	}

	if findProduct == nil {
		return presentations.Response{
			Code:    http.StatusNotFound,
			Message: "product not found",
		}, err
	}

	informationSeller, err := o.repoUser.FindId(ctx, findProduct.UserId)
	if err != nil {
		log.Printf("failed find data user %v", err)
		return presentations.Response{
			Code:    http.StatusInternalServerError,
			Message: "Internal Server Error",
		}, err
	}

	totalPrice := float64(payload.Quantity) * findProduct.Price

	saveOrder, err := o.repoOrder.Store(ctx, entity.Orders{
		UserId:                     authId,
		DeliverySourceAddress:      informationSeller.Address,
		DeliveryDestinationAddress: payload.DeliveryDestinationOrder,
		CreatedAt:                  time.Now(),
		UpdatedAt:                  time.Now(),
	})

	if err != nil {
		log.Printf("failed store data order, order failed %v", err)
		return presentations.Response{
			Code:    http.StatusInternalServerError,
			Message: "Internal Server Error",
		}, err
	}

	saveOrderDetail, err := o.repoOrderDetail.Store(ctx, entity.OrderDetails{
		ProductId:  findProduct.Id,
		OrderId:    saveOrder.Id,
		Quantity:   payload.Quantity,
		TotalPrice: totalPrice,
		Status:     consts.Pending,
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
	})

	if err != nil {
		log.Printf("failed store data order detail, order failed %v", err)
		return presentations.Response{
			Code:    http.StatusInternalServerError,
			Message: "Internal Server Error",
		}, err
	}

	return presentations.Response{
		Code:    http.StatusCreated,
		Message: "Your order will be processed immediately, please wait",
		Data:    saveOrderDetail,
	}, nil
}
