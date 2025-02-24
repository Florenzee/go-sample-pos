package services

import (
	"go-sample-post/models"
	"go-sample-post/repositories"
)

type EmployeeService interface {
	CreateEmployee(Employee *models.Employee) error
	GetAllEmployee() ([]models.Employee, error)
	GetEmployeeByID(id string) (*models.Employee, error)
	UpdateEmployee(id string, Employee *models.Employee) error
	DeleteEmployee(id string) error
}

type employeeServiceImpl struct {
	employeeRepo repositories.EmployeeRepository
}

func NewEmployeeService(repo repositories.EmployeeRepository) EmployeeService {
	return &employeeServiceImpl{repo}
}

func (s *employeeServiceImpl) CreateEmployee(Employee *models.Employee) error {
	return s.employeeRepo.Create(Employee)
}

func (s *employeeServiceImpl) GetAllEmployee() ([]models.Employee, error) {
	return s.employeeRepo.GetAll()
}

func (s *employeeServiceImpl) GetEmployeeByID(id string) (*models.Employee, error) {
	return s.employeeRepo.GetByID(id)
}

func (s *employeeServiceImpl) UpdateEmployee(id string, Employee *models.Employee) error {
	return s.employeeRepo.Update(id, Employee)
}

func (s *employeeServiceImpl) DeleteEmployee(id string) error {
	return s.employeeRepo.Delete(id)
}
