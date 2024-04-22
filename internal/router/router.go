package router

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"gokomodo_test/internal/config"
	orderproductservice "gokomodo_test/internal/handler/buyer/order_product"
	"gokomodo_test/internal/handler/login"
	acceptorderservice "gokomodo_test/internal/handler/seller/accept_order"
	productservice "gokomodo_test/internal/handler/seller/product"
	"gokomodo_test/internal/middleware"
	"gokomodo_test/internal/repository"
	"gokomodo_test/internal/usecase/auth"
	"gokomodo_test/internal/usecase/buyer/order_product"
	"gokomodo_test/internal/usecase/seller/accept_order"
	"gokomodo_test/internal/usecase/seller/product"
	"net/http"
)

func Router(r *mux.Router) {
	viperConfig := config.NewViper()
	log := config.NewLogger(viperConfig)
	db := config.NewDatabase(viperConfig, log)

	// Repository
	usersRepository := repository.NewUsers(db)
	userRolesRepository := repository.NewUserRoles(db)
	productsRepository := repository.NewProducts(db)
	ordersRepository := repository.NewOrders(db)
	OrderDetailsRepository := repository.NewOrderDetails(db)

	loginUseCase := auth.NewUserAuth(usersRepository, userRolesRepository)
	loginService := login.NewHandlerLogin(loginUseCase)

	r.HandleFunc("/api/login",
		loginService.Login,
	).Methods(http.MethodPost)

	// Sub-Router
	sub := r.PathPrefix("/api").Subrouter()

	productUseCase := product.NewProductUseCase(productsRepository)
	productService := productservice.NewHandlerSeller(productUseCase)
	sub.HandleFunc("/seller/products",
		productService.AddProduct,
	).Methods(http.MethodPost)

	acceptOrderUseCase := accept_order.NewAcceptOrderUseCase(OrderDetailsRepository, productsRepository)
	acceptOrderService := acceptorderservice.NewAcceptProductService(acceptOrderUseCase)
	sub.HandleFunc("/seller/accept-order",
		acceptOrderService.AcceptOrder,
	).Methods(http.MethodPost)

	orderProductUseCase := order_product.NewOrderBuyerUseCase(
		ordersRepository,
		OrderDetailsRepository,
		productsRepository,
		usersRepository,
	)
	orderProductService := orderproductservice.NewHandlerOrderProductService(orderProductUseCase)
	sub.HandleFunc("/buyer/orders",
		orderProductService.OrderProduct,
	).Methods(http.MethodPost)

	sub.Use(middleware.JWTMiddleware)
}
