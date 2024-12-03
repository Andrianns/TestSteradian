package service

import (
	"Stera/models"
	"Stera/repository"
)

type OrderService interface {
	GetAll() ([]models.Order, error)
}

type orderService struct {
	repo repository.OrderRepository
}

func NewOrderService(repo repository.OrderRepository) OrderService {
	return &orderService{
		repo: repo,
	}
}

func (s *orderService) GetAll() ([]models.Order, error) {
	return s.repo.GetAll()
}
