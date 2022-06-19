package users

import (
	"github.com/depri11/vehicle/src/middleware"
	"github.com/gorilla/mux"
	newrelic "github.com/newrelic/go-agent/v3/newrelic"
	"gorm.io/gorm"
)

func NewRoute(r *mux.Router, db *gorm.DB) {
	route := r.PathPrefix("/users").Subrouter()

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
	route.HandleFunc("/", controller.Register).Methods("POST")
	route.HandleFunc("/{id}", middleware.Do(controller.GetUserID, middleware.CheckAuth)).Methods("GET")
	route.HandleFunc("/{id}", middleware.Do(controller.UpdateUser, middleware.CheckAuth)).Methods("PUT")
	route.HandleFunc("/{id}", middleware.Do(controller.DeleteUser, middleware.CheckAuth)).Methods("DELETE")
}
