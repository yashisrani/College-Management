package api

import (
	"github.com/gin-gonic/gin"
	"github.com/yashisrani/Go-Backend/model"
	"github.com/yashisrani/Go-Backend/utils"
)

func (api *Apiroutes) TeacherRoutes(route *gin.Engine) {
	group := route.Group("Teacher")
	{
		group.POST("/create", api.AuthMiddleware(), api.CreateTeacher)
		group.GET("/getTeacheres", api.AuthMiddleware(), api.GetTeacher)
		group.GET("/:id", api.AuthMiddleware(), api.GetTeacherByID)
		group.GET("/filter", api.AuthMiddleware(), api.GetTeacherByFilter)
		group.PUT("/update/:id", api.AuthMiddleware(), api.UpdateTeacher)
		group.DELETE("/delete/:id", api.AuthMiddleware(), api.DeleteTeacher)
	}
}

// Handler to create a Teacher
// @router /Teacher/create [post]
// @summary Create a new Teacher
// @tags Teacher
// @accept json
// @produce json
// @param teacher body model.Teacher true "Teacher object"
// @success 201 {object} model.Teacher
// @Security ApiKeyAuth
func (api Apiroutes) CreateTeacher(ctx *gin.Context) {
	utils.Log(model.LogLevelInfo, model.ApiPackage, model.CreateTeacher, "Creating new Teacher", nil)
	api.Server.CreateTeacher(ctx)
}

// Handler to get all Teachers
// @router /Teacher/getTeachers [get]
// @summary Get all Teachers
// @tags Teacher
// @produce json
// @param page query int false "Page number (default: 1)"
// @param limit query int false "Number of results per page (default: 10)"
// @success 200 {array} model.Teacher
// @Security ApiKeyAuth
func (api Apiroutes) GetTeacher(ctx *gin.Context) {
	utils.Log(model.LogLevelInfo, model.ApiPackage, model.GetTeacher, "Fetching Teacheres", nil)
	api.Server.GetTeacher(ctx)
}

// Handler to get a Teacher by ID
// @router /Teacher/{id} [get]
// @summary Get a Teacher by ID
// @tags Teacher
// @produce json
// @param id path string true "Teacher ID"
// @success 200 {object} model.Teacher
// @Security ApiKeyAuth
func (api Apiroutes) GetTeacherByID(ctx *gin.Context) {
	utils.Log(model.LogLevelInfo, model.ApiPackage, model.GetTeacherByID, "Fetching Teacher by ID", nil)
	api.Server.GetTeacherByID(ctx)
}

// Handler to get Teachers based on filters
// @router /Teacher/filter [get]
// @summary Get Teachers based on given filters
// @tags Teacher
// @produce json
// @param first_name query string false "First Name"
// @param last_name query string false "Last Name"
// @param active query bool false "Active status"
// @param created_by query string false "Created By"
// @param updated_by query string false "Updated By"
// @param teacher_id query string false "Teacher ID"
// @param subject_code query string false "Subject Code"
// @param teacher_type query string false "Teacher Type"
// @param salary query int false "Salary"
// @param skills query string false "Skills"
// @param join_date query string false "Joining Date"
// @param page query int false "Page number (default: 1)"
// @param limit query int false "Number of results per page (default: 10)"
// @success 200 {array} model.Teacher
// @Security ApiKeyAuth
func (api Apiroutes) GetTeacherByFilter(ctx *gin.Context) {
	utils.Log(model.LogLevelInfo, model.ApiPackage, model.GetTeacherByFilter, "Fetching Teacheres by filter", nil)
	api.Server.GetTeacherByFilter(ctx)
}


// Handler to update a Teacher
// @router /Teacher/update/{id} [put]
// @summary Update a Teacher
// @tags Teacher
// @accept json
// @produce json
// @param id path string true "Teacher ID"
// @param teacher body model.Teacher true "Teacher object"
// @success 200 {object} model.Teacher
// @Security ApiKeyAuth
func (api Apiroutes) UpdateTeacher(ctx *gin.Context) {
	utils.Log(model.LogLevelInfo, model.ApiPackage, model.UpdateTeacher, "Updating Teacher", nil)
	api.Server.UpdateTeacher(ctx)
}

// Handler to delete a Teacher
// @router /Teacher/delete/{id} [delete]
// @summary Delete a Teacher
// @tags Teacher
// @param id path string true "Teacher ID"
// @Security ApiKeyAuth
func (api Apiroutes) DeleteTeacher(ctx *gin.Context) {
	utils.Log(model.LogLevelInfo, model.ApiPackage, model.DeleteTeacher, "Deleting Teacher", nil)
	api.Server.DeleteTeacher(ctx)
}
