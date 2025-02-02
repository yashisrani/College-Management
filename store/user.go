package store

import (
	"github.com/yashisrani/Go-Backend/model"
	"github.com/yashisrani/Go-Backend/utils"
)

func (store Postgress) CreateUser(user *model.User) error {
	utils.Log(model.LogLevelInfo, model.Store, model.CreateUser, "creating new user", nil)
	response := store.DB.Create(user)
	if response.Error != nil {
		utils.Log(model.LogLevelError, model.Store, model.CreateUser, "error while creating user", response.Error)
	}
	
	utils.Log(model.LogLevelInfo, model.Store, model.CreateUser, "user created successfully", nil)
	return nil
}
