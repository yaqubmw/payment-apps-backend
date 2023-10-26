package usecase

import (
	"fmt"
	"payment-apps-backend/model"
	"payment-apps-backend/repository"
)

type MerchantUseCase interface {
	RegisterMerchant(payload model.Merchant) error
	FindByIdMerch(id string) (model.Merchant, error)
}

type merchantUseCase struct {
	repo repository.MerchantRepository
}

// FindByIdCust implements MerchantUseCase.
func (m *merchantUseCase) FindByIdMerch(id string) (model.Merchant, error) {
	return m.repo.Get(id)
}

// RegisterMerchant implements MerchantUseCase.
func (m *merchantUseCase) RegisterMerchant(payload model.Merchant) error {
	err := m.repo.Create(payload)
	if err != nil {
		return fmt.Errorf("failed to create user: %v", err)
	}
	return nil
}

func NewMerchantUseCase(repo repository.MerchantRepository) MerchantUseCase {
	return &merchantUseCase{repo: repo}
}
