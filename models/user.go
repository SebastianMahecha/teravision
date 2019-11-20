package models

import (
	"time"
	"github.com/jinzhu/gorm"	
)

type User struct {
	gorm.Model
	FiscalNumber     string       `json:"-" db:"fiscal_number"`
	Name        	 string       `json:"-" db:"name"`
	Birthdate        time.Time    `json:"-" db:"birthdate"`	
}
