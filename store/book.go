package store

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/yashisrani/Go-Backend/model"
	"github.com/yashisrani/Go-Backend/utils"
	"gorm.io/gorm"
)

func (store Postgress) CreateBook(Book *model.Book) error {
	utils.Log(model.LogLevelInfo, model.Store, model.CreateBook, "creating new Book", nil)
	response := store.DB.Create(Book)
	if response.Error != nil {
		utils.Log(model.LogLevelError, model.Store, model.CreateBook, "error while creating Book", response.Error)
	}

	utils.Log(model.LogLevelInfo, model.Store, model.CreateBook, "Book created successfully", Book)
	return nil
}

func (store Postgress) GetBook() ([]model.Book, error) {

	// create db layer
	Bookes := []model.Book{}

	response := store.DB.Find(&Bookes)
	if response.Error != nil {
		utils.Log(model.LogLevelError, model.Store, model.GetBook, "error while fetching Book", response.Error)
	}

	utils.Log(model.LogLevelInfo, model.Store, model.GetBook, "Book fetched successfully", Bookes)
	return Bookes, nil
}

func (store Postgress) GetBookByID(BookID uuid.UUID) (*model.Book, error) {
	// Create a Book variable
	var Book model.Book

	// Fetch Book by ID
	response := store.DB.First(&Book, BookID)
	if response.Error != nil {
		if response.Error == gorm.ErrRecordNotFound {
			utils.Log(model.LogLevelError, model.Store, model.GetBookByID, "Book record is not found", response.Error)
		} else {
			utils.Log(model.LogLevelError, model.Store, model.GetBookByID, "error while fetching Book by ID", response.Error)
			return nil, response.Error
		}
	}

	utils.Log(model.LogLevelInfo, model.Store, model.GetBookByID, "Book fetched successfully", Book)
	return &Book, nil
}

func (store Postgress) UpdateBook(Book *model.Book) error {

	utils.Log(model.LogLevelInfo, model.StorePackage, model.UpdateBook, "updating Book data", *Book)
	resp := store.DB.Save(Book)
	if resp.Error != nil {
		utils.Log(model.LogLevelError, model.StorePackage, model.UpdateBook,
			"error while updating Book data", resp.Error)
		return fmt.Errorf("error while updating Book record, err = %v", resp.Error)
	}
	utils.Log(model.LogLevelInfo, model.StorePackage, model.UpdateBook,
		"successfully updated Book", nil)
	return nil
}

func (store Postgress) DeleteBook(BookID string) error {

	var Book model.Book

	utils.Log(model.LogLevelInfo, model.StorePackage, model.DeleteBook, "deleting Book data", BookID)
	if err := store.DB.First(&Book, `"id" = '`+BookID+`'`).Error; err != nil {
		utils.Log(model.LogLevelError, model.StorePackage, model.DeleteBook,
			"error while deleting Book data", err)
		return fmt.Errorf("Book not found for given id, ID = %v", BookID)
	}
	resp := store.DB.Delete(Book)
	if resp.Error != nil {
		return fmt.Errorf("error while deleting Book record from DB, err = %v", resp.Error)
	}
	utils.Log(model.LogLevelInfo, model.StorePackage, model.DeleteBook,
		"successfully deleted Book", nil)
	return nil
}

// soft delete: not actually deleting record, but change the status of active (eg. active status is true then make it flase)
// hard delete: actually deleting record.

// func (store Postgress) SoftDeleteBook(BookID string) error {
// 	var Book model.Book

// 	utils.Log(model.LogLevelInfo, model.StorePackage, model.DeleteBook, "soft deleting Book data", BookID)

// 	// Find the Book by ID
// 	if err := store.DB.First(&Book, `"id" = '`+BookID+`'`).Error; err != nil {
// 		utils.Log(model.LogLevelError, model.StorePackage, model.DeleteBook,
// 			"error while finding Book data", err)
// 		return fmt.Errorf("Book not found for given ID, ID = %v", BookID)
// 	}

// 	// Set active field to false (Soft delete)
// 	Book.Active = false
// 	if err := store.DB.Save(&Book).Error; err != nil {
// 		utils.Log(model.LogLevelError, model.StorePackage, model.DeleteBook,
// 			"error while updating Book active status", err)
// 		return fmt.Errorf("error while updating Book record in DB, err = %v", err)
// 	}

// 	utils.Log(model.LogLevelInfo, model.StorePackage, model.DeleteBook,
// 		"successfully soft deleted Book", BookID)
// 	return nil
// }


func (store Postgress) GetBookByFilter(filter map[string]interface{}) ([]model.Book, error) {

	utils.Log(model.LogLevelInfo, model.StorePackage, model.GetBookByFilter, "fetching records of Book from db", nil)
	var Bookes []model.Book
	query := store.DB    // 

	for key, value := range filter {
		if key == model.PageNumber || key == model.DataPerPage || key == model.StartDate || key == model.EndDate {
			continue
		}
		utils.Log(model.LogLevelInfo, model.StorePackage, model.GetBookByFilter,
			"filters key", key+" value = "+fmt.Sprintf("%v", value))
		query = query.Where(fmt.Sprintf("%s = ?", key), value)
	}
	SetLimitAndPage(filter, query)
	SetDateRangeFilter(filter, query)

	err := query.Find(&Bookes).Error
	if err != nil {
		utils.Log(model.LogLevelError, model.StorePackage, model.GetBookByFilter,
			"error while reading Book data", err)
		return nil, fmt.Errorf("error while fetching Book records from DB, err = %v", err)
	}

	utils.Log(model.LogLevelInfo, model.StorePackage, model.GetBookByFilter, "records of Bookes from db", Bookes)
	return Bookes, nil
}




