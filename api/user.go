package api

import (
	"github.com/gin-gonic/gin"
	"github.com/yashisrani/Go-Backend/model"
	"github.com/yashisrani/Go-Backend/utils"
)

func (api *Apiroutes) UserRoutes(route *gin.Engine) {
	group := route.Group("user")
	{
		group.POST("/createuser", api.CreateUser)
	}

}

func (api Apiroutes) CreateUser(ctx *gin.Context) {
	utils.Log(model.LogLevelInfo, model.ApiPackage, model.CreateUser, "creating new user", nil)
}
