package api

import (
	"github.com/yashisrani/Go-Backend/controllers"
	"github.com/yashisrani/Go-Backend/store"
)

type Apiroutes struct {
	Server controllers.ServerOperation
}

func (api *Apiroutes) StartApp(server controllers.Server)  {
	api.Server = &server
	api.Server.NewServer(store.Postgress{})
}