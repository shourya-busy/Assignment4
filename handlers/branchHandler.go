package handlers

import (
	"assignment4/database"
	"assignment4/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func CreateBranch(context *gin.Context) {
	var input models.Branch

	if err := context.ShouldBind(&input); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error" : err.Error()})
		return
	}

	if err := input.Validate(); err != nil {
        context.JSON(http.StatusBadRequest, gin.H{"error" : err.Error()})
		return
    }

	savedBranch,err := input.Save()

	if err != nil {
		context.JSON(http.StatusBadRequest,gin.H{"err":err.Error()})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message" : "Branch has been created successfully","Branch ID":savedBranch.ID})
	
}

func GetAllBranchesByBankID(context *gin.Context) {

	id := context.Param("id")
	ID,_ := strconv.ParseUint(id,10,0)


	branches,err := models.FindAllBranchesByBankID(uint(ID))

	if err != nil {
		context.JSON(http.StatusBadRequest,gin.H{"error":err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"Branches":branches})
}

func GetBranchByID(context *gin.Context) {

	id := context.Param("id")
	ID,_ := strconv.ParseUint(id,10,0)
	branch,err := models.FindBranchByID(uint(ID))

	if err != nil {
		context.JSON(http.StatusBadRequest,gin.H{"error":err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"Branches":branch})
}

func DeleteAllBranches(context *gin.Context) {

	err := models.DeleteAllBranches()

	if err != nil {
		context.JSON(http.StatusBadRequest,gin.H{"error":err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message":"All branches have been deleted"})
}

func DeleteBranchByID(context *gin.Context) {
	id := context.Param("id")
	ID,_ := strconv.ParseUint(id,10,0)
	branches,err := models.DeleteBranchByID(uint(ID))

	if err != nil {
		context.JSON(http.StatusBadRequest,gin.H{"error":err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message" : "Branch has been deleted successfully","Branches ID":branches.ID})
}

func UpdateBranch(context *gin.Context) {
	var input models.Branch

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

	updatedBranch,err := input.Update(tx)

	if err != nil {
		context.JSON(http.StatusBadRequest,gin.H{"error":err.Error()})
		return
	}

	tx.Commit()

	context.JSON(http.StatusCreated, gin.H{"message":"Branch has been updated successfully","Branch ID":updatedBranch.ID})
	
}



