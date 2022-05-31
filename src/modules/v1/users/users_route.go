package users

import (
	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

func NewRoute(r *mux.Router, db *gorm.DB) {
	route := r.PathPrefix("/users").Subrouter()

	repository := NewRepository(db)
	controller := NewController(repository)

	route.HandleFunc("/", controller.GetAll)
}
