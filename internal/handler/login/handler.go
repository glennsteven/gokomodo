package login

import (
	"gokomodo_test/internal/usecase/auth"
)

type loginService struct {
	authUseCase auth.Resolver
}

func NewHandlerLogin(
	authUseCase auth.Resolver,
) Resolver {
	return &loginService{
		authUseCase: authUseCase,
	}
}
