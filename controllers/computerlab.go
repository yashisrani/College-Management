package controllers

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/yashisrani/Go-Backend/model"
	"github.com/yashisrani/Go-Backend/utils"
)

func (server Server) CreateComputerLab(ctx *gin.Context) {

	utils.Log(model.LogLevelError, model.ControllersPackage, model.CreateComputerLab, "Creating new ComputerLab", nil)

	ComputerLab := model.ComputerLab{}
	err := ctx.ShouldBindJSON(&ComputerLab) // to bind data in json
	if err != nil {
		utils.Log(model.LogLevelError, model.ControllersPackage, model.CreateComputerLab, "error while json binding", err)
		ctx.JSON(http.StatusInternalServerError, err) // to return status code
		return
	}

	ComputerLab.ID = uuid.New()        // to generate new id
	ComputerLab.CreatedAt = time.Now() // to add created time in ComputerLab

	err = server.PostgressDB.CreateComputerLab(&ComputerLab)
	if err != nil {
		utils.Log(model.LogLevelError, model.ControllersPackage, model.CreateComputerLab, "error while inserting ComputerLab to database", err)
		ctx.JSON(http.StatusInternalServerError, err) // to return status code
		return
	}

	ctx.JSON(http.StatusCreated, ComputerLab)
}

func (server Server) GetComputerLab(ctx *gin.Context) {

	utils.Log(model.LogLevelError, model.ControllersPackage, model.GetComputerLab, "Fetching ComputerLabs", nil)

	ComputerLab, err := server.PostgressDB.GetComputerLab()
	if err != nil {
		utils.Log(model.LogLevelError, model.ControllersPackage, model.GetComputerLab, "error while fetching ComputerLabs from database", err)
		ctx.JSON(http.StatusInternalServerError, err) // to return status code
		return
	}

	ctx.JSON(http.StatusCreated, ComputerLab)

}

func (server Server) GetComputerLabByID(ctx *gin.Context) {

	id := ctx.Param("id")

	utils.Log(model.LogLevelInfo, model.ControllersPackage, model.GetComputerLabByID, "fetching ComputerLab by id", map[string]interface{}{"id": id})

	uuid, err := uuid.Parse(id)
	if err != nil {
		utils.Log(model.LogLevelError, model.ControllersPackage, model.GetComputerLabByID, "invalid ComputerLab ID format", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ComputerLab ID format"})
		return
	}

	ComputerLab, err := server.PostgressDB.GetComputerLabByID(uuid)
	if err != nil {
		utils.Log(model.LogLevelError, model.ControllersPackage, model.GetComputerLabByID, "error while inserting record ", err)
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}
	ctx.JSON(http.StatusCreated, ComputerLab)
}

func (server Server) GetComputerLabByFilter(ctx *gin.Context) {

	utils.Log(model.LogLevelInfo, model.ControllersPackage, model.GetComputerLabByFilter, "fetching ComputerLab by filter", nil)

	queryparams := ctx.Request.URL.Query()

	filter := utils.ConvertQueryParams(queryparams)

	ComputerLab, err := server.PostgressDB.GetComputerLabByFilter(filter)

	// uuid, err := uuid.Parse(id)
	// if err != nil {
	// 	utils.Log(model.LogLevelError, model.ControllersPackage, model.GetComputerLabByID, "invalid ComputerLab ID format", err)
	// 	ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ComputerLab ID format"})
	// 	return
	// }

	// ComputerLab, err = server.PostgressDB.GetComputerLabByID(uuid)
	if err != nil {
		utils.Log(model.LogLevelError, model.ControllersPackage, model.GetComputerLabByFilter, "error while fetching record ", err)
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}
	ctx.JSON(http.StatusCreated, ComputerLab)
}

func (server *Server) UpdateComputerLab(c *gin.Context) error {

	var ComputerLab model.ComputerLab
	//Unmarshal
	utils.Log(model.LogLevelInfo, model.ControllersPackage, model.UpdateComputerLab,
		"unmarshaling ComputerLab data", nil)
	err := c.ShouldBindJSON(&ComputerLab)
	if err != nil {
		utils.Log(model.LogLevelError, model.ControllersPackage, model.UpdateComputerLab,
			"error while unmarshaling payload", err)
		return fmt.Errorf("")
	}
	//validation is to be done here
	//DB call
	ComputerLab.UpdatedAt = time.Now().UTC()
	err = server.PostgressDB.UpdateComputerLab(&ComputerLab)
	if err != nil {
		utils.Log(model.LogLevelError, model.ControllersPackage, model.UpdateComputerLab,
			"error while updating record from pgress", err)
		return fmt.Errorf("")
	}
	utils.Log(model.LogLevelInfo, model.ControllersPackage, model.GetComputerLab,
		"successfully updated ComputerLab record and setting response", ComputerLab)
	c.JSON(http.StatusOK, ComputerLab)
	return nil

}

func (server *Server) DeleteComputerLab(c *gin.Context) error {

	//validation is to be done here
	utils.Log(model.LogLevelInfo, model.ControllersPackage, model.DeleteComputerLab,
		"reading ComputerLab id", nil)
	id := c.Param("id")
	if id == "" {
		utils.Log(model.LogLevelError, model.ControllersPackage, model.DeleteComputerLab,
			"missing ComputerLab id", nil)
		return fmt.Errorf("")
	}
	//DB call
	err := server.PostgressDB.DeleteComputerLab(id)
	if err != nil {
		utils.Log(model.LogLevelError, model.ControllersPackage, model.DeleteComputerLab,
			"error while deleting ComputerLab record from pgress", err)
		return fmt.Errorf("")
	}
	utils.Log(model.LogLevelInfo, model.ControllersPackage, model.DeleteComputerLab,
		"successfully deleted ComputerLab record ", nil)
	return nil

}
