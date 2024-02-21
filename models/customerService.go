package models

import (
	"assignment4/database"
	"errors"

	"github.com/go-pg/pg/v10"
	"github.com/go-pg/pg/v10/orm"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

func (customer *Customer) Validate() error {
    validate := validator.New()
    return validate.Struct(customer)
}

func (customer *Customer) Save(tx *pg.Tx) (*Customer, error) {
	_, insertErr := tx.Model(customer).Returning("*").Insert()

	if insertErr != nil {
		tx.Rollback()
		return nil,insertErr
	}

	return customer, nil
}

func FindCustomerByID(id uint) (*Customer, error){
	var output Customer
	getErr := database.Db.Model(&output).
		Where("id = ?",id).
		Select()


	if getErr != nil {
		return &Customer{},getErr
	}

	return &output,nil

}

func FindAllCustomers() ([]Customer,error) {
	var customer []Customer
	getErr := database.Db.Model(&customer).
		Select()


	if getErr != nil {
		return nil,getErr
	}

	
	return customer,nil
}

func FindAllCustomersByBranchID(id uint) ([]Customer, error) {
	var customer []Customer
	getErr := database.Db.Model(&customer).
		Where("branch_id = ?",id).
		Select()


	if getErr != nil {
		return nil,getErr
	}

	return customer,nil

}



func DeleteAllCustomers()  error {
	var customer Customer

	opts := &orm.DropTableOptions{
		IfExists: true,
		Cascade: true,
	}

	deleteErr := database.Db.Model(&customer).DropTable(opts)
	if deleteErr != nil {
		return deleteErr
	}

	return nil
}

func DeleteCustomersByID(id uint, tx *pg.Tx) (*Customer, error) {

	accounts,err := FindAllAccountsByCustomerID(id)

	if err != nil {
		return nil,err
	}

	// for _, account := range accounts{

	// 	_, updateErr := tx.Model((*Transaction)(nil)).Set("account_id = NULL").Where("account_id=?",account.ID).Update()
	// 	if updateErr != nil {
	// 		tx.Rollback()
	// 		return nil,updateErr
	// 	}
	// }

	for _, account := range accounts{
		_, deleteErr := tx.Model(account).Where("id=?",account.ID).Where("account_type != 'joint' ").Delete()
		if deleteErr != nil {
			tx.Rollback()
			return nil,deleteErr
		}
	}

	var customer Customer
	_, deleteErr := tx.Model(&customer).Where("id = ?",id).Returning("*").Delete(&customer)

	if deleteErr != nil {
		tx.Rollback()
		return nil,deleteErr
	}

	return &customer,nil
}

func (customer *Customer) Update(tx *pg.Tx) (*Customer, error)  {
	updateResult, updateErr := tx.Model(customer).WherePK().Returning("*").UpdateNotZero(customer)

	if updateErr != nil {
		tx.Rollback()
		return nil,updateErr
	}

	if updateResult.RowsAffected() == 0 {
		tx.Rollback()
		return nil, errors.New("no record updated")
	}

	return customer,nil
}


func FindAllCustomersByAccountNumber(accNumber uuid.UUID) ([]*Customer, error) {

	var account Account
	getErr := database.Db.Model(&account).
		Relation("Customer").
		Where("account_number =?",accNumber).
		Select()


	if getErr != nil {
		return nil,getErr
	}

	return account.Customer,nil

}