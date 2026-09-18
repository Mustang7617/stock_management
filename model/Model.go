package model

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	HouseID  uint    `gorm:"primaryKey"`
	UserName string  `gorm:"column:username"`
	Name     string  `gorm:"column:name"`
	Password string  `gorm:"column:password"`
	Role     string  `gorm:"column:role"`
	Visits   []Visit `gorm:"foreignKey:HouseID;references:HouseID"` // One user can have many visits
}

type Visit struct {
	gorm.Model
	VisitID      uint      `gorm:"primaryKey;autoIncrement" column:"visit_id"`
	HouseID      uint      `gorm:"column:house_id "` // Foreign key to the User's HouseID
	Reason       string    `gorm:"column:reason"`
	Status       string    `gorm:"column:status"`
	ArrivalTime  time.Time `gorm:"column:arrival_time"`
	LicensePlate string    `gorm:"column:license"`
}

// type Guest struct {
// 	gorm.Model
// 	GuestID   uint   `gorm:"primaryKey"`
// 	GuestName string `gorm:"column:guest_name"`
// 	GuestCar  string `gorm:"column:guest_car"`
// 	VisitID   uint   `gorm:"column:visit_id"` // Foreign key to Visit
// }
