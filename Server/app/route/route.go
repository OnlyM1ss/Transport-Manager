package route

import (
	"github.com/OnlyM1ss/transport-manager/app/controller"
	"github.com/gorilla/mux"
)

// Routes -> define endpoints
func Routes() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/transportTypes", controller.GetTransportTypesEndpoint).Methods("GET")
	router.HandleFunc("/transport/{name}", controller.GetTransportByTypeName).Methods("GET")
	router.HandleFunc("/transport", controller.GetTransportsEndpoint).Methods("GET")
	router.HandleFunc("/transport", controller.UpdateTransportsEndpoint).Methods("PUT")
	router.HandleFunc("/transport", controller.CreateTransportEndpoint).Methods("POST")
	router.HandleFunc("/transport/{id}", controller.DeleteTransportEndpoint).Methods("DELETE")
	router.HandleFunc("/transport/{id}", controller.GetTransportEndpoint).Methods("GET")
	//Authorization
	router.HandleFunc("/login", controller.Login)
	router.HandleFunc("/account", controller.Account)
	return router
}
