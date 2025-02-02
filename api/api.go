package api

import (
	"github.com/gin-gonic/gin"
	"github.com/yashisrani/Go-Backend/controllers"
	"github.com/yashisrani/Go-Backend/store"
)

type Apiroutes struct {
	Server controllers.ServerOperation
}

func (api *Apiroutes) StartApp(route *gin.Engine, server controllers.Server) {
	api.Server = &server
	api.Server.NewServer(store.Postgress{})


	api.UserRoutes(route)
}
