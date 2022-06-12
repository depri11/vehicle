package history

import (
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
	r.HandleFunc("/", controller.Create).Methods("POST")
	r.HandleFunc("/{id}", controller.Update).Methods("PUT")
	r.HandleFunc("/{id}", controller.GetHistorys).Methods("GET")
	r.HandleFunc("/{id}", controller.DeleteHistory).Methods("DELETE")
}
