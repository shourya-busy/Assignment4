package models

import (
	"assignment4/database"
	"errors"

	"github.com/go-pg/pg/v10"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

func (mapping *CustomerToAccount) Validate() error {
    validate := validator.New()
    return validate.Struct(mapping)
}

func (mapping *CustomerToAccount) Save(tx *pg.Tx) error {
	_, insertErr := tx.Model(mapping).Returning("*").Insert()

	if insertErr != nil {
		tx.Rollback()
		return insertErr
	}

	return nil
}


func DeleteNomineeFromAccountByID(accNumber uuid.UUID,id uint) error {

	account,err := FindAccountByAccountNumber(accNumber)

	if err != nil {
		return err
	}

    count, countErr := database.Db.Model((*CustomerToAccount)(nil)).
        Where("account_id = ?", account.ID).
        Count()
    if countErr != nil {
        return countErr
    }

    // Check if there is only one customer mapped to the account
    if count <= 1 {
        return errors.New("the account must have atleast one customer")
    }

	var mapping CustomerToAccount
	_, deleteErr := database.Db.Model(&mapping).Where("customer_id=?",id).Where("account_id=?",account.ID).Delete()
	if deleteErr != nil {
		return deleteErr
	}

	return nil
}
