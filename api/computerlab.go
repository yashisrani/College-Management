package api

import (
	"github.com/gin-gonic/gin"
	"github.com/yashisrani/Go-Backend/model"
	"github.com/yashisrani/Go-Backend/utils"
)

func (api *Apiroutes) ComputerLabRoutes(route *gin.Engine) {
	group := route.Group("ComputerLab")
	{
		group.POST("/create", api.AuthMiddleware(), api.CreateComputerLab)
		group.GET("/getComputerLabes", api.AuthMiddleware(), api.GetComputerLab)
		group.GET("/:id", api.AuthMiddleware(), api.GetComputerLabByID)
		group.GET("/filter", api.AuthMiddleware(), api.GetComputerLabByFilter)
		group.PUT("/update/:id", api.AuthMiddleware(), api.UpdateComputerLab)
		group.DELETE("/delete/:id", api.AuthMiddleware(), api.DeleteComputerLab)
	}
}

// Handler to create a ComputerLab
// @router /ComputerLab/create [post]
// @summary Create a new ComputerLab
// @tags ComputerLab
// @accept json
// @produce json
// @param ComputerLab body model.ComputerLab true "ComputerLab object"
// @success 201 {object} model.ComputerLab
// @Security ApiKeyAuth
func (api Apiroutes) CreateComputerLab(ctx *gin.Context) {
	utils.Log(model.LogLevelInfo, model.ApiPackage, model.CreateComputerLab, "Creating new ComputerLab", nil)
	api.Server.CreateComputerLab(ctx)
}

// Handler to get all ComputerLabs
// @router /ComputerLab/getComputerLabs [get]
// @summary Get all ComputerLabs
// @tags ComputerLab
// @produce json
// @param page query int false "Page number (default: 1)"
// @param limit query int false "Number of results per page (default: 10)"
// @success 200 {array} model.ComputerLab
// @Security ApiKeyAuth
func (api Apiroutes) GetComputerLab(ctx *gin.Context) {
	utils.Log(model.LogLevelInfo, model.ApiPackage, model.GetComputerLab, "Fetching ComputerLabes", nil)
	api.Server.GetComputerLab(ctx)
}

// Handler to get a ComputerLab by ID
// @router /ComputerLab/{id} [get]
// @summary Get a ComputerLab by ID
// @tags ComputerLab
// @produce json
// @param id path string true "ComputerLab ID"
// @success 200 {object} model.ComputerLab
// @Security ApiKeyAuth
func (api Apiroutes) GetComputerLabByID(ctx *gin.Context) {
	utils.Log(model.LogLevelInfo, model.ApiPackage, model.GetComputerLabByID, "Fetching ComputerLab by ID", nil)
	api.Server.GetComputerLabByID(ctx)
}

// Handler to get ComputerLabs based on filters
// @router /ComputerLab/filter [get]
// @summary Get ComputerLabs based on given filters
// @tags ComputerLab
// @produce json
// @param class_name query string false "Class Name"
// @param school_id query string false "School ID"
// @param active query bool false "Active status"
// @param instructor query string false "Instructor"
// @param total_students query int false "Total Students"
// @param classroom_number query string false "Classroom Number"
// @param syllabus query string false "Syllabus"
// @param created_by query string false "Created By"
// @param updated_by query string false "Updated By"
// @param computerlab_id query string false "ComputerLab ID"
// @param page query int false "Page number (default: 1)"
// @param limit query int false "Number of results per page (default: 10)"
// @success 200 {array} model.ComputerLab
// @Security ApiKeyAuth
func (api Apiroutes) GetComputerLabByFilter(ctx *gin.Context) {
	utils.Log(model.LogLevelInfo, model.ApiPackage, model.GetComputerLabByFilter, "Fetching ComputerLabs by filter", nil)
	api.Server.GetComputerLabByFilter(ctx)
}


// Handler to update a ComputerLab
// @router /ComputerLab/update/{id} [put]
// @summary Update a ComputerLab
// @tags ComputerLab
// @accept json
// @produce json
// @param id path string true "ComputerLab ID"
// @param ComputerLab body model.ComputerLab true "ComputerLab object"
// @success 200 {object} model.ComputerLab
// @Security ApiKeyAuth
func (api Apiroutes) UpdateComputerLab(ctx *gin.Context) {
	utils.Log(model.LogLevelInfo, model.ApiPackage, model.UpdateComputerLab, "Updating ComputerLab", nil)
	api.Server.UpdateComputerLab(ctx)
}

// Handler to delete a ComputerLab
// @router /ComputerLab/delete/{id} [delete]
// @summary Delete a ComputerLab
// @tags ComputerLab
// @param id path string true "ComputerLab ID"
// @Security ApiKeyAuth
func (api Apiroutes) DeleteComputerLab(ctx *gin.Context) {
	utils.Log(model.LogLevelInfo, model.ApiPackage, model.DeleteComputerLab, "Deleting ComputerLab", nil)
	api.Server.DeleteComputerLab(ctx)
}
