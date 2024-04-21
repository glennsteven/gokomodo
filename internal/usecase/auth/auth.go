package auth

import (
	"encoding/json"
	"github.com/golang-jwt/jwt/v5"
	"gokomodo_test/internal/consts"
	"gokomodo_test/internal/entity"
	"gokomodo_test/internal/helper"
	"gokomodo_test/internal/presentations"
	"gokomodo_test/internal/repository"
	"io"
	"log"
	"net/http"
	"time"
)

type userAuth struct {
	user repository.UserRepositories
}

func NewUserAuth(
	user repository.UserRepositories,
) UserAuth {
	return &userAuth{
		user: user,
	}
}

func (u *userAuth) Login(w http.ResponseWriter, r *http.Request) {
	var (
		payload  presentations.PayloadLogin
		response presentations.Response
	)

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&payload); err != nil {
		response.Code = http.StatusBadRequest
		response.Message = err.Error()
		helper.ResponseJSON(w, http.StatusBadRequest, response)
		return
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(r.Body)

	findUser, err := u.user.FindUser(r.Context(), entity.Users{
		Email: payload.Email,
	})
	if err != nil {
		log.Printf("failed to find user %v", err)
		response.Code = http.StatusInternalServerError
		response.Message = err.Error()
		helper.ResponseJSON(w, http.StatusInternalServerError, response)
		return
	}

	if findUser == nil {
		response.Code = http.StatusUnauthorized
		response.Message = "invalid username or password"
		helper.ResponseJSON(w, http.StatusUnauthorized, response)
		return
	}

	err = compareHashPassword([]byte(findUser.Password), []byte(payload.Password))
	if err != nil {
		response.Code = http.StatusUnauthorized
		response.Message = "invalid username or password"
		helper.ResponseJSON(w, http.StatusUnauthorized, response)
		return
	}

	exp := time.Now().Add(time.Hour * 15)
	claims := &consts.JWTClaim{
		Email: findUser.Email,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "go_komodo_service",
			ExpiresAt: jwt.NewNumericDate(exp),
		},
	}

	tokenAlgorithm := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenAlgorithm.SignedString(consts.JWT_KEY)
	if err != nil {
		response.Code = http.StatusInternalServerError
		response.Message = err.Error()
		helper.ResponseJSON(w, http.StatusInternalServerError, response)
		return
	}

	response.Code = http.StatusOK
	response.Message = "token created"
	response.Data = map[string]string{"token": token}
	helper.ResponseJSON(w, http.StatusOK, response)
	return
}
