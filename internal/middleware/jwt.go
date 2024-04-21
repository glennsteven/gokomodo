package middleware

import (
	"fmt"
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

func ParseJWT(tokenString string) (jwt.MapClaims, error) {
	bearerToken := strings.Split(tokenString, " ")
	token, err := jwt.Parse(bearerToken[1], func(token *jwt.Token) (interface{}, error) {
		// Validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		// Return the secret signing key (this could also come from an env variable or be loaded from a file)
		return consts.JWT_KEY, nil
	})

	if err != nil {
		return nil, err
	}

	// If the token is valid, claim the value
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	} else {
		return nil, err
	}
}
