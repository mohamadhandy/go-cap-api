package service

import (
	"capi/domain"
	"capi/dto"
	"capi/errs"
)

type CustomerService interface {
	GetAllCustomer(string) ([]dto.CustomerResponse, *errs.AppErr)
	GetCustomerByID(string) (*dto.CustomerResponse, *errs.AppErr)
}

type DefaultCustomerService struct {
	repository domain.CustomerRepository
}

func (s DefaultCustomerService) GetAllCustomer(customerStatus string) ([]dto.CustomerResponse, *errs.AppErr) {
	customers, err := s.repository.FindAll(customerStatus)
	if err != nil {
		return nil, errs.NewUnExpectedError("unexpected db error")
	}
	var dtoCustomers []dto.CustomerResponse
	for _, customer := range customers {
		dtoCustomers = append(dtoCustomers, customer.ToDTO())
	}
	return dtoCustomers, nil
}

func (s DefaultCustomerService) GetCustomerByID(customerId string) (*dto.CustomerResponse, *errs.AppErr) {
	customer, err := s.repository.FindByID(customerId)
	if err != nil {
		return nil, err
	}
	response := customer.ToDTO()
	return &response, nil
}

func NewCustomerService(repository domain.CustomerRepository) DefaultCustomerService {
	return DefaultCustomerService{repository: repository}
}
