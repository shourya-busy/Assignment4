package handlers

import (
	"assignment4/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func CreateBank(context *gin.Context) {
	var input models.Bank

	if err := context.ShouldBind(&input); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error" : err.Error()})
	}

	bank := models.Bank{
		Name: input.Name,
	}

	savedBank,err := bank.Save()

	if err != nil {
		context.JSON(http.StatusBadRequest,gin.H{"error":err.Error()})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"Bank":savedBank})
	
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

	context.JSON(http.StatusOK, gin.H{"Bank":bank})
}

func UpdateBank(context *gin.Context) {
	var input models.Bank

	if err := context.ShouldBind(&input); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error" : err.Error()})
	}

	updatedBank,err := input.Update()

	if err != nil {
		context.JSON(http.StatusBadRequest,gin.H{"error":err.Error()})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"Bank":updatedBank})
	
}




