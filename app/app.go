package app

import (
	"capi/domain"
	"capi/service"
	"net/http"

	"github.com/gorilla/mux"
)

func Start() {
	ch := CustomerHandler{service.NewCustomerService(domain.NewCustomerRepositoryDB())}

	mux := mux.NewRouter()

	// * defining routes
	// mux.HandleFunc("/greet", Greet).Methods(http.MethodGet)
	// mux.HandleFunc("/customers", AddCustomer).Methods(http.MethodPost)

	mux.HandleFunc("/customers", ch.GetAllCustomer).Methods(http.MethodGet)
	mux.HandleFunc("/customers/{customer_id:[0-9]+}", ch.GetCustomerByID).Methods(http.MethodGet)
	// mux.HandleFunc("/customers/{customer_id:[0-9]+}", DeleteCustomer).Methods(http.MethodDelete)
	// mux.HandleFunc("/customers/{customer_id:[0-9]+}", UpdateCustomer).Methods(http.MethodPut)

	// * starting the server
	http.ListenAndServe(":8080", mux)
}
