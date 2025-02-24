package services

import (
	"go-sample-post/models"
	"go-sample-post/repositories"
)

type CustomerService interface {
	CreateCustomer(customer *models.Customer) error
	GetAllCustomers() ([]models.Customer, error)
	GetCustomerByID(id string) (*models.Customer, error)
	UpdateCustomer(id string, customer *models.Customer) error
	DeleteCustomer(id string) error
}

type customerServiceImpl struct {
	Repo repositories.CustomerRepository
}

func NewCustomerService(repo repositories.CustomerRepository) CustomerService {
	return &customerServiceImpl{repo}
}

func (s *customerServiceImpl) CreateCustomer(customer *models.Customer) error {
	return s.Repo.Create(customer)
}

func (s *customerServiceImpl) GetAllCustomers() ([]models.Customer, error) {
	return s.Repo.GetAll()
}

func (s *customerServiceImpl) GetCustomerByID(id string) (*models.Customer, error) {
	return s.Repo.GetByID(id)
}

func (s *customerServiceImpl) UpdateCustomer(id string, customer *models.Customer) error {
	return s.Repo.Update(id, customer)
}

func (s *customerServiceImpl) DeleteCustomer(id string) error {
	return s.Repo.Delete(id)
}
