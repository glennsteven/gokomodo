package accept_order

import (
	"encoding/json"
	"gokomodo_test/internal/handler"
	"gokomodo_test/internal/helper"
	"gokomodo_test/internal/presentations"
	"net/http"
)

func (a *acceptProductService) AcceptOrder(w http.ResponseWriter, r *http.Request) {
	var (
		payload  presentations.PayloadAcceptOrder
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

	resultAccOrder, err := a.acceptProduct.AcceptOrder(r.Context(), presentations.PayloadAcceptOrder{
		OrderDetailId: payload.OrderDetailId,
		Status:        payload.Status,
	}, auth.Id)

	if err != nil {
		response.Code = http.StatusInternalServerError
		response.Message = "Internal Server Error"
		helper.ResponseJSON(w, http.StatusInternalServerError, response)
		return
	}

	helper.ResponseJSON(w, http.StatusBadRequest, resultAccOrder)
	return
}
