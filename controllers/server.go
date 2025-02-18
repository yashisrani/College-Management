package controllers

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/yashisrani/Go-Backend/model"
	"github.com/yashisrani/Go-Backend/store"
	"github.com/yashisrani/Go-Backend/utils"
)

type Server struct {
	PostgressDB store.StoreOperation
}

func (s *Server) NewServer(pgstore store.Postgress) {
	utils.SetLogger()
	utils.Log(model.LogLevelInfo, model.Controllers, "NewServer", "logger setup ", nil)
	s.PostgressDB = &pgstore
	err := s.PostgressDB.NewStore()
	if err != nil {
		utils.Log(model.LogLevelError, model.Controllers, "NewServer", "error creating store: %v", err)
	} else {
		utils.Log(model.LogLevelInfo, model.Controllers, "NewServer", "connected with store controller", nil)
	}
	fmt.Println("controller", s)
}

type ServerOperation interface {
	NewServer(pgstore store.Postgress)
	CreateUser(ctx *gin.Context)
	GetUsers(ctx *gin.Context)
	GetUserByID(ctx *gin.Context)
	GetUserByFilter(ctx *gin.Context)
	DeleteUser(c *gin.Context) error
	UpdateUser(c *gin.Context) error
	SignUp(c *gin.Context)
	SignIn(c *gin.Context)
}
