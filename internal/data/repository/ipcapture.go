package repository

import (
	"enzappmob/internal/data/models"

	"gorm.io/gorm"
)

type IpCaptureRepository struct {
	db *gorm.DB
}

func NewIpCaptureRepository(db *gorm.DB) *IpCaptureRepository {
	return &IpCaptureRepository{db}
}

func (r *IpCaptureRepository) Create(ipCaptureObject *models.IpCapture) error {
	return r.db.Create(ipCaptureObject).Error
}

func (r *IpCaptureRepository) FindByID(id uint) (*models.IpCapture, error) {
	var ipCaptureObject models.IpCapture

	err := r.db.First(&ipCaptureObject, id).Error

	return &ipCaptureObject, err
}

func (r *IpCaptureRepository) Update(ipCaptureObject *models.IpCapture) error {
	return r.db.Save(ipCaptureObject).Error
}

func (r *IpCaptureRepository) Delete(ipCaptureObject *models.IpCapture) error {
	return r.db.Delete(ipCaptureObject).Error
}

func (r *IpCaptureRepository) FindAll() ([]models.IpCapture, error) {
	var ipCaptureObject []models.IpCapture
	err := r.db.Find(&ipCaptureObject).Error
	return ipCaptureObject, err
}
