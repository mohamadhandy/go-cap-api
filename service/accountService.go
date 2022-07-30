package service

import (
	domain "capi/domain/account"
	"capi/dto"
	"capi/errs"
	"strconv"
)

type AccountService interface {
	CreateAccount(float64, string, string) (*dto.AccountResponse, *errs.AppErr)
}

type DefaultAccountService struct {
	domain.AccountRepositoryDB
}

func (s DefaultAccountService) CreateAccount(amount float64, customerId string, accountType string) (*dto.AccountResponse, *errs.AppErr) {
	account, accountId, err := s.AccountRepositoryDB.InsertAccount(amount, customerId, accountType)
	if err != nil {
		return nil, errs.NewUnExpectedError("unexpected error")
	}
	accountResponse := account.ToDTO()
	accountResponse.ID = strconv.Itoa(accountId)
	return &accountResponse, nil
}

func NewAccountService(repository domain.AccountRepositoryDB) DefaultAccountService {
	return DefaultAccountService{repository}
}
