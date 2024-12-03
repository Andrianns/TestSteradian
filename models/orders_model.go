package models

import "time"

type Order struct {
	OrderId         int       `json:"order_id gorm:"primarykey"`
	CarId           int       `json:"car_id gorm:foreignKey:cars_id"`
	PickupDate      time.Time `json:"pickup_date"`
	DropoffDate     time.Time `json:"dropoff_date"`
	PickupLocation  time.Time `json:"pickup_location"`
	DropoffLocation string    `json:"dropoff_location"`
}

func (Order) TableName() string {
	return "orders"
}
