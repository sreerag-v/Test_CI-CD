package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	// ID         uint   `json:"id" gorm:"primaryKey;unique"  `
	First_Name string `json:"first_name"  gorm:"not null" validate:"required,min=2,max=50"  `
	Last_Name  string `json:"last_name"    gorm:"not null"    validate:"required,min=1,max=50"  `
	Email      string `json:"email"   gorm:"not null;unique"  validate:"email,required"`
	Password   string `json:"password" gorm:"not null"  validate:"required"`
}
