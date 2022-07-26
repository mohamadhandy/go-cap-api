package app

import (
	"encoding/json"
	"encoding/xml"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Customer struct {
	ID      int    `json:"id" xml:"id"`
	Name    string `json:"name" xml:"name"`
	City    string `json:"city" xml:"city"`
	ZipCode string `json:"zipCode" xml:"zipcode"`
}

var customers []Customer = []Customer{
	{
		1, "User1", "Jakarta", "181818",
	},
	{
		2, "User2", "Bandung", "989898",
	},
}

func Greet(w http.ResponseWriter, r *http.Request) {
	log.Println("Greet")
}

func GetCustomers(w http.ResponseWriter, r *http.Request) {
	if r.Header.Get("Content-Type") == "application/xml" {
		w.Header().Add("Content-Type", "application/xml")
		xml.NewEncoder(w).Encode(customers)
	} else {
		w.Header().Add("Content-Type", "application/json")
		json.NewEncoder(w).Encode(customers)
	}
}

func GetCustomer(w http.ResponseWriter, r *http.Request) {
	customer := mux.Vars(r)
	customerId, _ := strconv.Atoi(customer["customer_id"])

	var getCustomer Customer
	for _, v := range customers {
		if v.ID == customerId {
			getCustomer = v
		}
	}
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(getCustomer)
}
