package login

import (
	"encoding/json"
	"gokomodo_test/internal/helper"
	"gokomodo_test/internal/presentations"
	"io"
	"net/http"
)

func (s *loginService) Login(w http.ResponseWriter, r *http.Request) {
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
		_ = Body.Close()
	}(r.Body)

	resultLogin, err := s.authUseCase.Login(r.Context(), presentations.PayloadLogin{
		Email:    payload.Email,
		Password: payload.Password,
		RoleId:   payload.RoleId,
	})

	if err != nil {
		response.Code = http.StatusInternalServerError
		response.Message = "Internal Server Error"
		helper.ResponseJSON(w, http.StatusInternalServerError, response)
		return
	}

	helper.ResponseJSON(w, http.StatusBadRequest, resultLogin)
	return
}
