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
	user := model.User{}
	err := ctx.Bind(&user) // to bind data in json
	if err != nil {
		utils.Log(model.LogLevelError, model.Controllers, model.CreateUser, "error while json binding", err)
		ctx.JSON(http.StatusInternalServerError, err) // to return status code
		return
	}

	user.ID = uuid.New() // to generate new id
	user.CreatedAt = time.Now() // to add created time in user 

	err = server.PostgressDB.CreateUser(&user)
	if err!= nil {
        utils.Log(model.LogLevelError, model.Controllers, model.CreateUser, "error while inserting user to database", err)
        ctx.JSON(http.StatusInternalServerError, err) // to return status code
        return
    }

	ctx.JSON(http.StatusCreated, err)

}
