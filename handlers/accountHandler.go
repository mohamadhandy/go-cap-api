package handlers

import (
	"capi/errs"
	"capi/logger"
	"capi/service"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func NewAccountHandler(as service.DefaultAccountService) *AccountHandler {
	return &AccountHandler{as}
}

type AccountHandler struct {
	service service.DefaultAccountService
}

type handleAccount struct {
	CustomerId  string  `json:"customer_id"`
	Amount      float64 `json:"amount"`
	AccountType string  `json:"account_type"`
}

func (ah *AccountHandler) CreateAccount(w http.ResponseWriter, r *http.Request) {
	var acc handleAccount
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		logger.Error(err.Error())
	}
	err = json.Unmarshal(b, &acc)
	if err != nil {
		logger.Error(err.Error())
	}

	if acc.Amount > 5000 {
		ah.service.CreateAccount(acc.Amount, acc.CustomerId, acc.AccountType)
		writeResponse(w, http.StatusCreated, acc)
	} else {
		writeResponse(w, http.StatusBadRequest, errs.NewBadRequestError("Invalid amount, amount should more than 5000"))
		return
	}
}
