package models

type Car struct {
	CarId     int     `json:"car_id" gorm:"primaryKey"`
	DayRate   float32 `json:"day_rate"`
	MonthRate float32 `json:"month_rate"`
	Image     string  `json:"image"`
	CarName   string  `json:"car_name"`
}

func (Car) TableName() string {
	return "cars"
}
