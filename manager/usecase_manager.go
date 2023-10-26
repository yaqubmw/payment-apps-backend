package manager

import "payment-apps-backend/usecase"

type UseCaseManager interface {
	CustomerUseCase() usecase.CustomerUseCase
	AuthUseCase() usecase.AuthUseCase
	MerchantUseCase() usecase.MerchantUseCase
	TransactionUseCase() usecase.TransactionUseCase
}

type useCaseManager struct {
	repoManager RepoManager
}

// TransactionUseCase implements UseCaseManager.
func (u *useCaseManager) TransactionUseCase() usecase.TransactionUseCase {
	return usecase.NewTransactionUseCase(u.repoManager.TransactionRepo())
}

// MerchantUseCase implements UseCaseManager.
func (u *useCaseManager) MerchantUseCase() usecase.MerchantUseCase {
	return usecase.NewMerchantUseCase(u.repoManager.MerchantRepo())
}

// AuthUseCase implements UseCaseManager.
func (u *useCaseManager) AuthUseCase() usecase.AuthUseCase {
	return usecase.NewAuthUseCase(u.CustomerUseCase())
}

// CustomerUseCase implements UseCaseManager.
func (u *useCaseManager) CustomerUseCase() usecase.CustomerUseCase {
	return usecase.NewCustomerUseCase(u.repoManager.CustomerRepo())
}

func NewUseCaseManager(repoManager RepoManager) UseCaseManager {
	return &useCaseManager{repoManager: repoManager}
}
