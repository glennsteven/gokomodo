package order_product

import "net/http"

type Resolver interface {
	OrderProduct(w http.ResponseWriter, r *http.Request)
}
