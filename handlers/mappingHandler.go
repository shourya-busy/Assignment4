package handlers

import (
	"assignment4/database"
	"assignment4/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)



func AddNominee(context *gin.Context){
	var input AddNomineeRequest

	if err := context.ShouldBind(&input); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error" : err.Error()})
		return
	}

	if err := input.Validate(); err != nil {
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

	tx, txErr := database.Db.Begin()
	if txErr != nil {
		context.JSON(http.StatusBadRequest,gin.H{"error" : txErr.Error()})
		return
	}

	err = mapping.Save(tx)

	if err != nil {
		context.JSON(http.StatusBadRequest,gin.H{"err":err.Error()})
		return
	}

	tx.Commit()

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

