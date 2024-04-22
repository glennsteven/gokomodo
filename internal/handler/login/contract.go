package login

import "net/http"

type Resolver interface {
	Login(w http.ResponseWriter, r *http.Request)
}
