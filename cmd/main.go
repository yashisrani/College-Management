package main

import (
	"github.com/gin-gonic/gin"
	"github.com/yashisrani/Go-Backend/api"
	"github.com/yashisrani/Go-Backend/controllers"
)

func main() {
	api := api.Apiroutes{}
	controller := controllers.Server{}
	route := gin.Default()
	api.StartApp(route, controller)
	
	route.Run(":8000")
	// fmt.Println("main", api)
}
