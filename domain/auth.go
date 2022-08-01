package domain

import (
	"capi/errs"

	"github.com/golang-jwt/jwt/v4"
)

type Login struct {
	Username   string `db:"username"`
	CustomerID string `db:"customer_id"`
	Accounts   string `db:"account_numbers"`
	Role       string `db:"role"`
}

type AccessTokenClaims struct {
	Username   string   `json:"username"`
	CustomerID string   `json:"customer_id"`
	Accounts   []string `json:"account_numbers"`
	Role       string   `json:"role"`
	jwt.StandardClaims
}

type AuthRepository interface {
	FindBy(username string, password string) (*Login, *errs.AppErr)
}
