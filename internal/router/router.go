package router

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"gokomodo_test/internal/config"
	"gokomodo_test/internal/repository"
	"gokomodo_test/internal/usecase/auth"
	"net/http"
)

func Router(r *mux.Router) {
	viperConfig := config.NewViper()
	log := config.NewLogger(viperConfig)
	db := config.NewDatabase(viperConfig, log)

	usersRepository := repository.NewUsers(db)

	userService := auth.NewUserAuth(usersRepository)

	r.HandleFunc("/api/login",
		userService.Login,
	).Methods(http.MethodPost)

	//sub.Use(middleware.JWTMiddleware)
}
