package api

import (
	"github.com/gin-gonic/gin"
	"github.com/yashisrani/Go-Backend/model"
	"github.com/yashisrani/Go-Backend/utils"
)

func (api *Apiroutes) BookRoutes(route *gin.Engine) {
	group := route.Group("Book")
	{
		group.POST("/create", api.AuthMiddleware(), api.CreateBook)
		group.GET("/getBookes", api.AuthMiddleware(), api.GetBook)
		group.GET("/:id", api.AuthMiddleware(), api.GetBookByID)
		group.GET("/filter", api.AuthMiddleware(), api.GetBookByFilter)
		group.PUT("/update/:id", api.AuthMiddleware(), api.UpdateBook)
		group.DELETE("/delete/:id", api.AuthMiddleware(), api.DeleteBook)
	}
}

// Handler to create a Book
// @router /Book/create [post]
// @summary Create a new Book
// @tags Book
// @accept json
// @produce json
// @param Book body model.Book true "Book object"
// @success 201 {object} model.Book
// @Security ApiKeyAuth
func (api Apiroutes) CreateBook(ctx *gin.Context) {
	utils.Log(model.LogLevelInfo, model.ApiPackage, model.CreateBook, "Creating new Book", nil)
	api.Server.CreateBook(ctx)
}

// Handler to get all Books
// @router /Book/getBooks [get]
// @summary Get all Books
// @tags Book
// @produce json
// @param page query int false "Page number (default: 1)"
// @param limit query int false "Number of results per page (default: 10)"
// @success 200 {array} model.Book
// @Security ApiKeyAuth
func (api Apiroutes) GetBook(ctx *gin.Context) {
	utils.Log(model.LogLevelInfo, model.ApiPackage, model.GetBook, "Fetching Bookes", nil)
	api.Server.GetBook(ctx)
}

// Handler to get a Book by ID
// @router /Book/{id} [get]
// @summary Get a Book by ID
// @tags Book
// @produce json
// @param id path string true "Book ID"
// @success 200 {object} model.Book
// @Security ApiKeyAuth
func (api Apiroutes) GetBookByID(ctx *gin.Context) {
	utils.Log(model.LogLevelInfo, model.ApiPackage, model.GetBookByID, "Fetching Book by ID", nil)
	api.Server.GetBookByID(ctx)
}

// Handler to get Books based on filters
// @router /Book/filter [get]
// @summary Get Books based on given filters
// @tags Book
// @produce json
// @param book_name query string false "Book Name"
// @param school_id query string false "School ID"
// @param active query bool false "Active status"
// @param created_by query string false "Created By"
// @param updated_by query string false "Updated By"
// @param book_id query string false "Book ID"
// @param publisher query string false "Publisher"
// @param author query string false "Author"
// @param edition_number query int false "Edition Number"
// @param page query int false "Page number (default: 1)"
// @param limit query int false "Number of results per page (default: 10)"
// @success 200 {array} model.Book
// @Security ApiKeyAuth
func (api Apiroutes) GetBookByFilter(ctx *gin.Context) {
	utils.Log(model.LogLevelInfo, model.ApiPackage, model.GetBookByFilter, "Fetching Bookes by filter", nil)
	api.Server.GetBookByFilter(ctx)
}


// Handler to update a Book
// @router /Book/update/{id} [put]
// @summary Update a Book
// @tags Book
// @accept json
// @produce json
// @param id path string true "Book ID"
// @param Book body model.Book true "Book object"
// @success 200 {object} model.Book
// @Security ApiKeyAuth
func (api Apiroutes) UpdateBook(ctx *gin.Context) {
	utils.Log(model.LogLevelInfo, model.ApiPackage, model.UpdateBook, "Updating Book", nil)
	api.Server.UpdateBook(ctx)
}

// Handler to delete a Book
// @router /Book/delete/{id} [delete]
// @summary Delete a Book
// @tags Book
// @param id path string true "Book ID"
// @Security ApiKeyAuth
func (api Apiroutes) DeleteBook(ctx *gin.Context) {
	utils.Log(model.LogLevelInfo, model.ApiPackage, model.DeleteBook, "Deleting Book", nil)
	api.Server.DeleteBook(ctx)
}
