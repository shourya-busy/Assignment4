package handlers

import (
	"assignment4/database"
	"assignment4/models"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func Deposit(context *gin.Context) {
	var input models.Transaction
	if err := context.ShouldBind(&input); err != nil {
		context.JSON(http.StatusBadRequest,gin.H{"error" : err.Error()})
		return
	}

	if err := input.Validate(); err != nil {
        context.JSON(http.StatusBadRequest, gin.H{"error" : err.Error()})
		return
    }

	transaction := models.Transaction{
		AccountID: input.AccountID,
		Amount: input.Amount,
		ModeOfPayment: input.ModeOfPayment,
		TypeOfTransaction: "Deposit",
		Time: time.Now(),
	}

	tx, txErr := database.Db.Begin()
	if txErr != nil {
		context.JSON(http.StatusBadRequest,gin.H{"error" : txErr.Error()})
		return
	}

	err := models.AccountDeposit(transaction.AccountID,transaction.Amount,tx)

	if err != nil {
		context.JSON(http.StatusNotModified,gin.H{"error" : err.Error()})
		return
	}

	savedTransaction,err := transaction.Save(tx)

	if err != nil {
		context.JSON(http.StatusResetContent,gin.H{"error" : err.Error()})
		return 
	}

	tx.Commit()

	context.JSON(http.StatusAccepted,gin.H{"message":"Your Transaction has been completed suxccessfully","data":savedTransaction.ID})

}

func Withdraw(context *gin.Context) {
	var input models.Transaction
	if err := context.ShouldBind(&input); err != nil {
		context.JSON(http.StatusBadRequest,gin.H{"error" : err.Error()})
		return
	}

	if err := input.Validate(); err != nil {
        context.JSON(http.StatusBadRequest, gin.H{"error" : err.Error()})
		return
    }

	transaction := models.Transaction{
		AccountID: input.AccountID,
		Amount: input.Amount,
		ModeOfPayment: input.ModeOfPayment,
		TypeOfTransaction: "Withdraw",
		Time: time.Now(),
	}

	tx, txErr := database.Db.Begin()
	if txErr != nil {
		context.JSON(http.StatusBadRequest,gin.H{"error" : txErr.Error()})
		return
	}

	err := models.AccountWithdrawal(transaction.AccountID,transaction.Amount,tx)

	if err != nil {
		context.JSON(http.StatusBadGateway,gin.H{"error" : err.Error()})
		return
	}

	savedTransaction,err := transaction.Save(tx)

	if err != nil {
		context.JSON(http.StatusResetContent,gin.H{"error" : err.Error()})
		return
	}

	tx.Commit()

	context.JSON(http.StatusAccepted,gin.H{"message":"Your Transaction has been completed suxccessfully","data":savedTransaction.ID})

}

func Transfer(context *gin.Context) {
	var input models.Transaction
	if err := context.ShouldBind(&input); err != nil {
		context.JSON(http.StatusBadRequest,gin.H{"error" : err.Error()})
		return
	}

	if err := input.Validate(); err != nil {
        context.JSON(http.StatusBadRequest, gin.H{"error" : err.Error()})
		return
    }


	transaction := models.Transaction{
		AccountID: input.AccountID,
		Amount: input.Amount,
		ModeOfPayment: input.ModeOfPayment,
		TypeOfTransaction: "Transfer",
		ReceiverAccountNumber: input.ReceiverAccountNumber,
		Time: time.Now(),
	}

	tx, txErr := database.Db.Begin()
	if txErr != nil {
		context.JSON(http.StatusBadRequest,gin.H{"error" : txErr.Error()})
		return
	}

	err := models.AccountTransfer(transaction.AccountID,transaction.ReceiverAccountNumber,transaction.Amount,tx)

	if err != nil {
		context.JSON(http.StatusBadRequest,gin.H{"error" : err.Error()})
		return
	}

	savedTransaction,err := transaction.Save(tx)

	if err != nil {
		context.JSON(http.StatusResetContent,gin.H{"error" : err.Error()})
		return
	}

	tx.Commit()

	context.JSON(http.StatusAccepted,gin.H{"message":"Your Transaction has been completed suxccessfully","data":savedTransaction.ID})

}

func GetAllTransactionsByAccountNumber(context *gin.Context) {

	number := uuid.MustParse(context.Param("number"))

	transactions,err := models.FindAllTransactionsByAccountNumber(number)

	if err != nil {
		context.JSON(http.StatusBadRequest,gin.H{"error":err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"Transactions":transactions})
}


func GetTransactionByID(context *gin.Context) {
	id := context.Param("id")
	ID,_ := strconv.ParseUint(id,10,0)
	transaction,err := models.FindTransactionByID(uint(ID))

	if err != nil {
		context.JSON(http.StatusBadRequest,gin.H{"error":err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"Transaction":transaction})
}


