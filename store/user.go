package store

import (
	"github.com/google/uuid"
	"github.com/yashisrani/Go-Backend/model"
	"github.com/yashisrani/Go-Backend/utils"
	"gorm.io/gorm"
)

func (store Postgress) CreateUser(user *model.User) error {
	utils.Log(model.LogLevelInfo, model.Store, model.CreateUser, "creating new user", nil)
	response := store.DB.Create(user)
	if response.Error != nil {
		utils.Log(model.LogLevelError, model.Store, model.CreateUser, "error while creating user", response.Error)
	}

	utils.Log(model.LogLevelInfo, model.Store, model.CreateUser, "user created successfully", user)
	return nil
}

func (store Postgress) GetUsers() ([]model.User, error) {

	// create db layer
	users := []model.User{}

	response := store.DB.Find(&users)
	if response.Error != nil {
		utils.Log(model.LogLevelError, model.Store, model.GetUsers, "error while fetching user", response.Error)
	}

	utils.Log(model.LogLevelInfo, model.Store, model.GetUsers, "user fetched successfully", users)
	return users, nil
}

func (store Postgress) GetUserByID(userID uuid.UUID) (*model.User, error) {
	// Create a user variable
	var user model.User

	// Fetch user by ID
	response := store.DB.First(&user, userID)
	if response.Error != nil {
		if response.Error == gorm.ErrRecordNotFound {
			utils.Log(model.LogLevelError, model.Store, model.GetUserByID, "user record is not found", response.Error)
		} else {
			utils.Log(model.LogLevelError, model.Store, model.GetUserByID, "error while fetching user by ID", response.Error)
			return nil, response.Error
		}
	}

	utils.Log(model.LogLevelInfo, model.Store, model.GetUserByID, "user fetched successfully", user)
	return &user, nil
}
