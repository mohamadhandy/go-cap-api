package domain

import (
	"capi/dto"
	"capi/errs"
)

type Account struct {
	ID          string  `json:"account_id" db:"account_id"`
	CustomerId  string  `json:"customer_id" db:"customer_id"`
	OpeningDate string  `json:"opening_date"`
	AccountType string  `json:"account_type"`
	Amount      float64 `json:"amount" db:"amount"`
	Status      string  `json:"status"`
}

type AccountRepository interface {
	InsertAccount(float64, string, string) (*Account, int, errs.AppErr)
}

func (a Account) ToDTO() dto.AccountResponse {

	return dto.AccountResponse{
		ID:          a.ID,
		OpeningDate: a.OpeningDate,
	}
}
