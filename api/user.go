package api

import (
	"github.com/gin-gonic/gin"
	"github.com/yashisrani/Go-Backend/model"
	"github.com/yashisrani/Go-Backend/utils"
)

func (api *Apiroutes) UserRoutes(route *gin.Engine) {
	group := route.Group("user")
	{
		group.POST("/create", api.CreateUser)
		group.GET("/getusers", api.GetUsers)
		group.GET("/:id", api.GetUserByID)
	}

}

// Handler to create a user
// @router /user/create [post]
// @summary Create a user
// @tags users
// @accept json
// @produce json
// @param user body model.User true "User object"
// @success 201 {object} model.User
func (api Apiroutes) CreateUser(ctx *gin.Context) {
	utils.Log(model.LogLevelInfo, model.ApiPackage, model.CreateUser, "creating new user", nil)
	api.Server.CreateUser(ctx)
}

// Handler to get all users
// @router /user/getusers [get]
// @summary Get all users
// @tags users
// @produce json
// @param page query int false "Page number (default: 1)"
// @param limit query int false "Number of results per page (default: 10)"
// @success 200 {array} model.User
// @Security ApiKeyAuth
func (api Apiroutes) GetUsers(ctx *gin.Context) {
	utils.Log(model.LogLevelInfo, model.ApiPackage, model.GetUsers, "fetching user", nil)
	api.Server.GetUsers(ctx)
}

// Handler to get a user by ID
// @router /user/{id} [get]
// @summary Get a user by ID
// @tags users
// @produce json
// @param id path string true "User ID"
// @success 200 {object} model.User
// @Security ApiKeyAuth
func (api Apiroutes) GetUserByID(ctx *gin.Context) {
	utils.Log(model.LogLevelInfo, model.ApiPackage, model.GetUserByID, "fetching user", nil)
	api.Server.GetUserByID(ctx)
}
