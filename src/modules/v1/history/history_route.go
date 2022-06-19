package history

import (
	"github.com/depri11/vehicle/src/middleware"
	"github.com/gorilla/mux"
	"github.com/newrelic/go-agent/v3/newrelic"
	"gorm.io/gorm"
)

func NewRoute(mux *mux.Router, db *gorm.DB) {
	r := mux.PathPrefix("/history").Subrouter()

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

	r.HandleFunc(newrelic.WrapHandleFunc(app, "/", controller.GetAll)).Methods("GET")
	r.HandleFunc(newrelic.WrapHandleFunc(app, "/all", controller.Query)).Methods("GET")
	r.HandleFunc("/", middleware.Do(controller.Create, middleware.CheckAuth)).Methods("POST")
	r.HandleFunc("/{id}", middleware.Do(controller.Update, middleware.CheckAuth)).Methods("PUT")
	r.HandleFunc("/{id}", middleware.Do(controller.GetHistorys, middleware.CheckAuth)).Methods("GET")
	r.HandleFunc("/{id}", middleware.Do(controller.DeleteHistory, middleware.CheckAuth)).Methods("DELETE")
}
