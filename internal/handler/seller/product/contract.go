package product

import "net/http"

type Resolver interface {
	AddProduct(w http.ResponseWriter, r *http.Request)
	ListingProducts(w http.ResponseWriter, r *http.Request)
}
