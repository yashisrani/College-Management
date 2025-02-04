package controllers

import (
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
