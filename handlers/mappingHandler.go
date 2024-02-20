package handlers

import (
	"assignment4/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type AddNomineeRequest struct{
	NomineeID uint
	AccountNumber uuid.UUID
}

func AddNominee(context *gin.Context){
	var input AddNomineeRequest

	if err := context.ShouldBind(&input); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error" : err.Error()})
		return
	}

	account,err := models.FindAccountByAccountNumber(input.AccountNumber)

	if err != nil {
		context.JSON(http.StatusBadRequest,gin.H{"error":err.Error()})
		return
	}

	mapping := models.CustomerToAccount {
		CustomerID : input.NomineeID,
		AccountID: account.ID,
	}

	err = mapping.Save()

	if err != nil {
		context.JSON(http.StatusBadRequest,gin.H{"err":err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"Message":"Nominee added"})

}


func DeleteNomineeFromAccountByID(context *gin.Context) {

	number := uuid.MustParse(context.Param("number"))
	id := context.Param("id")
	ID,_ := strconv.ParseUint(id,10,0)

	err := models.DeleteNomineeFromAccountByID(number, uint(ID))

	if err != nil {
		context.JSON(http.StatusBadRequest,gin.H{"error":err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"Message":"Nominee deleted"})
}

