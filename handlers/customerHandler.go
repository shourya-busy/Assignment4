package handlers

import (
	"assignment4/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type CreateCustomerRequest struct{
	BranchID uint
	Name string
	PAN string
	DOB string
	Age uint
	Phone uint
	Address string
	Balance float64
	AccountType string
}

func CreateCustomer(context *gin.Context) {
	var input CreateCustomerRequest

	if err := context.ShouldBind(&input); err != nil {
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

	savedCustomer,err := customer.Save()

	if err != nil {
		context.JSON(http.StatusBadRequest,gin.H{"err":err.Error()})
		return
	}

	savedAccount,err := account.Save()

	
	if err != nil {
		context.JSON(http.StatusBadRequest,gin.H{"err":err.Error()})
		return
	}
	
	savedCustomer.Account = append(savedCustomer.Account, savedAccount)

	mapping := models.CustomerToAccount {
		CustomerID : savedCustomer.ID,
		AccountID: savedAccount.ID,
	}

	err = mapping.Save()

	if err != nil {
		context.JSON(http.StatusBadRequest,gin.H{"err":err.Error()})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"Customer":savedCustomer})
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
	customer,err := models.DeleteCustomersByID(uint(ID))

	if err != nil {
		context.JSON(http.StatusBadRequest,gin.H{"error":err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"Customer":customer})
}

func UpdateCustomer(context *gin.Context) {
	var input models.Customer

	if err := context.ShouldBind(&input); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error" : err.Error()})
		return
	}

	updatedCustomer,err := input.Update()

	if err != nil {
		context.JSON(http.StatusBadRequest,gin.H{"error":err.Error()})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"Customer":updatedCustomer})
	
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


