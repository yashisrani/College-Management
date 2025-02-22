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

func (server Server) CreateCourse(ctx *gin.Context) {

	utils.Log(model.LogLevelError, model.ControllersPackage, model.CreateCourse, "Creating new Course", nil)

	Course := model.Course{}
	err := ctx.ShouldBindJSON(&Course) // to bind data in json
	if err != nil {
		utils.Log(model.LogLevelError, model.ControllersPackage, model.CreateCourse, "error while json binding", err)
		ctx.JSON(http.StatusInternalServerError, err) // to return status code
		return
	}

	Course.ID = uuid.New()        // to generate new id
	Course.CreatedAt = time.Now() // to add created time in Course

	err = server.PostgressDB.CreateCourse(&Course)
	if err != nil {
		utils.Log(model.LogLevelError, model.ControllersPackage, model.CreateCourse, "error while inserting Course to database", err)
		ctx.JSON(http.StatusInternalServerError, err) // to return status code
		return
	}

	ctx.JSON(http.StatusCreated, Course)
}

func (server Server) GetCourse(ctx *gin.Context) {

	utils.Log(model.LogLevelError, model.ControllersPackage, model.GetCourse, "Fetching Courses", nil)

	Course, err := server.PostgressDB.GetCourse()
	if err != nil {
		utils.Log(model.LogLevelError, model.ControllersPackage, model.GetCourse, "error while fetching Courses from database", err)
		ctx.JSON(http.StatusInternalServerError, err) // to return status code
		return
	}

	ctx.JSON(http.StatusCreated, Course)

}

func (server Server) GetCourseByID(ctx *gin.Context) {

	id := ctx.Param("id")

	utils.Log(model.LogLevelInfo, model.ControllersPackage, model.GetCourseByID, "fetching Course by id", map[string]interface{}{"id": id})

	uuid, err := uuid.Parse(id)
	if err != nil {
		utils.Log(model.LogLevelError, model.ControllersPackage, model.GetCourseByID, "invalid Course ID format", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Course ID format"})
		return
	}

	Course, err := server.PostgressDB.GetCourseByID(uuid)
	if err != nil {
		utils.Log(model.LogLevelError, model.ControllersPackage, model.GetCourseByID, "error while inserting record ", err)
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}
	ctx.JSON(http.StatusCreated, Course)
}

func (server Server) GetCourseByFilter(ctx *gin.Context) {

	utils.Log(model.LogLevelInfo, model.ControllersPackage, model.GetCourseByFilter, "fetching Course by filter", nil)

	queryparams := ctx.Request.URL.Query()

	filter := utils.ConvertQueryParams(queryparams)

	Course, err := server.PostgressDB.GetCourseByFilter(filter)

	// uuid, err := uuid.Parse(id)
	// if err != nil {
	// 	utils.Log(model.LogLevelError, model.ControllersPackage, model.GetCourseByID, "invalid Course ID format", err)
	// 	ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Course ID format"})
	// 	return
	// }

	// Course, err = server.PostgressDB.GetCourseByID(uuid)
	if err != nil {
		utils.Log(model.LogLevelError, model.ControllersPackage, model.GetCourseByFilter, "error while fetching record ", err)
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}
	ctx.JSON(http.StatusCreated, Course)
}

func (server *Server) UpdateCourse(c *gin.Context) error {

	var Course model.Course
	//Unmarshal
	utils.Log(model.LogLevelInfo, model.ControllersPackage, model.UpdateCourse,
		"unmarshaling Course data", nil)
	err := c.ShouldBindJSON(&Course)
	if err != nil {
		utils.Log(model.LogLevelError, model.ControllersPackage, model.UpdateCourse,
			"error while unmarshaling payload", err)
		return fmt.Errorf("")
	}
	//validation is to be done here
	//DB call
	Course.UpdatedAt = time.Now().UTC()
	err = server.PostgressDB.UpdateCourse(&Course)
	if err != nil {
		utils.Log(model.LogLevelError, model.ControllersPackage, model.UpdateCourse,
			"error while updating record from pgress", err)
		return fmt.Errorf("")
	}
	utils.Log(model.LogLevelInfo, model.ControllersPackage, model.GetCourse,
		"successfully updated Course record and setting response", Course)
	c.JSON(http.StatusOK, Course)
	return nil

}

func (server *Server) DeleteCourse(c *gin.Context) error {

	//validation is to be done here
	utils.Log(model.LogLevelInfo, model.ControllersPackage, model.DeleteCourse,
		"reading Course id", nil)
	id := c.Param("id")
	if id == "" {
		utils.Log(model.LogLevelError, model.ControllersPackage, model.DeleteCourse,
			"missing Course id", nil)
		return fmt.Errorf("")
	}
	//DB call
	err := server.PostgressDB.DeleteCourse(id)
	if err != nil {
		utils.Log(model.LogLevelError, model.ControllersPackage, model.DeleteCourse,
			"error while deleting Course record from pgress", err)
		return fmt.Errorf("")
	}
	utils.Log(model.LogLevelInfo, model.ControllersPackage, model.DeleteCourse,
		"successfully deleted Course record ", nil)
	return nil

}
