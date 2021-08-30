package models

import (
	"time"

	"gorm.io/gorm"
)

type Customer struct {
	Id          uint   `gorm:"primarykey"`
	Name        string `gorm:"type:varchar(200);not null"`
	Address     string `gorm:"type:varchar(200)"`
	LicenseId   string `gorm:"type:varchar(20)"`
	PhoneNumber string `gorm:"type:varchar(20)"`
	Email       string `gorm:"type:varchar(200)"`
	Password    string `gorm:"type:varchar(200)"`
	Active      bool   `gorm:"type:bool"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt `gorm:"index"`
}
