package main

import (
	"fmt"

	"github.com/yashisrani/Go-Backend/api"
	"github.com/yashisrani/Go-Backend/controllers"
)

func main() {
	api := api.Apiroutes{}

	api.StartApp(controllers.Server{})
	fmt.Println("main", api)
}
