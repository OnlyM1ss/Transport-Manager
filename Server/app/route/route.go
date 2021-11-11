package route

import (
	"github.com/gorilla/mux"
	"transport-manager/m/v1/app/controller"
)

// Routes -> define endpoints
func Routes() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/transport", controller.GetTransportsEndpoint).Methods("GET")
	router.HandleFunc("/transport", controller.CreateTransportEndpoint).Methods("POST")
	router.HandleFunc("/transport/{id}", controller.DeleteTransportEndpoint).Methods("DELETE")
	router.HandleFunc("/transport/{id}", controller.GetTransportEndpoint).Methods("GET")
	router.HandleFunc("/login", controller.Login).Methods("GET")
	router.HandleFunc("/account", controller.Account).Methods("GET")
	return router
}
