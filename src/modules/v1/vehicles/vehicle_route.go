package vehicle

import (
	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

func NewRoute(r *mux.Router, db *gorm.DB) {
	route := r.PathPrefix("/vehicle").Subrouter()

	repository := NewRepository(db)
	controller := NewController(repository)

	route.HandleFunc("/", controller.GetAll).Methods("GET")
	route.HandleFunc("/all", controller.QuerySort).Methods("GET")
	route.HandleFunc("/popular", controller.PopularVehicle).Methods("GET")
	route.HandleFunc("/", controller.Create).Methods("POST")
	route.HandleFunc("/{id}", controller.GetVehicle).Methods("GET")
	route.HandleFunc("/{id}", controller.UpdateVehicle).Methods("PUT")
	route.HandleFunc("/{id}", controller.DeleteVehicle).Methods("DELETE")

}
