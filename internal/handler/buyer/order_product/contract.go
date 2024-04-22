package order_product

import "net/http"

type Resolver interface {
	OrderProduct(w http.ResponseWriter, r *http.Request)
	ListOfProducts(w http.ResponseWriter, r *http.Request)
	ListOrders(w http.ResponseWriter, r *http.Request)
}
