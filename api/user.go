package api

import (
	"github.com/gin-gonic/gin"
	"github.com/yashisrani/Go-Backend/model"
	"github.com/yashisrani/Go-Backend/utils"
)

func (api *Apiroutes) UserRoutes(route *gin.Engine) {
	group := route.Group("user")
	{
		group.POST("/create", api.AuthMiddleware(), api.CreateUser)
		group.GET("/getusers", api.AuthMiddleware(), api.GetUsers)
		group.GET("/:id", api.AuthMiddleware(), api.GetUserByID)
		group.GET("/filter", api.AuthMiddleware(), api.GetUserByFilter)
		group.PUT("/update/:id", api.AuthMiddleware(), api.UpdateUser)
		group.DELETE("/delete/:id", api.AuthMiddleware(), api.DeleteUser)
		group.POST("/signup", api.SignUp)
		group.POST("/signin", api.SignIn)
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

// Handler to get all users based on filter
// @router /user/filter [get]
// @summary Get all users based on given filters
// @tags users
// @produce json
// @param email query string false "email"
// @param id query string false "id"
// @param password query string false "password"
// @param active query bool false "active"
// @param created_by query string false "created_by"
// @param updated_by query string false "updated_by"
// @param first_name query string false "first_name"
// @param last_name query string false "last_name"
// @param lane query string false "lane"
// @param village query string false "village"
// @param city query string false "city"
// @param district query string false "district"
// @param pincode query int false "pincode"
// @param state query string false "state"
// @param start_date query string false "start date"
// @param end_date query string false "end date"
// @param page query int false "Page number (default: 1)"
// @param limit query int false "Number of results per page (default: 10)"
// @success 200 {array} model.User
// @Security ApiKeyAuth
func (api Apiroutes) GetUserByFilter(ctx *gin.Context) {
	utils.Log(model.LogLevelInfo, model.ApiPackage, model.GetUserByFilter, "fetching user", nil)
	api.Server.GetUserByFilter(ctx)
}

// Handler to update a user
// @router /user/update/{id} [put]
// @summary Update a user
// @tags users
// @accept json
// @produce json
// @param id path string true "User ID"
// @param user body model.User true "User object"
// @success 200 {object} model.User
// @Security ApiKeyAuth
func (api Apiroutes) UpdateUser(ctx *gin.Context) {
	utils.Log(model.LogLevelInfo, model.ApiPackage, model.UpdateUser, "updating user", nil)
	api.Server.UpdateUser(ctx)
}

// Handler to delete a user
// @router  /user/delete/{id} [delete]
// @summary Delete a user
// @tags users
// @param id path string true "User ID"
// @Security ApiKeyAuth
func (api Apiroutes) DeleteUser(ctx *gin.Context) {
	utils.Log(model.LogLevelInfo, model.ApiPackage, model.DeleteUser, "deleteing user", nil)
	api.Server.DeleteUser(ctx)
}

// Handler to signIn a user by email and password
// @router /user/signin [post]
// @summary SighIn user
// @tags users
// @produce json
// @param user body model.UserSignIn true "User object"
// @Success 200 {string} string "Successful SignIn"
// @failure 404 {object} model.ErrorResponse
// @Security ApiKeyAuth
func (api Apiroutes) SignIn(ctx *gin.Context) {
	utils.Log(model.LogLevelInfo, model.ApiPackage, model.SignIn, "signin user", nil)
	api.Server.SignIn(ctx)
}

// Handler to SignUp a user
// @router /user/signup [post]
// @summary SignUp a user
// @tags users
// @accept json
// @produce json
// @param user body model.User true "User object"
// @Success 200 {string} string "Successful SignUp"
// @failure 400 {object} model.ErrorResponse
// @Security ApiKeyAuth
func (api Apiroutes) SignUp(ctx *gin.Context) {
	utils.Log(model.LogLevelInfo, model.ApiPackage, model.SignUP, "signup user", nil)
	api.Server.SignUp(ctx)
}
