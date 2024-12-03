package repository

import (
	"Stera/models"

	"gorm.io/gorm"
)

type OrderRepository interface {
	GetAll() ([]models.Order, error)
}

type orderRepository struct {
	db *gorm.DB
}

func NewOrderRepository(db *gorm.DB) OrderRepository {
	return &orderRepository{db: db}
}

func (r *orderRepository) GetAll() ([]models.Order, error) {
	var orders []models.Order
	if err := r.db.Find(&orders).Error; err != nil {
		return nil, err
	}
	return orders, nil
}
