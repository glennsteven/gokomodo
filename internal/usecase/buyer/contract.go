package buyer

import "net/http"

type OrderBuyerCase interface {
	Create(w http.ResponseWriter, r *http.Request)
}
