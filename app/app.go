package app

import (
	domain "capi/domain/customer"
	"capi/logger"
	"capi/service"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func sanityCheck() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func Start() {
	sanityCheck()
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
	serverAddr := os.Getenv("SERVER_ADDRESS")
	serverPort := os.Getenv("SERVER_PORT")
	logger.Info(fmt.Sprintf("Start server on %s:%s ...", serverAddr, serverPort))
	http.ListenAndServe(fmt.Sprintf("%s:%s", serverAddr, serverPort), mux)
}
