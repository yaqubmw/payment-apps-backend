package usecase

import (
	"fmt"
	"payment-apps-backend/security"
)

type AuthUseCase interface {
	Login(username string, password string) (string, error)
}

type authUseCase struct {
	usecase CustomerUseCase
}

func (a *authUseCase) Login(username, password string) (string, error) {
	user, err := a.usecase.FindUsernamePassword(username, password)
	if err != nil {
		return "", fmt.Errorf("invalid username and password or user is not active")
	}

	// mekanisme jika user itu ada akan membalikan sebuah token
	token, err := security.CreateAccessToken(user)
	if err != nil {
		return "", fmt.Errorf("failed to generate token")
	}
	return token, nil
}

func NewAuthUseCase(usecase CustomerUseCase) AuthUseCase {
	return &authUseCase{usecase: usecase}
}
