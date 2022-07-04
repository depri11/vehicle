package vehicle

import (
	"github.com/depri11/vehicle/src/middleware"
	"github.com/gorilla/mux"
	"github.com/newrelic/go-agent/v3/newrelic"
	"gorm.io/gorm"
)

func NewRoute(r *mux.Router, db *gorm.DB) {
	route := r.PathPrefix("/vehicle").Subrouter()

	repository := NewRepository(db)
	service := NewService(repository)
	controller := NewController(service)

	app, err := newrelic.NewApplication(
		newrelic.ConfigAppName("rentalvehicle"),
		newrelic.ConfigLicense("18451d107e236741290fd14cb735b2b19c20NRAL"),
		newrelic.ConfigDistributedTracerEnabled(true),
	)

	if err != nil {
		panic(err)
	}

	route.HandleFunc(newrelic.WrapHandleFunc(app, "/", controller.GetAll)).Methods("GET")
	route.HandleFunc(newrelic.WrapHandleFunc(app, "/all", controller.Query)).Methods("GET")
	route.HandleFunc("/popular", controller.PopularVehicle).Methods("GET")
	route.HandleFunc("/", middleware.Do(controller.Create, middleware.CheckAuth)).Methods("POST")
	route.HandleFunc("/{id}", controller.GetVehicle).Methods("GET")
	route.HandleFunc("/type/{type}", controller.GetVehiclesByType).Methods("GET")
	route.HandleFunc("/{id}", middleware.Do(controller.UpdateVehicle, middleware.CheckAuth)).Methods("PUT")
	route.HandleFunc("/{id}", middleware.Do(controller.DeleteVehicle, middleware.CheckAuth)).Methods("DELETE")

}
