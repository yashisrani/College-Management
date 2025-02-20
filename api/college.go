package api

import (
	"github.com/gin-gonic/gin"
	"github.com/yashisrani/Go-Backend/model"
	"github.com/yashisrani/Go-Backend/utils"
)

func (api *Apiroutes) CollegeRoutes(route *gin.Engine) {
	group := route.Group("College")
	{
		group.POST("/create", api.AuthMiddleware(), api.CreateCollege)
		group.GET("/getColleges", api.AuthMiddleware(), api.GetColleges)
		group.GET("/:id", api.AuthMiddleware(), api.GetCollegeByID)
		group.GET("/filter", api.AuthMiddleware(), api.GetCollegeByFilter)
		group.PUT("/update/:id", api.AuthMiddleware(), api.UpdateCollege)
		group.DELETE("/delete/:id", api.AuthMiddleware(), api.DeleteCollege)
	}

}

// Handler to create a College
// @router /college/create [post]
// @summary Create a College
// @tags Colleges
// @accept json
// @produce json
// @param College body model.College true "College object"
// @success 201 {object} model.College
// @Security ApiKeyAuth
func (api Apiroutes) CreateCollege(ctx *gin.Context) {
	utils.Log(model.LogLevelInfo, model.ApiPackage, model.CreateCollege, "creating new College", nil)
	api.Server.CreateCollege(ctx)
}

// Handler to get all Colleges
// @router /college/getColleges [get]
// @summary Get all Colleges
// @tags Colleges
// @produce json
// @param page query int false "Page number (default: 1)"
// @param limit query int false "Number of results per page (default: 10)"
// @success 200 {array} model.College
// @Security ApiKeyAuth
func (api Apiroutes) GetColleges(ctx *gin.Context) {
	utils.Log(model.LogLevelInfo, model.ApiPackage, model.GetCollege, "fetching College", nil)
	api.Server.GetColleges(ctx)
}

// Handler to get a College by ID
// @router /college/{id} [get]
// @summary Get a College by ID
// @tags Colleges
// @produce json
// @param id path string true "College ID"
// @success 200 {object} model.College
// @Security ApiKeyAuth
func (api Apiroutes) GetCollegeByID(ctx *gin.Context) {
	utils.Log(model.LogLevelInfo, model.ApiPackage, model.GetCollegeByID, "fetching College", nil)
	api.Server.GetCollegeByID(ctx)
}

// Handler to get all Colleges based on filter
// @router /college/filter [get]
// @summary Get all Colleges based on given filters
// @tags Colleges
// @produce json
// @param domain query string false "domain"
// @param id query string false "id"
// @param hostel query string false "hostel"
// @param active query bool false "active"
// @param created_by query string false "created_by"
// @param updated_by query string false "updated_by"
// @param admin_id query string false "admin_id"
// @param director_id query string false "director_id"
// @param principle_id query string false "principle_id"
// @param lane query string false "lane"
// @param village query string false "village"
// @param university_type query string false "university_type"
// @param city query string false "city"
// @param district query string false "district"
// @param pincode query int false "pincode"
// @param state query string false "state"
// @param start_date query string false "start date"
// @param end_date query string false "end date"
// @param page query int false "Page number (default: 1)"
// @param limit query int false "Number of results per page (default: 10)"
// @success 200 {array} model.College
// @Security ApiKeyAuth
func (api Apiroutes) GetCollegeByFilter(ctx *gin.Context) {
	utils.Log(model.LogLevelInfo, model.ApiPackage, model.GetCollegeByFilter, "fetching College", nil)
	api.Server.GetCollegeByFilter(ctx)
}

// Handler to update a College
// @router /college/update/{id} [put]
// @summary Update a College
// @tags Colleges
// @accept json
// @produce json
// @param id path string true "College ID"
// @param College body model.College true "College object"
// @success 200 {object} model.College
// @Security ApiKeyAuth
func (api Apiroutes) UpdateCollege(ctx *gin.Context) {
	utils.Log(model.LogLevelInfo, model.ApiPackage, model.UpdateCollege, "updating College", nil)
	api.Server.UpdateCollege(ctx)
}

// Handler to delete a College
// @router  /college/delete/{id} [delete]
// @summary Delete a College
// @tags Colleges
// @param id path string true "College ID"
// @Security ApiKeyAuth
func (api Apiroutes) DeleteCollege(ctx *gin.Context) {
	utils.Log(model.LogLevelInfo, model.ApiPackage, model.DeleteCollege, "deleteing College", nil)
	api.Server.DeleteCollege(ctx)
}
