package router

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"gokomodo_test/internal/config"
	"gokomodo_test/internal/middleware"
	"gokomodo_test/internal/repository"
	"gokomodo_test/internal/usecase/auth"
	"gokomodo_test/internal/usecase/buyer"
	"gokomodo_test/internal/usecase/seller"
	"net/http"
)

func Router(r *mux.Router) {
	viperConfig := config.NewViper()
	log := config.NewLogger(viperConfig)
	db := config.NewDatabase(viperConfig, log)

	usersRepository := repository.NewUsers(db)
	userRolesRepository := repository.NewUserRoles(db)
	productsRepository := repository.NewProducts(db)
	ordersRepository := repository.NewOrders(db)
	OrderDetailsRepository := repository.NewOrderDetails(db)

	userService := auth.NewUserAuth(usersRepository, userRolesRepository)

	r.HandleFunc("/api/login",
		userService.Login,
	).Methods(http.MethodPost)

	// Sub-Router
	sub := r.PathPrefix("/api").Subrouter()

	productService := seller.NewProductUseCase(productsRepository)

	sub.HandleFunc("/seller/products",
		productService.Create,
	).Methods(http.MethodPost)

	orderService := buyer.NewOrderBuyerUseCase(
		ordersRepository,
		OrderDetailsRepository,
		productsRepository,
		usersRepository,
	)

	sub.HandleFunc("/buyer/orders",
		orderService.Create,
	).Methods(http.MethodPost)

	sub.Use(middleware.JWTMiddleware)
}
