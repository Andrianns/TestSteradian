package repository

import (
	"Stera/models"

	"gorm.io/gorm"
)

type CarRepository interface {
	GetAll() ([]models.Car, error)
	GetById(id int) (models.Car, error)
	Update(car models.Car) (models.Car, error)
	Delete(id int) error
	Create(car models.Car) (models.Car, error)
}

type carRepository struct {
	db *gorm.DB
}

func NewCarRepository(db *gorm.DB) CarRepository {
	return &carRepository{db: db}
}

func (r *carRepository) GetAll() ([]models.Car, error) {
	var cars []models.Car
	if err := r.db.Find(&cars).Error; err != nil {
		return nil, err
	}
	return cars, nil
}
func (r *carRepository) GetById(id int) (models.Car, error) {
	var car models.Car
	err := r.db.First(&car, id).Error

	return car, err
}
func (r *carRepository) Update(car models.Car) (models.Car, error) {
	err := r.db.Save(&car).Error
	return car, err
}
func (r *carRepository) Delete(id int) error {
	err := r.db.Delete(&models.Car{}, id).Error
	return err
}

func (r *carRepository) Create(car models.Car) (models.Car, error) {
	err := r.db.Create(&car).Error
	return car, err
}
