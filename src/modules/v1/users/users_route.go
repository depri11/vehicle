package users

import (
	"github.com/depri11/vehicle/src/middleware.go"
	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

func NewRoute(r *mux.Router, db *gorm.DB) {
	route := r.PathPrefix("/users").Subrouter()

	repository := NewRepository(db)
	service := NewService(repository)
	controller := NewController(service)

	route.HandleFunc("/", controller.GetAll).Methods("GET")
	route.HandleFunc("/", controller.Register).Methods("POST")
	route.HandleFunc("/{id}", middleware.Do(controller.GetUserID, middleware.CheckAuth)).Methods("GET")
	// route.HandleFunc("/{id}", controller.UpdateUser).Methods("POST")
	// route.HandleFunc("/{id}", controller.DeleteUser).Methods("DELETE")
}
