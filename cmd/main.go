package main

import (
	"fmt"

	"github.com/yashisrani/Go-Backend/controllers"
	"github.com/yashisrani/Go-Backend/store"
)

func main()  {
	server:= controllers.Server{}
	server.NewServer(store.Postgress{})
	fmt.Println("main",server)
}