package models

import (
	"time"
	"github.com/jinzhu/gorm"	
	"github.com/teravision/config"
)

type User struct {
	gorm.Model
	FiscalNumber     string       `json:"-" db:"fiscal_number"`
	Name        	 string       `json:"-" db:"name"`
	Birthdate        time.Time    `json:"-" db:"birthdate"`	
}

func UserExists(fiscalNumber string) (bool) {

	user := User{}

	if config.DB.Where("fiscal_number = ?", fiscalNumber).First(&user).RecordNotFound() {
		return false
	}

	return true
}

func (u *User) Create() error{
	return config.DB.Create(u).Error
}
