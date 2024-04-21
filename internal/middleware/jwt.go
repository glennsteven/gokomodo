package middleware

import (
	"github.com/golang-jwt/jwt/v5"
	"gokomodo_test/internal/consts"
	"gokomodo_test/internal/helper"
	"gokomodo_test/internal/presentations"
	"net/http"
	"strings"
)

func JWTMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var response presentations.Response
		authorizationHeader := r.Header.Get("Authorization")
		if authorizationHeader != "" {
			bearerToken := strings.Split(authorizationHeader, " ")
			if len(bearerToken) == 2 {
				token, err := jwt.Parse(bearerToken[1], func(token *jwt.Token) (interface{}, error) {
					return consts.JWT_KEY, nil
				})

				if err != nil {
					response.Code = http.StatusUnauthorized
					response.Message = "unauthorized"
					helper.ResponseJSON(w, http.StatusUnauthorized, response)
					return
				}

				if !token.Valid {
					response.Code = http.StatusUnauthorized
					response.Message = "unauthorized"
					helper.ResponseJSON(w, http.StatusUnauthorized, response)
					return
				}
				next.ServeHTTP(w, r)
			}
		} else {
			response.Code = http.StatusUnauthorized
			response.Message = "authorization header is required"
			helper.ResponseJSON(w, http.StatusUnauthorized, response)
			return
		}
	})
}
