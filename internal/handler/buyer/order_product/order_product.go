package order_product

import (
	"encoding/json"
	"gokomodo_test/internal/handler"
	"gokomodo_test/internal/helper"
	"gokomodo_test/internal/presentations"
	"net/http"
)

func (o *orderProductService) OrderProduct(w http.ResponseWriter, r *http.Request) {
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

	tokenString := r.Header.Get("Authorization")
	if tokenString == "" {
		response.Code = http.StatusUnauthorized
		response.Message = "unauthorized"
		helper.ResponseJSON(w, http.StatusUnauthorized, response)
		return
	}

	auth, err := handler.DecodeJWT(tokenString)
	if err != nil {
		response.Code = http.StatusUnauthorized
		response.Message = "unauthorized"
		helper.ResponseJSON(w, http.StatusUnauthorized, response)
		return
	}

	// Close the request body to prevent resource leaks
	defer r.Body.Close()

	resultProduct, err := o.orderProductUseCase.OrderProduct(r.Context(), presentations.PayloadOrder{
		ProductId:                payload.ProductId,
		Quantity:                 payload.Quantity,
		DeliveryDestinationOrder: payload.DeliveryDestinationOrder,
	}, auth.Id)

	if err != nil {
		response.Code = http.StatusInternalServerError
		response.Message = "Internal Server Error"
		helper.ResponseJSON(w, http.StatusInternalServerError, response)
		return
	}

	helper.ResponseJSON(w, http.StatusBadRequest, resultProduct)
	return
}
