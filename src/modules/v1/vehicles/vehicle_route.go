package vehicle

import (
	"github.com/depri11/vehicle/src/middleware"
	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

func NewRoute(r *mux.Router, db *gorm.DB) {
	route := r.PathPrefix("/vehicle").Subrouter()

	repository := NewRepository(db)
	service := NewService(repository)
	controller := NewController(service)

	route.HandleFunc("/", controller.GetAll).Methods("GET")
	route.HandleFunc("/all", controller.Query).Methods("GET")
	route.HandleFunc("/popular", controller.PopularVehicle).Methods("GET")
	route.HandleFunc("/", middleware.Do(controller.Create, middleware.CheckAuth)).Methods("POST")
	route.HandleFunc("/{id}", controller.GetVehicle).Methods("GET")
	route.HandleFunc("/{id}", middleware.Do(controller.UpdateVehicle, middleware.CheckAuth)).Methods("PUT")
	route.HandleFunc("/{id}", middleware.Do(controller.DeleteVehicle, middleware.CheckAuth)).Methods("DELETE")

}
