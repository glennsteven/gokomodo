package handler

import (
	"errors"
	"gokomodo_test/internal/entity"
	"gokomodo_test/internal/middleware"
)

// DecodeJWT decodes JWT token and returns user information
func DecodeJWT(tokenString string) (*entity.Users, error) {
	claims, err := middleware.ParseJWT(tokenString)
	if err != nil {
		return nil, errors.New("unable to parse JWT token")
	}

	fullName, ok := claims["FullName"].(string)
	if !ok {
		return nil, errors.New("FullName field missing in claims")
	}

	userIdFloat64, ok := claims["Id"].(float64)
	if !ok {
		return nil, errors.New("id field missing in claims")
	}
	userId := int(userIdFloat64)

	address, ok := claims["Address"].(string)
	if !ok {
		return nil, errors.New("address field missing in claims")
	}

	// Construct User struct
	user := &entity.Users{
		Id:       userId,
		FullName: fullName,
		Address:  address,
	}

	return user, nil
}
