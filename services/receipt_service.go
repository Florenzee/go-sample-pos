package services

import (
	"go-sample-post/models"
	"go-sample-post/repositories"
)

type ReceiptService interface {
	CreateReceipt(receipt *models.Receipt) error
	GetAllReceipts() ([]models.Receipt, error)
	GetReceiptByID(id string) (*models.Receipt, error)
	UpdateReceipt(id string, receipt *models.Receipt) error
	DeleteReceipt(id string) error
}

type receiptServiceImpl struct {
	receiptRepo repositories.ReceiptRepository
}

func NewReceiptService(repo repositories.ReceiptRepository) ReceiptService {
	return &receiptServiceImpl{repo}
}

func (s *receiptServiceImpl) CreateReceipt(receipt *models.Receipt) error {
	return s.receiptRepo.Create(receipt)
}

func (s *receiptServiceImpl) GetAllReceipts() ([]models.Receipt, error) {
	return s.receiptRepo.GetAll()
}

func (s *receiptServiceImpl) GetReceiptByID(id string) (*models.Receipt, error) {
	return s.receiptRepo.GetByID(id)
}

func (s *receiptServiceImpl) UpdateReceipt(id string, receipt *models.Receipt) error {
	return s.receiptRepo.Update(id, receipt)
}

func (s *receiptServiceImpl) DeleteReceipt(id string) error {
	return s.receiptRepo.Delete(id)
}
