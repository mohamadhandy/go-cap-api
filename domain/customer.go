package domain

import (
	"capi/dto"
	"capi/errs"
)

type Customer struct {
	ID          string `json:"id" db:"customer_id"`
	Name        string `json:"name"`
	City        string `json:"city"`
	ZipCode     string `json:"zip_code"`
	DateOfBirth string `json:"date_of_birth" db:"date_of_birth"`
	Status      string `json:"status"`
}

type CustomerRepository interface {
	FindAll(string) ([]Customer, *errs.AppErr)
	FindByID(string) (*Customer, *errs.AppErr)
}

func (c Customer) convertStatusName() string {
	statusName := "active"
	if c.Status == "0" {
		statusName = "inactive"
	}
	return statusName
}

func (c Customer) ToDTO() dto.CustomerResponse {

	return dto.CustomerResponse{
		ID:          c.ID,
		Name:        c.Name,
		City:        c.City,
		DateOfBirth: c.DateOfBirth,
		ZipCode:     c.ZipCode,
		Status:      c.convertStatusName(),
	}
}
