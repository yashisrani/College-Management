package main

import (
	"github.com/gin-gonic/gin"
	"github.com/yashisrani/Go-Backend/api"
	"github.com/yashisrani/Go-Backend/controllers"
	"github.com/yashisrani/Go-Backend/middleware"
)

// @title College-Management-System
// @version 1.0
// @description APIs to manage College operations
// @host localhost:8000
// @BasePath /
// @schemes http https
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name X-Token
func main() {
	api := api.Apiroutes{}
	controller := controllers.Server{}
	route := gin.Default()
	api.StartApp(route, controller)

	middleware.PrometheusInit()
	route.Use(middleware.TrackMetrics())

	// Expose the metrics endpoint
	route.GET("/metrics", middleware.PrometheusHandler())

	route.Run(":8000")
	// fmt.Println("main", api)
}
