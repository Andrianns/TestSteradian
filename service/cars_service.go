package service

import (
	"Stera/models"
	"Stera/repository"
)

type CarService interface {
	GetAll() ([]models.Car, error)
	GetById(id int) (models.Car, error)
	Update(car models.Car) (models.Car, error)
	Delete(id int) error
	Create(models.Car) (models.Car, error)
}

type carService struct {
	repo repository.CarRepository
}

func NewCarService(repo repository.CarRepository) CarService {
	return &carService{
		repo: repo,
	}
}

func (s *carService) GetAll() ([]models.Car, error) {
	return s.repo.GetAll()
}

func (s *carService) GetById(id int) (models.Car, error) {
	return s.repo.GetById(id)
}
func (s *carService) Update(car models.Car) (models.Car, error) {
	return s.repo.Update(car)
}
func (s *carService) Delete(id int) error {
	return s.repo.Delete(id)
}

func (s *carService) Create(car models.Car) (models.Car, error) {
	return s.repo.Create(car)
}
