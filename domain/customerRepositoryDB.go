package domain

import (
	"database/sql"
	"errors"
	"log"

	_ "github.com/lib/pq"
)

type CustomerRepositoryDB struct {
	db *sql.DB
}

func NewCustomerRepositoryDB() CustomerRepositoryDB {
	connStr := "postgres://postgres:admin@localhost/banking?sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	return CustomerRepositoryDB{db}
}

func (d CustomerRepositoryDB) FindByID(customerId string) (*Customer, error) {
	query := "select * from customers where customer_id = $1"
	row := d.db.QueryRow(query, customerId)
	var c Customer
	err := row.Scan(&c.ID, &c.Name, &c.DateOfBirth, &c.City, &c.ZipCode, &c.Status)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Fatal("error customer data not found", err.Error())
			return nil, errors.New("customer data not found")
		} else {
			log.Fatal("error scanning customer data", err.Error())
			return nil, err
		}
	}
	return &c, nil
}

func (d CustomerRepositoryDB) FindAll() ([]Customer, error) {

	query := "select * from customers"
	rows, err := d.db.Query(query)
	if err != nil {
		log.Fatal("Error query customer table", err.Error())
		return nil, err
	}
	var customers []Customer
	for rows.Next() {
		var c Customer
		err := rows.Scan(&c.ID, &c.Name, &c.DateOfBirth, &c.City, &c.ZipCode, &c.Status)
		if err != nil {
			log.Fatal("error scanning customer data", err.Error())
			return nil, err
		}
		customers = append(customers, c)
	}
	return customers, nil
}
