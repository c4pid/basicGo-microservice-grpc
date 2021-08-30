package models

import (
	"time"

	"gorm.io/gorm"
)

type Flight struct {
	Id            uint   `gorm:"primarykey"`
	Name          string `gorm:"type:varchar(200);not null"`
	From          string `gorm:"type:varchar(10)"`
	To            string `gorm:"type:varchar(10)"`
	Date          time.Time
	Status        string `gorm:"type:varchar(10)"`
	AvailableSlot int64  `gorm:"type:integer"`
	CreatedAt     time.Time
	UpdatedAt     time.Time
	DeletedAt     gorm.DeletedAt `gorm:"index"`
}

type ListFlights struct {
	Flights []*Flight
}
