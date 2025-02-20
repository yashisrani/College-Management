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

func (server Server) CreateCollege(ctx *gin.Context) {

	utils.Log(model.LogLevelError, model.ControllersPackage, model.CreateCollege, "Creating new College", nil)

	College := model.College{}
	err := ctx.ShouldBindJSON(&College) // to bind data in json
	if err != nil {
		utils.Log(model.LogLevelError, model.ControllersPackage, model.CreateCollege, "error while json binding", err)
		ctx.JSON(http.StatusInternalServerError, err) // to return status code
		return
	}

	College.ID = uuid.New()        // to generate new id
	College.CreatedAt = time.Now() // to add created time in College

	err = server.PostgressDB.CreateCollege(&College)
	if err != nil {
		utils.Log(model.LogLevelError, model.ControllersPackage, model.CreateCollege, "error while inserting College to database", err)
		ctx.JSON(http.StatusInternalServerError, err) // to return status code
		return
	}

	ctx.JSON(http.StatusCreated, College)
}

func (server Server) GetColleges(ctx *gin.Context) {

	utils.Log(model.LogLevelError, model.ControllersPackage, model.GetCollege, "Fetching Colleges", nil)

	College, err := server.PostgressDB.GetColleges()
	if err != nil {
		utils.Log(model.LogLevelError, model.ControllersPackage, model.GetCollege, "error while fetching Colleges from database", err)
		ctx.JSON(http.StatusInternalServerError, err) // to return status code
		return
	}

	ctx.JSON(http.StatusCreated, College)

}

func (server Server) GetCollegeByID(ctx *gin.Context) {

	id := ctx.Param("id")

	utils.Log(model.LogLevelInfo, model.ControllersPackage, model.GetCollegeByID, "fetching College by id", map[string]interface{}{"id": id})

	uuid, err := uuid.Parse(id)
	if err != nil {
		utils.Log(model.LogLevelError, model.ControllersPackage, model.GetCollegeByID, "invalid College ID format", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid College ID format"})
		return
	}

	College, err := server.PostgressDB.GetCollegeByID(uuid)
	if err != nil {
		utils.Log(model.LogLevelError, model.ControllersPackage, model.GetCollegeByID, "error while inserting record ", err)
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}
	ctx.JSON(http.StatusCreated, College)
}

func (server Server) GetCollegeByFilter(ctx *gin.Context) {

	utils.Log(model.LogLevelInfo, model.ControllersPackage, model.GetCollegeByFilter, "fetching College by filter", nil)

	queryparams := ctx.Request.URL.Query()

	filter := utils.ConvertQueryParams(queryparams)

	College, err := server.PostgressDB.GetCollegeByFilter(filter)

	// uuid, err := uuid.Parse(id)
	// if err != nil {
	// 	utils.Log(model.LogLevelError, model.ControllersPackage, model.GetCollegeByID, "invalid College ID format", err)
	// 	ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid College ID format"})
	// 	return
	// }

	// College, err = server.PostgressDB.GetCollegeByID(uuid)
	if err != nil {
		utils.Log(model.LogLevelError, model.ControllersPackage, model.GetCollegeByFilter, "error while fetching record ", err)
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}
	ctx.JSON(http.StatusCreated, College)
}

func (server *Server) UpdateCollege(c *gin.Context) error {

	var College model.College
	//Unmarshal
	utils.Log(model.LogLevelInfo, model.ControllersPackage, model.UpdateCollege,
		"unmarshaling College data", nil)
	err := c.ShouldBindJSON(&College)
	if err != nil {
		utils.Log(model.LogLevelError, model.ControllersPackage, model.UpdateCollege,
			"error while unmarshaling payload", err)
		return fmt.Errorf("")
	}
	//validation is to be done here
	//DB call
	College.UpdatedAt = time.Now().UTC()
	err = server.PostgressDB.UpdateCollege(&College)
	if err != nil {
		utils.Log(model.LogLevelError, model.ControllersPackage, model.UpdateCollege,
			"error while updating record from pgress", err)
		return fmt.Errorf("")
	}
	utils.Log(model.LogLevelInfo, model.ControllersPackage, model.GetCollege,
		"successfully updated College record and setting response", College)
	c.JSON(http.StatusOK, College)
	return nil

}

func (server *Server) DeleteCollege(c *gin.Context) error {

	//validation is to be done here
	utils.Log(model.LogLevelInfo, model.ControllersPackage, model.DeleteCollege,
		"reading College id", nil)
	id := c.Param("id")
	if id == "" {
		utils.Log(model.LogLevelError, model.ControllersPackage, model.DeleteCollege,
			"missing College id", nil)
		return fmt.Errorf("")
	}
	//DB call
	err := server.PostgressDB.DeleteCollege(id)
	if err != nil {
		utils.Log(model.LogLevelError, model.ControllersPackage, model.DeleteCollege,
			"error while deleting College record from pgress", err)
		return fmt.Errorf("")
	}
	utils.Log(model.LogLevelInfo, model.ControllersPackage, model.DeleteCollege,
		"successfully deleted College record ", nil)
	return nil

}
