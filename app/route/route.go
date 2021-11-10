package route

import (
	"github.com/gorilla/mux"
	"transport-manager/m/v1/app/controller"
)

// Routes -> define endpoints
func Routes() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/transports", controller.GetTransportEndpoint).Methods("GET")
	//router.HandleFunc("/", controller.GetTransportEndpoint).Methods("POST")
	return router
}
