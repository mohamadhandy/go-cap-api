package app

import (
	"net/http"

	"github.com/gorilla/mux"
)

func Start() {

	mux := mux.NewRouter()

	// * defining routes
	mux.HandleFunc("/greet", Greet).Methods(http.MethodGet)
	mux.HandleFunc("/customers", AddCustomer).Methods(http.MethodPost)

	mux.HandleFunc("/customers", GetCustomers).Methods(http.MethodGet)
	mux.HandleFunc("/customers/{customer_id:[0-9]+}", GetCustomer).Methods(http.MethodGet)
	mux.HandleFunc("/customers/{customer_id:[0-9]+}", DeleteCustomer).Methods(http.MethodDelete)
	mux.HandleFunc("/customers/{customer_id:[0-9]+}", UpdateCustomer).Methods(http.MethodPut)

	// * starting the server
	http.ListenAndServe(":8080", mux)
}
