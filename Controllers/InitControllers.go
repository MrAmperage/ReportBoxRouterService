package Controllers

import (
	"github.com/gorilla/mux"
)

func InitControllers(Router *mux.Router) {

	Router.HandleFunc("/", RootController).Methods("GET")
}
