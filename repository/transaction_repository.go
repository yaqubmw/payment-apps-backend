package repository

import (
	"fmt"
	"payment-apps-backend/model"

	"gorm.io/gorm"
)

type TransactionRepository interface {
	Create(payload model.Transaction) error
	Get(id string) (model.Transaction, error)
	List() ([]model.Transaction, error)
}

type transactionRepository struct {
	db *gorm.DB
}

func (t *transactionRepository) Create(payload model.Transaction) error {
	return t.db.Create(&payload).Error
}

func (t *transactionRepository) List() ([]model.Transaction, error) {
	var transaction []model.Transaction
	result := t.db.Find(&transaction)
	if result.Error != nil {
		return nil, result.Error
	}

	fmt.Println("transaction retrieve all successfully")
	return transaction, nil
}

func (t *transactionRepository) Get(id string) (model.Transaction, error) {
	transaction := model.Transaction{}
	err := t.db.Preload("Role").Where("id = ?", id).First(&transaction).Error
	return transaction, err
}

func NewTransactionRepository(db *gorm.DB) TransactionRepository {
	return &transactionRepository{
		db: db,
	}
}
