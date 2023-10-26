package usecase

import (
	"fmt"
	"payment-apps-backend/model"
	"payment-apps-backend/repository"
)

type TransactionUseCase interface {
	RegisterTransaction(payload model.Transaction) error
	FindByIdMerch(id string) (model.Transaction, error)
	FindAllTransaction() ([]model.Transaction, error)
}

type transactionUseCase struct {
	repo repository.TransactionRepository
}

// FindAllTransaction implements TransactionUseCase.
func (t *transactionUseCase) FindAllTransaction() ([]model.Transaction, error) {
	return t.repo.List()
}

// FindByIdCust implements TransactionUseCase.
func (t *transactionUseCase) FindByIdMerch(id string) (model.Transaction, error) {
	return t.repo.Get(id)
}

// RegisterTransaction implements TransactionUseCase.
func (t *transactionUseCase) RegisterTransaction(payload model.Transaction) error {
	err := t.repo.Create(payload)
	if err != nil {
		return fmt.Errorf("failed to create user: %v", err)
	}
	return nil
}

func NewTransactionUseCase(repo repository.TransactionRepository) TransactionUseCase {
	return &transactionUseCase{repo: repo}
}
