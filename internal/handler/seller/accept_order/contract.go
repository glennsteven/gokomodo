package accept_order

import "net/http"

type Resolver interface {
	AcceptOrder(w http.ResponseWriter, r *http.Request)
}
