package handlers

import (
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

type CreateAccountRequest struct {
	CustomerID uint `validate:"number"`
	Balance float64 `validate:"gte=0,numeric"`
	AccountType string `validate:"oneof=savings current joint"`
	NomineeID uint `validate:"number"`
}

func (request *CreateAccountRequest) Validate() error {
    validate := validator.New()
    return validate.Struct(request)
}

type AddNomineeRequest struct{
	NomineeID uint `validate:"number"`
	AccountNumber uuid.UUID `validate:"uuid"`
}

func (request *AddNomineeRequest) Validate() error {
    validate := validator.New()
    return validate.Struct(request)
}


type CreateCustomerRequest struct{
	BranchID uint `validate:"number"`
	Name string `validate:"alpha"`
	PAN string `validate:"alphanum"` 
	DOB string `pg:"type:date" `// validate:"date"`
	Age uint `validate:"gte=0,lte=130"`
	Phone uint `validate:"number"`
	Address string `validate:"alphanum"`
	Balance float64 `validate:"gte=0,numeric"`
	AccountType string `validate:"oneof=savings current joint"`
}

func (request *CreateCustomerRequest) Validate() error {
    validate := validator.New()
    return validate.Struct(request)
}


