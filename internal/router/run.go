package router

import (
	"github.com/gorilla/mux"
	"gokomodo_test/internal/config"
	"log"
	"net/http"
)

func Run() {
	r := mux.NewRouter()
	cfg := config.NewViper()
	Router(r)
	port := cfg.GetString("web.port")
	appName := cfg.GetString("app.name")
	log.Printf("%s is running with port %s", appName, port)
	log.Println(http.ListenAndServe(":"+port, r))
}
