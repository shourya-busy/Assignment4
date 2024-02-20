package handlers

import (
	"assignment4/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type CreateAccountRequest struct {
	CustomerID uint
	Balance float64
	AccountType string
	NomineeID uint
}

func CreateAccount(context *gin.Context) {
	var input CreateAccountRequest

	if err := context.ShouldBind(&input); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error" : err.Error()})
		return
	}

	customer, err := models.FindCustomerByID(input.CustomerID)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error" : err.Error()})
		return
	}

	account := models.Account{
		BranchID: customer.BranchID,
		Balance: input.Balance,
		AccountType: input.AccountType,
	}

	account.Customer = append(account.Customer, customer)
	
	if input.NomineeID != 0 {
		nominee, err := models.FindCustomerByID(input.NomineeID)
		
		if err != nil {
			context.JSON(http.StatusBadRequest, gin.H{"error" : err.Error()})
			return
		}
		
		account.Customer = append(account.Customer, nominee)
	}
	
	savedAccount,err := account.Save()
	
	if err != nil {
		context.JSON(http.StatusBadRequest,gin.H{"error":err.Error()})
		return
    }

	mapping := models.CustomerToAccount {
		CustomerID : customer.ID,
		AccountID: savedAccount.ID,
	}

	err = mapping.Save()

	if err != nil {
		context.JSON(http.StatusBadRequest,gin.H{"err":err.Error()})
		return
	}

	if input.NomineeID != 0 {

		mapping := models.CustomerToAccount {
			CustomerID : input.NomineeID,
			AccountID: savedAccount.ID,
		}
	
		err = mapping.Save()
	
		if err != nil {
			context.JSON(http.StatusBadRequest,gin.H{"err":err.Error()})
			return
		}

	}

	context.JSON(http.StatusCreated, gin.H{"Account":savedAccount})
}

func GetAllAccountsByBranchID(context *gin.Context) {
	id := context.Param("id")
	ID,_ := strconv.ParseUint(id,10,0)
	accounts,err := models.FindAllAccountsByBranchID(uint(ID))

	if err != nil {
		context.JSON(http.StatusBadRequest,gin.H{"error":err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"Account":accounts})
}

func GetAccountById(context *gin.Context) {

	id := context.Param("id")
	ID,_ := strconv.ParseUint(id,10,0)
	account,err := models.FindAccountByID(uint(ID))

	if err != nil {
		context.JSON(http.StatusBadRequest,gin.H{"error":err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"Account":account})
}

func GetAllAccountsByCustomerID(context *gin.Context) {
	id := context.Param("id")
	ID,_ := strconv.ParseUint(id,10,0)

	accounts,err := models.FindAllAccountsByCustomerID(uint(ID))

	if err != nil {
		context.JSON(http.StatusBadRequest,gin.H{"error":err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"Account":accounts})
}

func GetAccountByAccountNumber(context *gin.Context) {
	id := uuid.MustParse(context.Param("number"))


	account,err := models.FindAccountByAccountNumber(id)

	if err != nil {
		context.JSON(http.StatusBadRequest,gin.H{"error":err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"Account":account})
}

func DeleteAllAccounts(context *gin.Context) {

	err := models.DeleteAllAccounts()

	if err != nil {
		context.JSON(http.StatusBadRequest,gin.H{"error":err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message":"All tables have been deleted"})
}

func DeleteAccountByID(context *gin.Context) {
	id := context.Param("id")
	ID,_ := strconv.ParseUint(id,10,0)
	branches,err := models.DeleteAccountByID(uint(ID))

	if err != nil {
		context.JSON(http.StatusBadRequest,gin.H{"error":err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"Branches":branches})
}

func UpdateAccount(context *gin.Context) {
	var input models.Account

	if err := context.ShouldBind(&input); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error" : err.Error()})
	}

	updatedAccount,err := input.Update()

	if err != nil {
		context.JSON(http.StatusBadRequest,gin.H{"error":err.Error()})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"Account":updatedAccount})
	
}

