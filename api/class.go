package api

import (
	"github.com/gin-gonic/gin"
	"github.com/yashisrani/Go-Backend/model"
	"github.com/yashisrani/Go-Backend/utils"
)

func (api *Apiroutes) ClassRoutes(route *gin.Engine) {
	group := route.Group("Class")
	{
		group.POST("/create", api.AuthMiddleware(), api.CreateClass)
		group.GET("/getClasses", api.AuthMiddleware(), api.GetClass)
		group.GET("/:id", api.AuthMiddleware(), api.GetClassByID)
		group.GET("/filter", api.AuthMiddleware(), api.GetClassByFilter)
		group.PUT("/update/:id", api.AuthMiddleware(), api.UpdateClass)
		group.DELETE("/delete/:id", api.AuthMiddleware(), api.DeleteClass)
	}
}

// Handler to create a Class
// @router /Class/create [post]
// @summary Create a new Class
// @tags Class
// @accept json
// @produce json
// @param class body model.Class true "Class object"
// @success 201 {object} model.Class
// @Security ApiKeyAuth
func (api Apiroutes) CreateClass(ctx *gin.Context) {
	utils.Log(model.LogLevelInfo, model.ApiPackage, model.CreateClass, "Creating new Class", nil)
	api.Server.CreateClass(ctx)
}

// Handler to get all Classes
// @router /Class/getClasses [get]
// @summary Get all Classes
// @tags Class
// @produce json
// @param page query int false "Page number (default: 1)"
// @param limit query int false "Number of results per page (default: 10)"
// @success 200 {array} model.Class
// @Security ApiKeyAuth
func (api Apiroutes) GetClass(ctx *gin.Context) {
	utils.Log(model.LogLevelInfo, model.ApiPackage, model.GetClass, "Fetching Classes", nil)
	api.Server.GetClass(ctx)
}

// Handler to get a Class by ID
// @router /Class/{id} [get]
// @summary Get a Class by ID
// @tags Class
// @produce json
// @param id path string true "Class ID"
// @success 200 {object} model.Class
// @Security ApiKeyAuth
func (api Apiroutes) GetClassByID(ctx *gin.Context) {
	utils.Log(model.LogLevelInfo, model.ApiPackage, model.GetClassByID, "Fetching Class by ID", nil)
	api.Server.GetClassByID(ctx)
}

// Handler to get Classes based on filters
// @router /Class/filter [get]
// @summary Get Classes based on given filters
// @tags Class
// @produce json
// @param class_name query string false "Class Name"
// @param active query bool false "Active status"
// @param created_by query string false "Created By"
// @param updated_by query string false "Updated By"
// @param teacher_id query string false "Teacher ID"
// @param class_monitor query string false "Class Monitor"
// @param floor_number query int false "Floor Number"
// @param page query int false "Page number (default: 1)"
// @param limit query int false "Number of results per page (default: 10)"
// @success 200 {array} model.Class
// @Security ApiKeyAuth
func (api Apiroutes) GetClassByFilter(ctx *gin.Context) {
	utils.Log(model.LogLevelInfo, model.ApiPackage, model.GetClassByFilter, "Fetching Classes by filter", nil)
	api.Server.GetClassByFilter(ctx)
}

// Handler to update a Class
// @router /Class/update/{id} [put]
// @summary Update a Class
// @tags Class
// @accept json
// @produce json
// @param id path string true "Class ID"
// @param class body model.Class true "Class object"
// @success 200 {object} model.Class
// @Security ApiKeyAuth
func (api Apiroutes) UpdateClass(ctx *gin.Context) {
	utils.Log(model.LogLevelInfo, model.ApiPackage, model.UpdateClass, "Updating Class", nil)
	api.Server.UpdateClass(ctx)
}

// Handler to delete a Class
// @router /Class/delete/{id} [delete]
// @summary Delete a Class
// @tags Class
// @param id path string true "Class ID"
// @Security ApiKeyAuth
func (api Apiroutes) DeleteClass(ctx *gin.Context) {
	utils.Log(model.LogLevelInfo, model.ApiPackage, model.DeleteClass, "Deleting Class", nil)
	api.Server.DeleteClass(ctx)
}
