package history

import (
	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

func NewRoute(mux *mux.Router, db *gorm.DB) {
	r := mux.PathPrefix("/history").Subrouter()

	repository := NewRepository(db)
	controller := NewController(repository)

	r.HandleFunc("/", controller.GetAll).Methods("GET")
	r.HandleFunc("/all", controller.QuerySort).Methods("GET")
	r.HandleFunc("/{id}", controller.GetHistorys).Methods("GET")
	r.HandleFunc("/{id}", controller.DeleteHistory).Methods("DELETE")
}
