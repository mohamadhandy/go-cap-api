package app

import (
	"capi/service"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

// type Customer struct {
// 	ID      int    `json:"id" xml:"id"`
// 	Name    string `json:"name" xml:"name"`
// 	City    string `json:"city" xml:"city"`
// 	ZipCode string `json:"zipCode" xml:"zipcode"`
// }

// var customers []Customer = []Customer{
// 	{
// 		1, "User1", "Jakarta", "181818",
// 	},
// 	{
// 		2, "User2", "Bandung", "989898",
// 	},
// }

// func Greet(w http.ResponseWriter, r *http.Request) {
// 	log.Println("Greet")
// }

type CustomerHandler struct {
	service service.CustomerService
}

func (ch *CustomerHandler) GetAllCustomer(w http.ResponseWriter, r *http.Request) {
	customers, err := ch.service.GetAllCustomer()
	if err != nil {
		panic(err.Error())
	}
	if r.Header.Get("Content-Type") == "application/xml" {
		w.Header().Add("Content-Type", "application/xml")
		xml.NewEncoder(w).Encode(customers)
	} else {
		w.Header().Add("Content-Type", "application/json")
		json.NewEncoder(w).Encode(customers)
	}
}

func (ch *CustomerHandler) GetCustomerByID(w http.ResponseWriter, r *http.Request) {
	customerVars := mux.Vars(r)
	customerId := customerVars["customer_id"]
	customer, err := ch.service.GetCustomerByID(customerId)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprint(w, err.Error())
		return
	}
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(customer)
}

// func AddCustomer(w http.ResponseWriter, r *http.Request) {
// 	var cust Customer
// 	json.NewDecoder(r.Body).Decode(&cust)

// 	nextId := getNextID()
// 	cust.ID = nextId

// 	customers = append(customers, cust)
// 	w.WriteHeader(http.StatusCreated)
// 	log.Println("Add customer successfully")
// }

// func DeleteCustomer(w http.ResponseWriter, r *http.Request) {
// 	customer := mux.Vars(r)
// 	customerId, _ := strconv.Atoi(customer["customer_id"])

// 	var counter int = 0
// 	for index, v := range customers {
// 		if v.ID == customerId {
// 			customers = append(customers[:index], customers[index+1:]...)
// 			counter++
// 		}
// 	}
// 	if counter > 0 {
// 		w.Header().Add("Content-Type", "application/json")
// 		fmt.Fprint(w, "Remove customer successfully")
// 	} else {
// 		fmt.Fprint(w, "Delete customer unsuccessful, customer not found")
// 	}
// }

// func UpdateCustomer(w http.ResponseWriter, r *http.Request) {
// 	fmt.Println("UpdateCustomer")
// 	w.Header().Set("Content-Type", "application/json")
// 	params := mux.Vars(r)
// 	cust_id, _ := strconv.Atoi(params["customer_id"])
// 	for index, item := range customers {
// 		if item.ID == cust_id {
// 			customers = append(customers[:index], customers[index+1:]...)
// 			var customerData Customer
// 			_ = json.NewDecoder(r.Body).Decode(&customerData)
// 			customerData.ID = cust_id
// 			customers = append(customers, customerData)
// 			json.NewEncoder(w).Encode(customerData)
// 			return
// 		}
// 	}
// }

// func getNextID() int {
// 	cust := customers[len(customers)-1]
// 	return cust.ID + 1
// }
