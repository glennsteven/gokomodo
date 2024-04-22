package order_product

import (
	"gokomodo_test/internal/helper"
	"gokomodo_test/internal/presentations"
	"net/http"
	"strconv"
)

func (o *orderProductService) ListOfProducts(w http.ResponseWriter, r *http.Request) {
	var (
		payload  presentations.Paging
		response presentations.ResponseListingProduct
	)

	tokenString := r.Header.Get("Authorization")
	if tokenString == "" {
		response.Code = http.StatusUnauthorized
		response.Message = "unauthorized"
		helper.ResponseJSON(w, http.StatusUnauthorized, response)
		return
	}

	page := r.URL.Query().Get("page")
	limit := r.URL.Query().Get("limit")
	pg, _ := strconv.Atoi(page)
	lm, _ := strconv.Atoi(limit)

	payload.Page = pg
	payload.Limit = lm

	resultProducts, err := o.orderProductUseCase.ListOfProducts(r.Context(), presentations.Paging{
		Limit: payload.Limit,
		Page:  payload.Page,
	})
	if err != nil {
		response.Code = http.StatusInternalServerError
		response.Message = "Internal Server Error"
		helper.ResponseJSON(w, http.StatusInternalServerError, response)
		return
	}

	helper.ResponseJSON(w, http.StatusOK, resultProducts)
	return
}
