package helpers

import (
	"fmt"	
	"time"
	"errors"
	"github.com/teravision/models"
	"github.com/teravision/httpmodels"	
)

//Validation and create user information

func CreateUser(u *httpmodels.UserRequest) error{
	
	if models.UserExists(u.FiscalNumber){
		return errors.New("User already exists.")
	}

	date, err := time.Parse("2006-01-02", u.Birthdate)
	if err != nil {
		date, _ = time.Parse("2006-01-02", "0001-01-01")
	}

	user := &models.User{
		FiscalNumber: u.FiscalNumber,
		Name:    u.Name,		
		Birthdate:    date,		
	}

	err = user.Create()
	if err != nil {
		fmt.Println(err)
		return  err
	}

	return nil
}
