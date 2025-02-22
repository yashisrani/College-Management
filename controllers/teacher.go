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

func (server Server) CreateTeacher(ctx *gin.Context) {

	utils.Log(model.LogLevelError, model.ControllersPackage, model.CreateTeacher, "Creating new Teacher", nil)

	Teacher := model.Teacher{}
	err := ctx.ShouldBindJSON(&Teacher) // to bind data in json
	if err != nil {
		utils.Log(model.LogLevelError, model.ControllersPackage, model.CreateTeacher, "error while json binding", err)
		ctx.JSON(http.StatusInternalServerError, err) // to return status code
		return
	}

	Teacher.ID = uuid.New()        // to generate new id
	Teacher.CreatedAt = time.Now() // to add created time in Teacher

	err = server.PostgressDB.CreateTeacher(&Teacher)
	if err != nil {
		utils.Log(model.LogLevelError, model.ControllersPackage, model.CreateTeacher, "error while inserting Teacher to database", err)
		ctx.JSON(http.StatusInternalServerError, err) // to return status code
		return
	}

	ctx.JSON(http.StatusCreated, Teacher)
}

func (server Server) GetTeacher(ctx *gin.Context) {

	utils.Log(model.LogLevelError, model.ControllersPackage, model.GetTeacher, "Fetching Teachers", nil)

	Teacher, err := server.PostgressDB.GetTeacher()
	if err != nil {
		utils.Log(model.LogLevelError, model.ControllersPackage, model.GetTeacher, "error while fetching Teachers from database", err)
		ctx.JSON(http.StatusInternalServerError, err) // to return status code
		return
	}

	ctx.JSON(http.StatusCreated, Teacher)

}

func (server Server) GetTeacherByID(ctx *gin.Context) {

	id := ctx.Param("id")

	utils.Log(model.LogLevelInfo, model.ControllersPackage, model.GetTeacherByID, "fetching Teacher by id", map[string]interface{}{"id": id})

	uuid, err := uuid.Parse(id)
	if err != nil {
		utils.Log(model.LogLevelError, model.ControllersPackage, model.GetTeacherByID, "invalid Teacher ID format", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Teacher ID format"})
		return
	}

	Teacher, err := server.PostgressDB.GetTeacherByID(uuid)
	if err != nil {
		utils.Log(model.LogLevelError, model.ControllersPackage, model.GetTeacherByID, "error while inserting record ", err)
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}
	ctx.JSON(http.StatusCreated, Teacher)
}

func (server Server) GetTeacherByFilter(ctx *gin.Context) {

	utils.Log(model.LogLevelInfo, model.ControllersPackage, model.GetTeacherByFilter, "fetching Teacher by filter", nil)

	queryparams := ctx.Request.URL.Query()

	filter := utils.ConvertQueryParams(queryparams)

	Teacher, err := server.PostgressDB.GetTeacherByFilter(filter)

	// uuid, err := uuid.Parse(id)
	// if err != nil {
	// 	utils.Log(model.LogLevelError, model.ControllersPackage, model.GetTeacherByID, "invalid Teacher ID format", err)
	// 	ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Teacher ID format"})
	// 	return
	// }

	// Teacher, err = server.PostgressDB.GetTeacherByID(uuid)
	if err != nil {
		utils.Log(model.LogLevelError, model.ControllersPackage, model.GetTeacherByFilter, "error while fetching record ", err)
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}
	ctx.JSON(http.StatusCreated, Teacher)
}

func (server *Server) UpdateTeacher(c *gin.Context) error {

	var Teacher model.Teacher
	//Unmarshal
	utils.Log(model.LogLevelInfo, model.ControllersPackage, model.UpdateTeacher,
		"unmarshaling Teacher data", nil)
	err := c.ShouldBindJSON(&Teacher)
	if err != nil {
		utils.Log(model.LogLevelError, model.ControllersPackage, model.UpdateTeacher,
			"error while unmarshaling payload", err)
		return fmt.Errorf("")
	}
	//validation is to be done here
	//DB call
	Teacher.UpdatedAt = time.Now().UTC()
	err = server.PostgressDB.UpdateTeacher(&Teacher)
	if err != nil {
		utils.Log(model.LogLevelError, model.ControllersPackage, model.UpdateTeacher,
			"error while updating record from pgress", err)
		return fmt.Errorf("")
	}
	utils.Log(model.LogLevelInfo, model.ControllersPackage, model.GetTeacher,
		"successfully updated Teacher record and setting response", Teacher)
	c.JSON(http.StatusOK, Teacher)
	return nil

}

func (server *Server) DeleteTeacher(c *gin.Context) error {

	//validation is to be done here
	utils.Log(model.LogLevelInfo, model.ControllersPackage, model.DeleteTeacher,
		"reading Teacher id", nil)
	id := c.Param("id")
	if id == "" {
		utils.Log(model.LogLevelError, model.ControllersPackage, model.DeleteTeacher,
			"missing Teacher id", nil)
		return fmt.Errorf("")
	}
	//DB call
	err := server.PostgressDB.DeleteTeacher(id)
	if err != nil {
		utils.Log(model.LogLevelError, model.ControllersPackage, model.DeleteTeacher,
			"error while deleting Teacher record from pgress", err)
		return fmt.Errorf("")
	}
	utils.Log(model.LogLevelInfo, model.ControllersPackage, model.DeleteTeacher,
		"successfully deleted Teacher record ", nil)
	return nil

}
