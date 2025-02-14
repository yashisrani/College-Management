package controllers

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/yashisrani/Go-Backend/model"
	"github.com/yashisrani/Go-Backend/utils"
)

func (server Server) CreateUser(ctx *gin.Context) {

	utils.Log(model.LogLevelError, model.ControllersPackage, model.CreateUser, "Creating new user", nil)

	user := model.User{}
	err := ctx.Bind(&user) // to bind data in json
	if err != nil {
		utils.Log(model.LogLevelError, model.ControllersPackage, model.CreateUser, "error while json binding", err)
		ctx.JSON(http.StatusInternalServerError, err) // to return status code
		return
	}

	user.ID = uuid.New()        // to generate new id
	user.CreatedAt = time.Now() // to add created time in user

	err = server.PostgressDB.CreateUser(&user)
	if err != nil {
		utils.Log(model.LogLevelError, model.ControllersPackage, model.CreateUser, "error while inserting user to database", err)
		ctx.JSON(http.StatusInternalServerError, err) // to return status code
		return
	}

	ctx.JSON(http.StatusCreated, user)
}

func (server Server) GetUsers(ctx *gin.Context) {

	utils.Log(model.LogLevelError, model.ControllersPackage, model.GetUsers, "Fetching users", nil)

	user, err := server.PostgressDB.GetUsers()
	if err != nil {
		utils.Log(model.LogLevelError, model.ControllersPackage, model.GetUsers, "error while fetching users from database", err)
		ctx.JSON(http.StatusInternalServerError, err) // to return status code
		return
	}

	ctx.JSON(http.StatusCreated, user)

}

func (server Server) GetUserByID(ctx *gin.Context) {

	id := ctx.Param("id")

	utils.Log(model.LogLevelInfo, model.ControllersPackage, model.GetUserByID, "fetching user by id", map[string]interface{}{"id": id})

	uuid, err := uuid.Parse(id)
	if err != nil {
		utils.Log(model.LogLevelError, model.ControllersPackage, model.GetUserByID, "invalid user ID format", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID format"})
		return
	}

	user, err := server.PostgressDB.GetUserByID(uuid)
	if err != nil {
		utils.Log(model.LogLevelError, model.ControllersPackage, model.GetUserByID, "error while inserting record ", err)
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}
	ctx.JSON(http.StatusCreated, user)
}

func (server Server) GetUserByFilter(ctx *gin.Context) {

	utils.Log(model.LogLevelInfo, model.ControllersPackage, model.GetUserByFilter, "fetching user by filter", nil)

	queryparams := ctx.Request.URL.Query()

	filter := utils.ConvertQueryParams(queryparams)

	user, err := server.PostgressDB.GetUserByFilter(filter)

	// uuid, err := uuid.Parse(id)
	// if err != nil {
	// 	utils.Log(model.LogLevelError, model.ControllersPackage, model.GetUserByID, "invalid user ID format", err)
	// 	ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID format"})
	// 	return
	// }

	// user, err = server.PostgressDB.GetUserByID(uuid)
	if err != nil {
		utils.Log(model.LogLevelError, model.ControllersPackage, model.GetUserByFilter, "error while fetching record ", err)
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}
	ctx.JSON(http.StatusCreated, user)
}

func (server *Server) UpdateUser(c *gin.Context) error {

	var user model.User
	//Unmarshal
	utils.Log(model.LogLevelInfo, model.ControllersPackage, model.UpdateUser,
		"unmarshaling user data", nil)
	err := c.ShouldBindJSON(&user)
	if err != nil {
		utils.Log(model.LogLevelError, model.ControllersPackage, model.UpdateUser,
			"error while unmarshaling payload", err)
		return fmt.Errorf("")
	}
	//validation is to be done here
	//DB call
	user.UpdatedAt = time.Now().UTC()
	err = server.PostgressDB.UpdateUser(&user)
	if err != nil {
		utils.Log(model.LogLevelError, model.ControllersPackage, model.UpdateUser,
			"error while updating record from pgress", err)
		return fmt.Errorf("")
	}
	utils.Log(model.LogLevelInfo, model.ControllersPackage, model.GetUsers,
		"successfully updated user record and setting response", user)
	c.JSON(http.StatusOK, user)
	return nil

}

func (server *Server) DeleteUser(c *gin.Context) error {

	//validation is to be done here
	utils.Log(model.LogLevelInfo, model.ControllersPackage, model.DeleteUser,
		"reading user id", nil)
	id := c.Param("id")
	if id == "" {
		utils.Log(model.LogLevelError, model.ControllersPackage, model.DeleteUser,
			"missing user id", nil)
		return fmt.Errorf("")
	}
	//DB call
	err := server.PostgressDB.DeleteUser(id)
	if err != nil {
		utils.Log(model.LogLevelError, model.ControllersPackage, model.DeleteUser,
			"error while deleting user record from pgress", err)
		return fmt.Errorf("")
	}
	utils.Log(model.LogLevelInfo, model.ControllersPackage, model.DeleteUser,
		"successfully deleted user record ", nil)
	return nil

}
