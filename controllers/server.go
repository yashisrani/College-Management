package controllers

import (
	"fmt"

	"github.com/yashisrani/Go-Backend/store"
)

type Server struct {
	PostgressDB store.StoreOperation
}

func (s *Server) NewServer(pgstore store.Postgress)  {
	s.PostgressDB = &pgstore
	s.PostgressDB.NewStore()
	fmt.Println("controller",s)
}

type ServerOperation interface {
	NewServer(pgstore store.Postgress)
}