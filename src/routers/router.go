package routers

import (
	"errors"

	"github.com/depri11/vehicle/src/configs/database"
	"github.com/gorilla/mux"
)

func SetupRouter() (*mux.Router, error) {
	route := mux.NewRouter()
	db, err := database.SetupDB()
	if err != nil {
		return nil, errors.New("failed connecting to database")
	}
}
