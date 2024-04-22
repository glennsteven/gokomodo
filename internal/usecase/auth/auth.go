package auth

import (
	"context"
	"github.com/golang-jwt/jwt/v5"
	"gokomodo_test/internal/consts"
	"gokomodo_test/internal/entity"
	"gokomodo_test/internal/presentations"
	"gokomodo_test/internal/repository"
	"log"
	"net/http"
	"time"
)

type userAuth struct {
	user     repository.UserRepositories
	userRole repository.UserRoleRepositories
}

func NewUserAuth(
	user repository.UserRepositories,
	userRole repository.UserRoleRepositories,
) Resolver {
	return &userAuth{
		user:     user,
		userRole: userRole,
	}
}

func (u *userAuth) Login(ctx context.Context, payload presentations.PayloadLogin) (presentations.Response, error) {
	findUser, err := u.user.FindUser(ctx, entity.Users{
		Email: payload.Email,
	})
	if err != nil {
		log.Printf("failed to find user %v", err)
		return presentations.Response{
			Code:    http.StatusInternalServerError,
			Message: "Internal Server Error",
		}, err
	}

	if findUser == nil {
		return presentations.Response{
			Code:    http.StatusNotFound,
			Message: "User not found",
		}, err
	}

	findRole, err := u.userRole.FindUserRole(ctx, entity.UserRoles{
		UserId: findUser.Id,
		RoleId: payload.RoleId,
	})

	if err != nil {
		log.Printf("failed to find user role %v", err)
		return presentations.Response{
			Code:    http.StatusInternalServerError,
			Message: "Internal Server Error",
		}, err
	}

	if findUser == nil {
		return presentations.Response{
			Code:    http.StatusForbidden,
			Message: "Forbidden Access",
		}, err
	}

	err = compareHashPassword([]byte(findUser.Password), []byte(payload.Password))
	if err != nil {
		return presentations.Response{
			Code:    http.StatusUnauthorized,
			Message: "Unauthorized",
		}, err
	}

	exp := time.Now().Add(time.Hour * 15)
	claims := &consts.JWTClaim{
		Id:       findUser.Id,
		FullName: findUser.FullName,
		Email:    findUser.Email,
		RoleId:   findRole.RoleId,
		Address:  findUser.Address,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "go_komodo_service",
			ExpiresAt: jwt.NewNumericDate(exp),
		},
	}

	tokenAlgorithm := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenAlgorithm.SignedString(consts.JWT_KEY)
	if err != nil {
		return presentations.Response{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		}, err
	}

	return presentations.Response{
		Code:    http.StatusCreated,
		Message: "Token created",
		Data:    map[string]string{"token": token},
	}, nil
}
