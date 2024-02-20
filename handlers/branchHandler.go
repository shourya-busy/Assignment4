package handlers

import (
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

	savedBranch,err := input.Save()

	if err != nil {
		context.JSON(http.StatusBadRequest,gin.H{"err":err.Error()})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"Branch":savedBranch})
	
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

	context.JSON(http.StatusOK, gin.H{"Branches":branches})
}

func UpdateBranch(context *gin.Context) {
	var input models.Branch

	if err := context.ShouldBind(&input); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error" : err.Error()})
	}

	updatedBranch,err := input.Update()

	if err != nil {
		context.JSON(http.StatusBadRequest,gin.H{"error":err.Error()})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"Branch":updatedBranch})
	
}



