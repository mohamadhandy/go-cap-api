package domain

import (
	"capi/errs"
	"capi/logger"
	"database/sql"
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type CustomerRepositoryDB struct {
	db *sqlx.DB
}

func NewCustomerRepositoryDB() CustomerRepositoryDB {
	connStr := "postgres://postgres:admin@localhost/banking?sslmode=disable"
	db, err := sqlx.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	return CustomerRepositoryDB{db}
}

func (d CustomerRepositoryDB) FindByID(customerId string) (*Customer, *errs.AppErr) {
	query := "select * from customers where customer_id = $1"
	// row := d.db.QueryRow(query, customerId)
	var c Customer
	// err := row.Scan(&c.ID, &c.Name, &c.DateOfBirth, &c.City, &c.ZipCode, &c.Status)
	err := d.db.Get(&c, query, customerId)
	if err != nil {
		if err == sql.ErrNoRows {
			logger.Error("error customer data not found" + err.Error())
			return nil, errs.NewNotFoundError("Customer Not found")
		} else {
			logger.Error("error scanning customer data" + err.Error())
			return nil, errs.NewUnExpectedError("unexpected database error")
		}
	}
	return &c, nil
}

func (d CustomerRepositoryDB) FindAll(customerStatus string) ([]Customer, *errs.AppErr) {

	var c []Customer
	if customerStatus == "" {
		query := "select * from customers"
		err := d.db.Select(&c, query)
		if err != nil {
			logger.Error("Error query customer table" + err.Error())
			return nil, errs.NewUnExpectedError("unexpected database error")
		}
	} else {
		if customerStatus == "active" {
			customerStatus = "1"
		} else {
			customerStatus = "0"
		}
		query := "select * from customers where status = $1"
		err := d.db.Select(&c, query, customerStatus)
		if err != nil {
			logger.Error("Error query customer table" + err.Error())
			return nil, errs.NewUnExpectedError("unexpected database error")
		}
	}
	// var customers []Customer
	// for rows.Next() {
	// 	var c Customer
	// 	err := rows.Scan(&c.ID, &c.Name, &c.DateOfBirth, &c.City, &c.ZipCode, &c.Status)
	// 	if err != nil {
	// 		log.Fatal("error scanning customer data", err.Error())
	// 		return nil, err
	// 	}
	// 	customers = append(customers, c)
	// }
	return c, nil
}
