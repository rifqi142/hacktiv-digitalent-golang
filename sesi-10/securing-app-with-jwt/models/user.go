package models

import (
	"hacktiv-digitalent-golang/sesi-10/securing-app-with-jwt/helpers"

	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type User struct {
	GormModel
	FullName string    `gorm:"not null" json:"full_name" form:"full_name" valid:"required-Your full name is required!"`
	Email    string    `gorm:"not null;uniqueIndex" json:"email" form:"email" valid:"email,required-Please enter a valid email address!"`
	Password string    `gorm:"not null" json:"password" valid:"required-Password is required, minstringlength(6)-Password should be at least 6 characters long!"`
	Products []Product `gorm:"constraint:OnUpdate:CASCADE, OnDelete:SET NULL;" json:"products" `
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(u)

	if errCreate != nil {
		err = errCreate
		return
	}

	// Hash the password
	u.Password = helpers.HashPass(u.Password)

	err = nil
	return
}