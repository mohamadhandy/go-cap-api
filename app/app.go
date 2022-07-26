package app

import (
	"net/http"

	"github.com/gorilla/mux"
)

func Start() {

	mux := mux.NewRouter()

	// * defining routes
	mux.HandleFunc("/greet", Greet).Methods(http.MethodGet)

	mux.HandleFunc("/customers", GetCustomers).Methods(http.MethodGet)
	mux.HandleFunc("/customers/{customer_id}", GetCustomer).Methods(http.MethodGet)

	// * starting the server
	http.ListenAndServe(":8080", mux)
}
