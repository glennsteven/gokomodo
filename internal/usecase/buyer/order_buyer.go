package buyer

import (
	"encoding/json"
	"gokomodo_test/internal/consts"
	"gokomodo_test/internal/entity"
	"gokomodo_test/internal/helper"
	"gokomodo_test/internal/middleware"
	"gokomodo_test/internal/presentations"
	"gokomodo_test/internal/repository"
	"net/http"
	"time"
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
) OrderBuyerCase {
	return &orderBuyerUseCase{
		repoOrder:       repoOrder,
		repoOrderDetail: repoOrderDetail,
		repoProduct:     repoProduct,
		repoUser:        repoUser,
	}
}

func (o *orderBuyerUseCase) Create(w http.ResponseWriter, r *http.Request) {
	var (
		payload  presentations.PayloadOrder
		response presentations.Response
	)

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&payload); err != nil {
		response.Code = http.StatusBadRequest
		response.Message = err.Error()
		helper.ResponseJSON(w, http.StatusBadRequest, response)
		return
	}

	// Extract JWT token from request headers
	tokenString := r.Header.Get("Authorization")
	if tokenString == "" {
		response.Code = http.StatusUnauthorized
		response.Message = "Authorization header is missing"
		helper.ResponseJSON(w, http.StatusUnauthorized, response)
		return
	}

	// Decode JWT claims
	claims, err := middleware.ParseJWT(tokenString)
	if err != nil {
		response.Code = http.StatusUnauthorized
		response.Message = "Failed to decode token: " + err.Error()
		helper.ResponseJSON(w, http.StatusUnauthorized, response)
		return
	}

	// Type assert UserId field from claims as float64
	userIdFloat64, ok := claims["Id"].(float64)
	if !ok {
		response.Code = http.StatusUnauthorized
		response.Message = "UserId is not present in the token claims or not of type float64"
		helper.ResponseJSON(w, http.StatusUnauthorized, response)
		return
	}

	// Convert float64 to int
	userId := int(userIdFloat64)
	// Close the request body to prevent resource leaks
	defer r.Body.Close()

	findProduct, err := o.repoProduct.FindOne(r.Context(), payload.ProductId)
	if err != nil {
		response.Code = http.StatusInternalServerError
		response.Message = err.Error()
		helper.ResponseJSON(w, http.StatusInternalServerError, response)
		return
	}

	if findProduct == nil {
		response.Code = http.StatusUnauthorized
		response.Message = "product not found"
		helper.ResponseJSON(w, http.StatusNotFound, response)
		return
	}

	informationSeller, err := o.repoUser.FindId(r.Context(), findProduct.UserId)
	if err != nil {
		response.Code = http.StatusInternalServerError
		response.Message = err.Error()
		helper.ResponseJSON(w, http.StatusInternalServerError, response)
		return
	}

	totalPrice := float64(payload.Quantity) * findProduct.Price

	saveOrder, err := o.repoOrder.Store(r.Context(), entity.Orders{
		UserId:                     userId,
		DeliverySourceAddress:      informationSeller.Address,
		DeliveryDestinationAddress: payload.DeliveryDestinationOrder,
		CreatedAt:                  time.Now(),
		UpdatedAt:                  time.Now(),
	})
	if err != nil {
		response.Code = http.StatusInternalServerError
		response.Message = err.Error()
		helper.ResponseJSON(w, http.StatusInternalServerError, response)
		return
	}

	saveOrderDetail, err := o.repoOrderDetail.Store(r.Context(), entity.OrderDetails{
		ProductId:  findProduct.Id,
		OrderId:    saveOrder.Id,
		Quantity:   payload.Quantity,
		TotalPrice: totalPrice,
		Status:     consts.Pending,
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
	})

	if err != nil {
		response.Code = http.StatusInternalServerError
		response.Message = err.Error()
		helper.ResponseJSON(w, http.StatusInternalServerError, response)
		return
	}

	responses := presentations.Response{
		Code:    http.StatusCreated,
		Message: "Your order will be processed immediately, please wait",
		Data:    saveOrderDetail,
	}

	helper.ResponseJSON(w, http.StatusCreated, responses)
}
