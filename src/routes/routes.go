package routes

import (
	"app/src/app/controller"
	"net/http"

	"github.com/gorilla/mux"
)

func Routes() *mux.Router {
	routes := mux.NewRouter()

	routes.HandleFunc("/users", controller.Index).Methods(http.MethodGet)
	routes.HandleFunc("/user/{id}", controller.Show).Methods(http.MethodGet)
	routes.HandleFunc("/user", controller.Store).Methods(http.MethodPost)
	routes.HandleFunc("/user/{id}", controller.Update).Methods(http.MethodPut)
	routes.HandleFunc("/user/{id}", controller.Delete).Methods(http.MethodDelete)
	routes.HandleFunc("/authentication", controller.Auth).Methods(http.MethodPost)

	return routes
}
