package usecase

import (
	middleware "praktikum/middlewares"
	"praktikum/models/payload"
	"praktikum/repository/database"

	"github.com/labstack/echo"
)

type AuthUsecase interface {
	LoginUser(req *payload.LoginRequest) (res payload.LoginResponse, err error)
}

type authUsecase struct {
	authRepository database.AuthRepository
	userRepository database.UserRepository
}

func NewAuthUsecase(
	authRepository database.AuthRepository,
	userRepository database.UserRepository,
) *authUsecase {
	return &authUsecase{authRepository, userRepository}
}

func (a *authUsecase) LoginUser(req *payload.LoginRequest) (res payload.LoginResponse, err error) {

	user, err := a.authRepository.Login(req.Username, req.Password)
	if err != nil {
		return res, echo.NewHTTPError(400, "Account not Found !")
	}

	token, err := middleware.CreateToken(int(user.ID))
	if err != nil {
		return res, echo.NewHTTPError(400, "Failed to generate token")
	}

	user.Token = token

	res = payload.LoginResponse{
		Username: user.Username,
		Token:    user.Token,
	}

	return
}
