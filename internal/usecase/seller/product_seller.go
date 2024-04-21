package seller

import (
	"encoding/json"
	"gokomodo_test/internal/entity"
	"gokomodo_test/internal/helper"
	"gokomodo_test/internal/middleware"
	"gokomodo_test/internal/presentations"
	"gokomodo_test/internal/repository"
	"log"
	"net/http"
	"time"
)

type productUseCase struct {
	repoProduct repository.ProductRepositories
}

func NewProductUseCase(
	repoProduct repository.ProductRepositories,
) ProductCase {
	return &productUseCase{repoProduct: repoProduct}
}

func (p *productUseCase) Create(w http.ResponseWriter, r *http.Request) {
	var (
		payload  presentations.PayloadProduct
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

	valueProduct := entity.Products{
		UserId:      userId,
		ProductName: payload.Name,
		Description: payload.Description,
		Price:       payload.Price,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
	saveProducts, err := p.repoProduct.Store(r.Context(), valueProduct)
	if err != nil {
		log.Printf("failed save new product %v", err)
		response.Code = http.StatusInternalServerError
		response.Message = err.Error()
		helper.ResponseJSON(w, http.StatusInternalServerError, response)
		return
	}

	responses := presentations.Response{
		Code:    http.StatusCreated,
		Message: "success create product",
		Data:    saveProducts,
	}

	helper.ResponseJSON(w, http.StatusCreated, responses)
}
