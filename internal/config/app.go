package config

import (
	"database/sql"
	"github.com/gorilla/mux"
	"github.com/spf13/viper"
)

type BootstrapConfig struct {
	DB     *sql.DB
	App    *mux.Router
	Config *viper.Viper
}
