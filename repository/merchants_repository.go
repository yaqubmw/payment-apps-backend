package repository

import (
	"payment-apps-backend/model"

	"gorm.io/gorm"
)

type MerchantRepository interface {
	Create(payload model.Merchant) error
	Get(id string) (model.Merchant, error)
}

type merchantRepository struct {
	db *gorm.DB
}

// Create implements MerchantRepository.
func (m *merchantRepository) Create(payload model.Merchant) error {
	return m.db.Create(&payload).Error
}

// Get implements MerchantRepository.
func (m *merchantRepository) Get(id string) (model.Merchant, error) {
	merchant := model.Merchant{}
	err := m.db.Preload("Role").Where("id = ?", id).First(&merchant).Error
	return merchant, err
}

func NewMerchantRepository(db *gorm.DB) MerchantRepository {
	return &merchantRepository{
		db: db,
	}
}
