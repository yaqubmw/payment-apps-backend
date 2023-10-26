package repository

import (
	"payment-apps-backend/model"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type CustomerRepository interface {
	Create(payload model.Customer) error
	Get(id string) (model.Customer, error)
	GetUsername(username string) (model.Customer, error)
	GetUsernamePassword(username string, password string) (model.Customer, error)
	Update(payload model.Customer) error
}

type customerRepository struct {
	db *gorm.DB
}

// Create implements CustomerRepository.
func (c *customerRepository) Create(payload model.Customer) error {
	return c.db.Create(&payload).Error
}

// Get implements CustomerRepository.
func (c *customerRepository) Get(id string) (model.Customer, error) {
	user := model.Customer{}
	err := c.db.Preload("Role").Where("id = ?", id).First(&user).Error
	return user, err
}

// GetUsername implements CustomerRepository.
func (c *customerRepository) GetUsername(username string) (model.Customer, error) {
	var customer model.Customer
	err := c.db.Where("username = ?", username).First(&customer).Error
	return customer, err
}

// GetUsernamePassword implements CustomerRepository.
func (c *customerRepository) GetUsernamePassword(username string, password string) (model.Customer, error) {
	customer, err := c.GetUsername(username)
	if err != nil {
		return model.Customer{}, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(customer.Password), []byte(password))
	if err != nil {
		return model.Customer{}, err
	}
	return customer, nil
}

// UpdateUser implements CustomerRepository.
func (c *customerRepository) Update(payload model.Customer) error {
	err := c.db.Model(&payload).Updates(payload).Error
	return err
}

func NewCustomerRepository(db *gorm.DB) CustomerRepository {
	return &customerRepository{
		db: db,
	}
}
