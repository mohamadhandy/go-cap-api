package service

import (
	"capi/domain"
	"capi/dto"
	"capi/errs"
	"capi/logger"
	"strings"

	"github.com/golang-jwt/jwt/v4"
)

type AuthService interface {
	Login(dto.LoginRequest) (*dto.LoginResponse, *errs.AppErr)
}

type DefaultAuthService struct {
	repository domain.AuthRepository
}

func NewAuthService(repository domain.AuthRepository) DefaultAuthService {
	return DefaultAuthService{repository}
}

func (s DefaultAuthService) Login(req dto.LoginRequest) (*dto.LoginResponse, *errs.AppErr) {
	login, errApp := s.repository.FindBy(req.Username, req.Password)
	if errApp != nil {
		return nil, errApp
	}
	accounts := strings.Split(login.Accounts, ",")
	claims := domain.AccessTokenClaims{
		CustomerID: login.CustomerID,
		Username:   login.Username,
		Role:       login.Role,
		Accounts:   accounts,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString([]byte("rahasia"))
	if err != nil {
		logger.Error("Failed while signing access token: " + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected error")
	}
	return &dto.LoginResponse{signedToken}, nil
}
