package repositories

import (
	"go-sample-post/models"
	"gorm.io/gorm"
)

type ReceiptRepository interface {
	Create(receipt *models.Receipt) error
	GetAll() ([]models.Receipt, error)
	GetByID(id string) (*models.Receipt, error)
	Update(id string, receipt *models.Receipt) error
	Delete(id string) error
}

type receiptRepositoryImpl struct {
	db *gorm.DB
}

func NewReceiptRepository(db *gorm.DB) ReceiptRepository {
	return &receiptRepositoryImpl{db}
}

func (r *receiptRepositoryImpl) Create(receipt *models.Receipt) error {
	return r.db.Create(receipt).Error
}

func (r *receiptRepositoryImpl) GetAll() ([]models.Receipt, error) {
	var receipt []models.Receipt
	err := r.db.Find(&receipt).Error
	return receipt, err
}

func (r *receiptRepositoryImpl) GetByID(id string) (*models.Receipt, error) {
	var receipt models.Receipt
	err := r.db.First(&receipt, "receipt_id = ?", id).Error
	return &receipt, err
}

func (r *receiptRepositoryImpl) Update(id string, receipt *models.Receipt) error {
	return r.db.Model(&models.Receipt{}).Where("receipt_id = ?", id).Updates(receipt).Error
}

func (r *receiptRepositoryImpl) Delete(id string) error {
	return r.db.Delete(&models.Receipt{}, "receipt_id = ?", id).Error
}
