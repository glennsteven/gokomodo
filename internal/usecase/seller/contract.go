package seller

import "net/http"

type ProductCase interface {
	Create(w http.ResponseWriter, r *http.Request)
}
