package handlers

import (
	"assignment4/database"
	"assignment4/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func CreateCustomer(context *gin.Context) {
	var input CreateCustomerRequest

	if err := context.ShouldBind(&input); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error" : err.Error()})
		return
	}

	if err := input.Validate(); err != nil {
        context.JSON(http.StatusBadRequest, gin.H{"error" : err.Error()})
		return
    }

	customer := models.Customer{
		BranchID: input.BranchID,
		Name: input.Name,
		PAN: input.PAN,
		DOB: input.DOB,
		Age: input.Age,
		Phone: input.Phone,
		Address: input.Address,
	}

	account := &models.Account{
		BranchID: input.BranchID,
		Balance: input.Balance,
		AccountType: input.AccountType,
	}

	tx, txErr := database.Db.Begin()
	if txErr != nil {
		context.JSON(http.StatusBadRequest,gin.H{"error" : txErr.Error()})
		return
	}

	savedCustomer,err := customer.Save(tx)

	if err != nil {
		context.JSON(http.StatusBadRequest,gin.H{"err":err.Error()})
		return
	}

	savedAccount,err := account.Save(tx)

	
	if err != nil {
		context.JSON(http.StatusBadRequest,gin.H{"err":err.Error()})
		return
	}
	
	savedCustomer.Account = append(savedCustomer.Account, savedAccount)

	mapping := models.CustomerToAccount {
		CustomerID : savedCustomer.ID,
		AccountID: savedAccount.ID,
	}

	err = mapping.Save(tx)

	if err != nil {
		context.JSON(http.StatusBadRequest,gin.H{"err":err.Error()})
		return
	}

	tx.Commit()

	context.JSON(http.StatusCreated, gin.H{"message" : "Customer has been created successfully","Customer ID":savedCustomer.ID})
}

func GetAllCustomersByBranchID(context *gin.Context) {

	id := context.Param("id")
	ID,_ := strconv.ParseUint(id,10,0)

	customer,err := models.FindAllCustomersByBranchID(uint(ID))

	if err != nil {
		context.JSON(http.StatusBadRequest,gin.H{"error":err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"Customer":customer})
}

func GetCustomerByID(context *gin.Context) {

	id := context.Param("id")
	ID,_ := strconv.ParseUint(id,10,0)
	customer,err := models.FindCustomerByID(uint(ID))

	if err != nil {
		context.JSON(http.StatusBadRequest,gin.H{"error":err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"Customer":customer})
}


func DeleteAllCustomers(context *gin.Context) {

	err := models.DeleteAllCustomers()

	if err != nil {
		context.JSON(http.StatusBadRequest,gin.H{"error":err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message":"All Customers have been deleted"})
}

func DeleteCustomerByID(context *gin.Context) {
	id := context.Param("id")
	ID,_ := strconv.ParseUint(id,10,0)

	tx, txErr := database.Db.Begin()
	if txErr != nil {
		context.JSON(http.StatusBadRequest,gin.H{"error" : txErr.Error()})
		return
	}
	customer,err := models.DeleteCustomersByID(uint(ID),tx)

	if err != nil {
		context.JSON(http.StatusBadRequest,gin.H{"error":err.Error()})
		return
	}

	tx.Commit()

	context.JSON(http.StatusOK, gin.H{"message" : "Customer has been deleted successfully","Customer ID":customer.ID})
}

func UpdateCustomer(context *gin.Context) {
	var input models.Customer

	if err := context.ShouldBind(&input); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error" : err.Error()})
		return
	}

	if err := input.Validate(); err != nil {
        context.JSON(http.StatusBadRequest, gin.H{"error" : err.Error()})
		return
    }

	tx, txErr := database.Db.Begin()
	if txErr != nil {
		context.JSON(http.StatusBadRequest,gin.H{"error" : txErr.Error()})
		return
	}

	updatedCustomer,err := input.Update(tx)

	if err != nil {
		context.JSON(http.StatusBadRequest,gin.H{"error":err.Error()})
		return
	}

	tx.Commit()

	context.JSON(http.StatusCreated, gin.H{"message" : "Customer has been updated succesfully","Customer ID":updatedCustomer.ID})
	
}


func GetAllNomineesByAccountNumber(context *gin.Context) {
	number := uuid.MustParse(context.Param("number"))

	customers,err := models.FindAllCustomersByAccountNumber(number)

	if err != nil {
		context.JSON(http.StatusBadRequest,gin.H{"error":err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"Nominees":customers})
}


