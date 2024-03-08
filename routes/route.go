package routes

import (
	"github.com/gorilla/mux"
	Handlers "github.com/maadiab/aldifaArchive/handlers"
)

func Router() *mux.Router {

	router := mux.NewRouter()

	router.HandleFunc("/", Handlers.ServeHomePage).Methods("GET")

	return router
}
