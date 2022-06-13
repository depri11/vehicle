package history

import (
	"github.com/depri11/vehicle/src/middleware"
	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

func NewRoute(mux *mux.Router, db *gorm.DB) {
	r := mux.PathPrefix("/history").Subrouter()

	repository := NewRepository(db)
	service := NewService(repository)
	controller := NewController(service)

	r.HandleFunc("/", controller.GetAll).Methods("GET")
	r.HandleFunc("/all", controller.Query).Methods("GET")
	r.HandleFunc("/", middleware.Do(controller.Create, middleware.CheckAuth)).Methods("POST")
	r.HandleFunc("/{id}", middleware.Do(controller.Update, middleware.CheckAuth)).Methods("PUT")
	r.HandleFunc("/{id}", middleware.Do(controller.GetHistorys, middleware.CheckAuth)).Methods("GET")
	r.HandleFunc("/{id}", middleware.Do(controller.DeleteHistory, middleware.CheckAuth)).Methods("DELETE")
}
