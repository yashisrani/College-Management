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

func (api Apiroutes) CreateUser(ctx *gin.Context) {
	utils.Log(model.LogLevelInfo, model.ApiPackage, model.CreateUser, "creating new user", nil)
	api.Server.CreateUser(ctx)
}

func (api Apiroutes) GetUsers(ctx *gin.Context) {
	utils.Log(model.LogLevelInfo, model.ApiPackage, model.GetUsers, "fetching user", nil)
	api.Server.GetUsers(ctx)
}

func (api Apiroutes) GetUserByID(ctx *gin.Context) {
	utils.Log(model.LogLevelInfo, model.ApiPackage, model.GetUserByID, "fetching user", nil)
	api.Server.GetUserByID(ctx)
}
