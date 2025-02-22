package store

import (
	"fmt"

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

func (store Postgress) UpdateUser(user *model.User) error {

	utils.Log(model.LogLevelInfo, model.StorePackage, model.UpdateUser, "updating user data", *user)
	resp := store.DB.Save(user)
	if resp.Error != nil {
		utils.Log(model.LogLevelError, model.StorePackage, model.UpdateUser,
			"error while updating user data", resp.Error)
		return fmt.Errorf("error while updating user record, err = %v", resp.Error)
	}
	utils.Log(model.LogLevelInfo, model.StorePackage, model.UpdateUser,
		"successfully updated user", nil)
	return nil
}

func (store Postgress) DeleteUser(userID string) error {

	var user model.User

	utils.Log(model.LogLevelInfo, model.StorePackage, model.DeleteUser, "deleting user data", userID)
	if err := store.DB.First(&user, `"id" = '`+userID+`'`).Error; err != nil {
		utils.Log(model.LogLevelError, model.StorePackage, model.DeleteUser,
			"error while deleting user data", err)
		return fmt.Errorf("user not found for given id, ID = %v", userID)
	}
	resp := store.DB.Delete(user)
	if resp.Error != nil {
		return fmt.Errorf("error while deleting user record from DB, err = %v", resp.Error)
	}
	utils.Log(model.LogLevelInfo, model.StorePackage, model.DeleteUser,
		"successfully deleted user", nil)
	return nil
}

// soft delete: not actually deleting record, but change the status of active (eg. active status is true then make it flase)
// hard delete: actually deleting record.

// func (store Postgress) SoftDeleteUser(userID string) error {
// 	var user model.User

// 	utils.Log(model.LogLevelInfo, model.StorePackage, model.DeleteUser, "soft deleting user data", userID)

// 	// Find the user by ID
// 	if err := store.DB.First(&user, `"id" = '`+userID+`'`).Error; err != nil {
// 		utils.Log(model.LogLevelError, model.StorePackage, model.DeleteUser,
// 			"error while finding user data", err)
// 		return fmt.Errorf("user not found for given ID, ID = %v", userID)
// 	}

// 	// Set active field to false (Soft delete)
// 	user.Active = false
// 	if err := store.DB.Save(&user).Error; err != nil {
// 		utils.Log(model.LogLevelError, model.StorePackage, model.DeleteUser,
// 			"error while updating user active status", err)
// 		return fmt.Errorf("error while updating user record in DB, err = %v", err)
// 	}

// 	utils.Log(model.LogLevelInfo, model.StorePackage, model.DeleteUser,
// 		"successfully soft deleted user", userID)
// 	return nil
// }


func (store Postgress) GetUserByFilter(filter map[string]interface{}) ([]model.User, error) {

	utils.Log(model.LogLevelInfo, model.StorePackage, model.GetUserByFilter, "fetching records of user from db", nil)
	var users []model.User
	query := store.DB    // 

	for key, value := range filter {
		if key == model.PageNumber || key == model.DataPerPage || key == model.StartDate || key == model.EndDate {
			continue
		}
		utils.Log(model.LogLevelInfo, model.StorePackage, model.GetUserByFilter,
			"filters key", key+" value = "+fmt.Sprintf("%v", value))
		query = query.Where(fmt.Sprintf("%s = ?", key), value)
	}
	SetLimitAndPage(filter, query)
	SetDateRangeFilter(filter, query)

	err := query.Find(&users).Error
	if err != nil {
		utils.Log(model.LogLevelError, model.StorePackage, model.GetUserByFilter,
			"error while reading user data", err)
		return nil, fmt.Errorf("error while fetching user records from DB, err = %v", err)
	}

	utils.Log(model.LogLevelInfo, model.StorePackage, model.GetUserByFilter, "records of users from db", users)
	return users, nil
}


func (store Postgress) SignUp(user *model.User) error {

	utils.Log(model.LogLevelInfo, model.StorePackage, model.SignUP, "creating user data with signup api", *user)
	resp := store.DB.Create(user)
	if resp.Error != nil {
		utils.Log(model.LogLevelError, model.StorePackage, model.SignUP,
			"error while creating user data", resp.Error)
		return fmt.Errorf("error while creating user record with signup api, err = %v", resp.Error)
	}
	utils.Log(model.LogLevelInfo, model.StorePackage, model.SignUP,
		"successfully created user with signup api", nil)
	return nil
}

func (store Postgress) SignIn(userSignIn model.UserSignIn) (*model.User, error) {
	var user model.User
	utils.Log(model.LogLevelInfo, model.StorePackage, model.SignIn,
		"reading user data from db based on email", userSignIn)
	resp := store.DB.Where("email = ? AND password = ?", userSignIn.EmailID, userSignIn.Password).First(&user)
	if resp.Error != nil {
		utils.Log(model.LogLevelError, model.StorePackage, model.SignIn,
			"error while reading user data", resp.Error)
		return &user, fmt.Errorf("error while fetching user record from DB for given id, err = %v", resp.Error)
	}
	utils.Log(model.LogLevelInfo, model.StorePackage, model.SignIn,
		"returning user data", user)
	return &user, nil
}

