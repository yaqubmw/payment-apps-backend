package usecase

import (
	"fmt"
	"payment-apps-backend/model"
	"payment-apps-backend/repository"

	"golang.org/x/crypto/bcrypt"
)

type CustomerUseCase interface {
	RegisterCustomer(payload model.Customer) error
	FindByIdCust(id string) (model.Customer, error)
	FindUsername(username string) (model.Customer, error)
	FindUsernamePassword(username, password string) (model.Customer, error)
	UpdateCustomer(payload model.Customer) error
}

type customerUseCase struct {
	repo repository.CustomerRepository
}

// RegisterCustomer implements CustomerUseCase.
func (c *customerUseCase) RegisterCustomer(payload model.Customer) error {
	bytes, _ := bcrypt.GenerateFromPassword([]byte(payload.Password), bcrypt.DefaultCost)
	payload.Password = string(bytes)
	err := c.repo.Create(payload)
	if err != nil {
		return fmt.Errorf("failed to create user: %v", err)
	}
	return nil
}

// FindByIdCust implements CustomerUseCase.
func (c *customerUseCase) FindByIdCust(id string) (model.Customer, error) {
	return c.repo.Get(id)
}

// FindUsername implements CustomerUseCase.
func (c *customerUseCase) FindUsername(username string) (model.Customer, error) {
	return c.repo.GetUsername(username)
}

// FindUsernamePassword implements CustomerUseCase.
func (c *customerUseCase) FindUsernamePassword(username string, password string) (model.Customer, error) {
	return c.repo.GetUsernamePassword(username, password)
}

// UpdateCustomer implements CustomerUseCase.
func (c *customerUseCase) UpdateCustomer(payload model.Customer) error {
	bytes, _ := bcrypt.GenerateFromPassword([]byte(payload.Password), bcrypt.DefaultCost)
	payload.Password = string(bytes)
	err := c.repo.Update(payload)
	if err != nil {
		return fmt.Errorf("failed to create user: %v", err)
	}
	return nil
}

func NewCustomerUseCase(repo repository.CustomerRepository) CustomerUseCase {
	return &customerUseCase{repo: repo}
}
