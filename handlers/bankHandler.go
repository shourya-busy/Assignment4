package handlers

import (
	"assignment4/database"
	"assignment4/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func CreateBank(context *gin.Context) {
	var input models.Bank

	if err := context.ShouldBind(&input); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error" : err.Error()})
		return
	}

	if err := input.Validate(); err != nil {
        context.JSON(http.StatusBadRequest, gin.H{"error" : err.Error()})
		return
    }

	savedBank,err := input.Save()

	if err != nil {
		context.JSON(http.StatusBadRequest,gin.H{"error":err.Error()})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message" : "Bank has been created successfully","Bank ID":savedBank.ID})
	
}


func GetAllBanks(context *gin.Context) {

	banks,err := models.FindAllBanks()

	if err != nil {
		context.JSON(http.StatusBadRequest,gin.H{"error":err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"Banks":banks})
}

func GetBankByID(context *gin.Context) {
	id := context.Param("id")
	ID,_ := strconv.ParseUint(id,10,64)
	bank,err := models.FindBankByID(uint(ID))

	if err != nil {
		context.JSON(http.StatusBadRequest,gin.H{"error":err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"Bank":bank})
}

func DeleteAllBanks(context *gin.Context) {

	err := models.DeleteAllBanks()

	if err != nil {
		context.JSON(http.StatusBadRequest,gin.H{"error":err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message":"All rows have been deleted"})
}

func DeleteBankByID(context *gin.Context) {
	id := context.Param("id")
	ID,_ := strconv.ParseUint(id,10,0)
	bank,err := models.DeleteBankByID(uint(ID))

	if err != nil {
		context.JSON(http.StatusBadRequest,gin.H{"error":err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message" : "Bank has been deleted successfully","Bank ID":bank.ID})
}

func UpdateBank(context *gin.Context) {
	var input models.Bank

	if err := context.ShouldBind(&input); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error" : err.Error()})
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

	updatedBank,err := input.Update(tx)

	if err != nil {
		context.JSON(http.StatusBadRequest,gin.H{"error":err.Error()})
		return
	}

	tx.Commit()

	context.JSON(http.StatusCreated, gin.H{"message" : "Bank has been updated successfully","Bank ID":updatedBank.ID})
	
}




