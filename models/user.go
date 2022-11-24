package models

import "gorm.io/gorm"

type User struct {
	gorm.Model

	First_name string `gorm:"not null" json:"firstname"`
	Last_name  string `gorm:"not null" json:"lastname"`
	Email      string `gorm:"unique; not null" json:"email"`
}
