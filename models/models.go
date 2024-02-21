package models

import (
	"time"

	"github.com/google/uuid"
)

type Bank struct {
	ID uint `validate:"omitempty,gte=0,numeric"`
	Name string `validate:"omitempty,alpha"`
	Branch []*Branch `pg:"rel:has-many"`
}

type Branch struct {
	ID uint `validate:"omitempty,gte=0,numeric"`
	Address string `validate:"omitempty,alphanum"` //no spaces allowed
	BankID uint `pg:"on_delete:CASCADE" validate:"omitempty,number"`
	Bank *Bank `pg:"rel:has-one"`
	IFSC_CODE uuid.UUID `pg:"type:uuid" validate:"omitempty,uuid"`
	Account []*Account `pg:"rel:has-many"`
	Customer []*Customer `pg:"rel:has-many"`
}

type Customer struct{
	ID uint `validate:"omitempty,gte=0,numeric"`
	BranchID uint `pg:"on_delete:CASCADE" validate:"omitempty,number"`
	Branch *Branch `pg:"rel:has-one"`
	Name string `validate:"omitempty,alpha"`
	PAN string `validate:"omitempty,alphanum"` 
	DOB string `pg:"type:date" ` //validate:"datetime"`
	Age uint `validate:"omitempty,gte=0,lte=130"`
	Phone uint `validate:"omitempty,number"`
	Address string `validate:"omitempty,alphanum"`
	Account []*Account `pg:"many2many:customer_to_accounts"`
}

type Account struct{
	ID uint `validate:"omitempty,gte=0,numeric"`
	BranchID uint `pg:"on_delete:CASCADE" validate:"omitempty,number"`
	Branch *Branch `pg:"rel:has-one"`
	AccountNumber uuid.UUID `pg:"type:uuid" validate:"omitempty,uuid"`
	Balance float64 `validate:"omitempty,gte=0,numeric"`
	AccountType string `validate:"omitempty,oneof=savings current joint"`
	Customer []*Customer `pg:"many2many:customer_to_accounts"`
	Transaction []*Transaction `pg:"rel:has-many"`
}

type CustomerToAccount struct{
	ID uint `validate:"omitempty,gte=0,numeric"`
	AccountID uint `pg:"on_delete:CASCADE " validate:"omitempty,number"`
	CustomerID uint `pg:"on_delete:CASCADE" validate:"omitempty,number"`

	Customer *Customer `pg:"rel:has-one"`
    Account  *Account  `pg:"rel:has-one"`
}

type Transaction struct{
	ID uint `validate:"omitempty,gte=0,numeric"`
	AccountID uint `pg:"on_delete:SET NULL" validate:"omitempty,number"`
	Account *Account `pg:"rel:has-one"`
	ReceiverAccountNumber uuid.UUID `pg:"type:uuid" validate:"omitempty,uuid"`
	ModeOfPayment string `validate:"omitempty,oneof=cash card UPI"`
	TypeOfTransaction string `validate:"omitempty,oneof=Deposit Withdraw Transafer"`
	Amount float64 `validate:"omitempty,gte=0,numeric"`
	Time time.Time `validate:"omitempty,datetime"`
}