package httpmodels

type UserRequest struct {
	FiscalNumber  string `json:"fiscal_number" binding:"required"`
	Name     	  string `json:"name" binding:"required"`	
	Birthdate     string `json:"birthdate"`	
}
