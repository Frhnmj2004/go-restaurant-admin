package models

import "github.com/jinzhu/gorm"

type Table struct {
	gorm.Model
	FirstName    *string `json:"firstname" validate:"required,min=3,max=50"`
	LastName     *string `json:"lastname" validate:"required,min=3,max=50"`
	Email        *string `json:"email" validate:"email,required"`
	Phone        *string `json:"phone" validate:"required"`
	Password     *string `json:"password" validate:"required,min=3,max=20"`
	UserType     *string `json:"usertype" validate:"required,eq=ADMIN|eq=USER"`
	Token        *string `json:"token"`
	RefreshToken *string `json:"refreshtoken"`
}
