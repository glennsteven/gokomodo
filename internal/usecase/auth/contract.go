package auth

import "net/http"

type UserAuth interface {
	Login(w http.ResponseWriter, r *http.Request)
}
