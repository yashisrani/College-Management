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

func (server Server) CreateBook(ctx *gin.Context) {

	utils.Log(model.LogLevelError, model.ControllersPackage, model.CreateBook, "Creating new Book", nil)

	Book := model.Book{}
	err := ctx.ShouldBindJSON(&Book) // to bind data in json
	if err != nil {
		utils.Log(model.LogLevelError, model.ControllersPackage, model.CreateBook, "error while json binding", err)
		ctx.JSON(http.StatusInternalServerError, err) // to return status code
		return
	}

	Book.ID = uuid.New()        // to generate new id
	Book.CreatedAt = time.Now() // to add created time in Book

	err = server.PostgressDB.CreateBook(&Book)
	if err != nil {
		utils.Log(model.LogLevelError, model.ControllersPackage, model.CreateBook, "error while inserting Book to database", err)
		ctx.JSON(http.StatusInternalServerError, err) // to return status code
		return
	}

	ctx.JSON(http.StatusCreated, Book)
}

func (server Server) GetBook(ctx *gin.Context) {

	utils.Log(model.LogLevelError, model.ControllersPackage, model.GetBook, "Fetching Books", nil)

	Book, err := server.PostgressDB.GetBook()
	if err != nil {
		utils.Log(model.LogLevelError, model.ControllersPackage, model.GetBook, "error while fetching Books from database", err)
		ctx.JSON(http.StatusInternalServerError, err) // to return status code
		return
	}

	ctx.JSON(http.StatusCreated, Book)

}

func (server Server) GetBookByID(ctx *gin.Context) {

	id := ctx.Param("id")

	utils.Log(model.LogLevelInfo, model.ControllersPackage, model.GetBookByID, "fetching Book by id", map[string]interface{}{"id": id})

	uuid, err := uuid.Parse(id)
	if err != nil {
		utils.Log(model.LogLevelError, model.ControllersPackage, model.GetBookByID, "invalid Book ID format", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Book ID format"})
		return
	}

	Book, err := server.PostgressDB.GetBookByID(uuid)
	if err != nil {
		utils.Log(model.LogLevelError, model.ControllersPackage, model.GetBookByID, "error while inserting record ", err)
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}
	ctx.JSON(http.StatusCreated, Book)
}

func (server Server) GetBookByFilter(ctx *gin.Context) {

	utils.Log(model.LogLevelInfo, model.ControllersPackage, model.GetBookByFilter, "fetching Book by filter", nil)

	queryparams := ctx.Request.URL.Query()

	filter := utils.ConvertQueryParams(queryparams)

	Book, err := server.PostgressDB.GetBookByFilter(filter)

	// uuid, err := uuid.Parse(id)
	// if err != nil {
	// 	utils.Log(model.LogLevelError, model.ControllersPackage, model.GetBookByID, "invalid Book ID format", err)
	// 	ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Book ID format"})
	// 	return
	// }

	// Book, err = server.PostgressDB.GetBookByID(uuid)
	if err != nil {
		utils.Log(model.LogLevelError, model.ControllersPackage, model.GetBookByFilter, "error while fetching record ", err)
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}
	ctx.JSON(http.StatusCreated, Book)
}

func (server *Server) UpdateBook(c *gin.Context) error {

	var Book model.Book
	//Unmarshal
	utils.Log(model.LogLevelInfo, model.ControllersPackage, model.UpdateBook,
		"unmarshaling Book data", nil)
	err := c.ShouldBindJSON(&Book)
	if err != nil {
		utils.Log(model.LogLevelError, model.ControllersPackage, model.UpdateBook,
			"error while unmarshaling payload", err)
		return fmt.Errorf("")
	}
	//validation is to be done here
	//DB call
	Book.UpdatedAt = time.Now().UTC()
	err = server.PostgressDB.UpdateBook(&Book)
	if err != nil {
		utils.Log(model.LogLevelError, model.ControllersPackage, model.UpdateBook,
			"error while updating record from pgress", err)
		return fmt.Errorf("")
	}
	utils.Log(model.LogLevelInfo, model.ControllersPackage, model.GetBook,
		"successfully updated Book record and setting response", Book)
	c.JSON(http.StatusOK, Book)
	return nil

}

func (server *Server) DeleteBook(c *gin.Context) error {

	//validation is to be done here
	utils.Log(model.LogLevelInfo, model.ControllersPackage, model.DeleteBook,
		"reading Book id", nil)
	id := c.Param("id")
	if id == "" {
		utils.Log(model.LogLevelError, model.ControllersPackage, model.DeleteBook,
			"missing Book id", nil)
		return fmt.Errorf("")
	}
	//DB call
	err := server.PostgressDB.DeleteBook(id)
	if err != nil {
		utils.Log(model.LogLevelError, model.ControllersPackage, model.DeleteBook,
			"error while deleting Book record from pgress", err)
		return fmt.Errorf("")
	}
	utils.Log(model.LogLevelInfo, model.ControllersPackage, model.DeleteBook,
		"successfully deleted Book record ", nil)
	return nil

}
