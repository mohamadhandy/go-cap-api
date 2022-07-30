package service

import (
	domain "capi/domain/account"
	"capi/errs"
)

type AccountService interface {
	CreateAccount(float64, string, string) (*domain.Account, *errs.AppErr)
}

type DefaultAccountService struct {
	domain.AccountRepositoryDB
}

func (s DefaultAccountService) CreateAccount(amount float64, customerId string, accountType string) (*domain.Account, *errs.AppErr) {
	account, err := s.AccountRepositoryDB.InsertAccount(amount, customerId, accountType)
	if err != nil {
		return nil, errs.NewUnExpectedError("unexpected error")
	}
	return account, nil
}

func NewAccountService(repository domain.AccountRepositoryDB) DefaultAccountService {
	return DefaultAccountService{repository}
}
