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

func (server Server) CreateClass(ctx *gin.Context) {

	utils.Log(model.LogLevelError, model.ControllersPackage, model.CreateClass, "Creating new Class", nil)

	Class := model.Class{}
	err := ctx.ShouldBindJSON(&Class) // to bind data in json
	if err != nil {
		utils.Log(model.LogLevelError, model.ControllersPackage, model.CreateClass, "error while json binding", err)
		ctx.JSON(http.StatusInternalServerError, err) // to return status code
		return
	}

	Class.ID = uuid.New()        // to generate new id
	Class.CreatedAt = time.Now() // to add created time in Class

	err = server.PostgressDB.CreateClass(&Class)
	if err != nil {
		utils.Log(model.LogLevelError, model.ControllersPackage, model.CreateClass, "error while inserting Class to database", err)
		ctx.JSON(http.StatusInternalServerError, err) // to return status code
		return
	}

	ctx.JSON(http.StatusCreated, Class)
}

func (server Server) GetClass(ctx *gin.Context) {

	utils.Log(model.LogLevelError, model.ControllersPackage, model.GetClass, "Fetching Classs", nil)

	Class, err := server.PostgressDB.GetClass()
	if err != nil {
		utils.Log(model.LogLevelError, model.ControllersPackage, model.GetClass, "error while fetching Classs from database", err)
		ctx.JSON(http.StatusInternalServerError, err) // to return status code
		return
	}

	ctx.JSON(http.StatusCreated, Class)

}

func (server Server) GetClassByID(ctx *gin.Context) {

	id := ctx.Param("id")

	utils.Log(model.LogLevelInfo, model.ControllersPackage, model.GetClassByID, "fetching Class by id", map[string]interface{}{"id": id})

	uuid, err := uuid.Parse(id)
	if err != nil {
		utils.Log(model.LogLevelError, model.ControllersPackage, model.GetClassByID, "invalid Class ID format", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Class ID format"})
		return
	}

	Class, err := server.PostgressDB.GetClassByID(uuid)
	if err != nil {
		utils.Log(model.LogLevelError, model.ControllersPackage, model.GetClassByID, "error while inserting record ", err)
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}
	ctx.JSON(http.StatusCreated, Class)
}

func (server Server) GetClassByFilter(ctx *gin.Context) {

	utils.Log(model.LogLevelInfo, model.ControllersPackage, model.GetClassByFilter, "fetching Class by filter", nil)

	queryparams := ctx.Request.URL.Query()

	filter := utils.ConvertQueryParams(queryparams)

	Class, err := server.PostgressDB.GetClassByFilter(filter)

	// uuid, err := uuid.Parse(id)
	// if err != nil {
	// 	utils.Log(model.LogLevelError, model.ControllersPackage, model.GetClassByID, "invalid Class ID format", err)
	// 	ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Class ID format"})
	// 	return
	// }

	// Class, err = server.PostgressDB.GetClassByID(uuid)
	if err != nil {
		utils.Log(model.LogLevelError, model.ControllersPackage, model.GetClassByFilter, "error while fetching record ", err)
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}
	ctx.JSON(http.StatusCreated, Class)
}

func (server *Server) UpdateClass(c *gin.Context) error {

	var Class model.Class
	//Unmarshal
	utils.Log(model.LogLevelInfo, model.ControllersPackage, model.UpdateClass,
		"unmarshaling Class data", nil)
	err := c.ShouldBindJSON(&Class)
	if err != nil {
		utils.Log(model.LogLevelError, model.ControllersPackage, model.UpdateClass,
			"error while unmarshaling payload", err)
		return fmt.Errorf("")
	}
	//validation is to be done here
	//DB call
	Class.UpdatedAt = time.Now().UTC()
	err = server.PostgressDB.UpdateClass(&Class)
	if err != nil {
		utils.Log(model.LogLevelError, model.ControllersPackage, model.UpdateClass,
			"error while updating record from pgress", err)
		return fmt.Errorf("")
	}
	utils.Log(model.LogLevelInfo, model.ControllersPackage, model.GetClass,
		"successfully updated Class record and setting response", Class)
	c.JSON(http.StatusOK, Class)
	return nil

}

func (server *Server) DeleteClass(c *gin.Context) error {

	//validation is to be done here
	utils.Log(model.LogLevelInfo, model.ControllersPackage, model.DeleteClass,
		"reading Class id", nil)
	id := c.Param("id")
	if id == "" {
		utils.Log(model.LogLevelError, model.ControllersPackage, model.DeleteClass,
			"missing Class id", nil)
		return fmt.Errorf("")
	}
	//DB call
	err := server.PostgressDB.DeleteClass(id)
	if err != nil {
		utils.Log(model.LogLevelError, model.ControllersPackage, model.DeleteClass,
			"error while deleting Class record from pgress", err)
		return fmt.Errorf("")
	}
	utils.Log(model.LogLevelInfo, model.ControllersPackage, model.DeleteClass,
		"successfully deleted Class record ", nil)
	return nil

}
