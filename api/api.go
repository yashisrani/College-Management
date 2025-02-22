package api

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/yashisrani/Go-Backend/controllers"
	"github.com/yashisrani/Go-Backend/store"
	_ "github.com/yashisrani/Go-Backend/docs" // which is the generated folder after swag init
)

type Apiroutes struct {
	Server controllers.ServerOperation
}

func (api *Apiroutes) StartApp(router *gin.Engine, server controllers.Server) {
	api.Server = &server
	api.Server.NewServer(store.Postgress{})

	// Swagger documentation
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	api.UserRoutes(router)
	api.CollegeRoutes(router)
	api.ClassRoutes(router)
	api.TeacherRoutes(router)
	api.BookRoutes(router)
	api.ComputerLabRoutes(router)
	api.CourseRoutes(router)
}
