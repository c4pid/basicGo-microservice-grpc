package models

import (
	"time"

	"gorm.io/gorm"
)

type Booking struct {
	BookingCode int64 `gorm:"primarykey"`
	BookingDate time.Time
	CustomerId  int64  `gorm:"type:integer;not null"`
	FlightId    int64  `gorm:"type:integer;not null"`
	Status      string `gorm:"type:varchar(10);not null"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt `gorm:"index"`
}
