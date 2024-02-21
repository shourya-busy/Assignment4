package models

import (
	"assignment4/database"
	"errors"

	"github.com/go-pg/pg/v10"
	"github.com/go-pg/pg/v10/orm"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

const MIN_BALANCE = 2000.00

func (transaction *Transaction) Validate() error {
    validate := validator.New()
    return validate.Struct(transaction)
}

func (transaction *Transaction) Save(tx *pg.Tx) (*Transaction, error) {
	_, insertErr := tx.Model(transaction).Returning("*").Insert()

	if insertErr != nil {
		return nil,insertErr
	}

	return transaction,nil
}

func AccountDeposit(accountID uint, amount float64, tx *pg.Tx) error {


	var account Account
	updateResult, updateErr := tx.Model(&account).
		Set("balance = balance + ?",amount).
		Where("id = ?",accountID).
		Returning("*").
		Update(&account)

	if updateErr != nil {
		tx.Rollback()
		return updateErr
	}

	if updateResult.RowsAffected() == 0 {
		tx.Rollback()
		return errors.New("account does not exists")
	}

	return nil
}

func AccountWithdrawal(accountID uint, amount float64, tx *pg.Tx) error {


	exists, err := tx.Model(&Account{}).Where("id = ?", accountID).Exists()
    if err != nil {
        tx.Rollback()
        return err
    }
    if !exists {
        tx.Rollback()
        return errors.New("account does not exist")
    }

	var account Account
	updateResult, updateErr := tx.Model(&account).
		Set("balance = balance - ?",amount).
		Where("id = ?",accountID).
		Where("(balance - ?) >= ?",amount,MIN_BALANCE).
		Returning("*").
		Update(&account)

	if updateErr != nil {
		tx.Rollback()
		return updateErr
	}

	if updateResult.RowsAffected() == 0 {
		tx.Rollback()
		return errors.New("insufficient balance")
	}

	return nil
}

func AccountTransfer(accountID uint,receiverAccountNo uuid.UUID,amount float64, tx *pg.Tx) error {


	err := AccountWithdrawal(accountID, amount,tx)

	if err != nil {
		tx.Rollback()
		return err
	}

	var receiver Account
	updateResult, updateErr := tx.Model(&receiver).
		Set("balance = balance + ?",amount).
		Where("account_number = ?",receiverAccountNo).
		Returning("*").
		Update(&receiver)

	if updateErr != nil {
		tx.Rollback()
		return updateErr
	}

	if updateResult.RowsAffected() == 0 {
		tx.Rollback()
		return errors.New("account does not exists")
	}

	return nil
}

func FindAllTransactions() ([]Transaction,error) {
	var transactions []Transaction
	getErr := database.Db.Model(&transactions).
		Select()


	if getErr != nil {
		return nil,getErr
	}

	return transactions,nil
}


func FindTransactionByID(id uint) (*Transaction, error){
	var output Transaction
	getErr := database.Db.Model(&output).
		Where("id = ?",id).
		Select()


	if getErr != nil {
		return &Transaction{},getErr
	}

	return &output,nil

}

func FindAllTransactionsByAccountNumber(accNumber uuid.UUID) ([]Transaction, error){
	account,err := FindAccountByAccountNumber(accNumber)

	if err != nil {
		return nil, err
	}

	var transactions []Transaction
	getErr := database.Db.Model(&transactions).
		Where("receiver_account_number = ?",accNumber).
		WhereOr("account_id = ?",account.ID).
		Select()


	if getErr != nil {
		return nil,getErr
	}

	
	return transactions,nil

}

func DeleteAllTransactions()  error {
	var transaction Transaction

	opts := &orm.DropTableOptions{
		IfExists: true,
		Cascade: true,
	}

	deleteErr := database.Db.Model(&transaction).DropTable(opts)
	if deleteErr != nil {
		return deleteErr
	}

	return nil
}

func DeleteTransactionByID(id uint) (*Transaction, error) {
	var transaction Transaction
	_, deleteErr := database.Db.Model(&transaction).Where("id=?",id).Returning("*").Delete(&transaction)
	if deleteErr != nil {
		return nil,deleteErr
	}

	return &transaction,nil
}

