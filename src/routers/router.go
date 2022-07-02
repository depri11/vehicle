package routers

import (
	"errors"
	"net/http"

	"github.com/depri11/vehicle/src/database"
	"github.com/depri11/vehicle/src/modules/v1/auth"
	"github.com/depri11/vehicle/src/modules/v1/history"
	"github.com/depri11/vehicle/src/modules/v1/users"
	vehicle "github.com/depri11/vehicle/src/modules/v1/vehicles"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func SetupRouter() (http.Handler, error) {
	route := mux.NewRouter()
	db, err := database.SetupDB()
	if err != nil {
		return nil, errors.New("failed connecting to database")
	}

	c := cors.New(cors.Options{
		AllowedHeaders:   []string{"X-Requested-With", "Content-Type", "Authorization"},
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE"},
		AllowCredentials: true,
		Debug:            true,
	}).Handler(route)

	auth.NewRoute(route, db)
	users.NewRoute(route, db)
	vehicle.NewRoute(route, db)
	history.NewRoute(route, db)

	return c, nil
}
