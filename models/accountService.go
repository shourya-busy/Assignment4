package models

import (
	"assignment4/database"
	"context"
	"errors"

	"github.com/go-pg/pg/v10"
	"github.com/go-pg/pg/v10/orm"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

func (account *Account) Validate() error {
    validate := validator.New()
    return validate.Struct(account)
}

func (account *Account) Save(tx *pg.Tx) (*Account, error) {
	_, insertErr := tx.Model(account).Returning("*").Insert()

	if insertErr != nil {
		tx.Rollback()
		return nil,insertErr
	}

	return account, nil
}

func (account *Account) BeforeInsert (context context.Context) (context.Context,error) {

	account.AccountNumber = uuid.New()
	return context,nil

}

func FindAccountByID(id uint) (*Account, error){
	var output Account
	getErr := database.Db.Model(&output).
		Where("id = ?",id).
		Select()


	if getErr != nil {
		return &Account{},getErr
	}

	return &output,nil

}


func FindAllAccounts() ([]Account,error) {
	var accounts []Account
	getErr := database.Db.Model(&accounts).
		Select()


	if getErr != nil {
		return nil,getErr
	}

	
	return accounts,nil
}

func FindAllAccountsByBranchID(id uint) ([]Account,error) {
	var accounts []Account
	getErr := database.Db.Model(&accounts).
		Where("branch_id =?",id).
		Select()

	if getErr != nil {
		return nil,getErr
	}

	
	return accounts,nil
}



func FindAccountByAccountNumber(accNumber uuid.UUID) (*Account, error) {
	var account Account
	getErr := database.Db.Model(&account).
		Where("account_number = ?",accNumber).
		Select()


	if getErr != nil {
		return nil,getErr
	}

	
	return &account,nil

}

func FindAllAccountsByCustomerID(id uint) ([]*Account, error) {

	var customer Customer
	getErr := database.Db.Model(&customer).
		Relation("Account").
		Where("id =?",id).
		Select()


	if getErr != nil {
		return nil,getErr
	}

	return customer.Account,nil

}

func DeleteAllAccounts()  error {
	var account Account

	opts := &orm.DropTableOptions{
		IfExists: true,
		Cascade: true,
	}

	deleteErr := database.Db.Model(&account).DropTable(opts)
	if deleteErr != nil {
		return deleteErr
	}

	return nil
}

func DeleteAccountByID(id uint, tx *pg.Tx) (*Account, error) {

	// _, updateErr := tx.Model((*Transaction)(nil)).Set("account_id = NULL").Where("account_id=?",id).Update()
	// if updateErr != nil {
	// 	tx.Rollback()
	// 	return nil,updateErr
	// }

	var account Account
	_, deleteErr := tx.Model(&account).Where("id=?",id).Returning("*").Delete(&account)
	if deleteErr != nil {
		tx.Rollback()
		return nil,deleteErr
	}

	return &account,nil
}

func (account *Account) Update(tx *pg.Tx) (*Account, error)  {


	updateResult, updateErr := tx.Model(account).WherePK().Returning("*").UpdateNotZero(account)

	if updateErr != nil {
		tx.Rollback()
		return nil,updateErr
	}

	if updateResult.RowsAffected() == 0 {
		tx.Rollback()
		return nil, errors.New("no record updated")
	}


	
	return account,nil
}