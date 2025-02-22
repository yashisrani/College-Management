package store

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/yashisrani/Go-Backend/model"
	"github.com/yashisrani/Go-Backend/utils"
	"gorm.io/gorm"
)

func (store Postgress) CreateClass(Class *model.Class) error {
	utils.Log(model.LogLevelInfo, model.Store, model.CreateClass, "creating new Class", nil)
	response := store.DB.Create(Class)
	if response.Error != nil {
		utils.Log(model.LogLevelError, model.Store, model.CreateClass, "error while creating Class", response.Error)
	}

	utils.Log(model.LogLevelInfo, model.Store, model.CreateClass, "Class created successfully", Class)
	return nil
}

func (store Postgress) GetClass() ([]model.Class, error) {

	// create db layer
	classes := []model.Class{}

	response := store.DB.Find(&classes)
	if response.Error != nil {
		utils.Log(model.LogLevelError, model.Store, model.GetClass, "error while fetching Class", response.Error)
	}

	utils.Log(model.LogLevelInfo, model.Store, model.GetClass, "Class fetched successfully", classes)
	return classes, nil
}

func (store Postgress) GetClassByID(ClassID uuid.UUID) (*model.Class, error) {
	// Create a Class variable
	var Class model.Class

	// Fetch Class by ID
	response := store.DB.First(&Class, ClassID)
	if response.Error != nil {
		if response.Error == gorm.ErrRecordNotFound {
			utils.Log(model.LogLevelError, model.Store, model.GetClassByID, "Class record is not found", response.Error)
		} else {
			utils.Log(model.LogLevelError, model.Store, model.GetClassByID, "error while fetching Class by ID", response.Error)
			return nil, response.Error
		}
	}

	utils.Log(model.LogLevelInfo, model.Store, model.GetClassByID, "Class fetched successfully", Class)
	return &Class, nil
}

func (store Postgress) UpdateClass(Class *model.Class) error {

	utils.Log(model.LogLevelInfo, model.StorePackage, model.UpdateClass, "updating Class data", *Class)
	resp := store.DB.Save(Class)
	if resp.Error != nil {
		utils.Log(model.LogLevelError, model.StorePackage, model.UpdateClass,
			"error while updating Class data", resp.Error)
		return fmt.Errorf("error while updating Class record, err = %v", resp.Error)
	}
	utils.Log(model.LogLevelInfo, model.StorePackage, model.UpdateClass,
		"successfully updated Class", nil)
	return nil
}

func (store Postgress) DeleteClass(ClassID string) error {

	var Class model.Class

	utils.Log(model.LogLevelInfo, model.StorePackage, model.DeleteClass, "deleting Class data", ClassID)
	if err := store.DB.First(&Class, `"id" = '`+ClassID+`'`).Error; err != nil {
		utils.Log(model.LogLevelError, model.StorePackage, model.DeleteClass,
			"error while deleting Class data", err)
		return fmt.Errorf("Class not found for given id, ID = %v", ClassID)
	}
	resp := store.DB.Delete(Class)
	if resp.Error != nil {
		return fmt.Errorf("error while deleting Class record from DB, err = %v", resp.Error)
	}
	utils.Log(model.LogLevelInfo, model.StorePackage, model.DeleteClass,
		"successfully deleted Class", nil)
	return nil
}

// soft delete: not actually deleting record, but change the status of active (eg. active status is true then make it flase)
// hard delete: actually deleting record.

// func (store Postgress) SoftDeleteClass(ClassID string) error {
// 	var Class model.Class

// 	utils.Log(model.LogLevelInfo, model.StorePackage, model.DeleteClass, "soft deleting Class data", ClassID)

// 	// Find the Class by ID
// 	if err := store.DB.First(&Class, `"id" = '`+ClassID+`'`).Error; err != nil {
// 		utils.Log(model.LogLevelError, model.StorePackage, model.DeleteClass,
// 			"error while finding Class data", err)
// 		return fmt.Errorf("Class not found for given ID, ID = %v", ClassID)
// 	}

// 	// Set active field to false (Soft delete)
// 	Class.Active = false
// 	if err := store.DB.Save(&Class).Error; err != nil {
// 		utils.Log(model.LogLevelError, model.StorePackage, model.DeleteClass,
// 			"error while updating Class active status", err)
// 		return fmt.Errorf("error while updating Class record in DB, err = %v", err)
// 	}

// 	utils.Log(model.LogLevelInfo, model.StorePackage, model.DeleteClass,
// 		"successfully soft deleted Class", ClassID)
// 	return nil
// }


func (store Postgress) GetClassByFilter(filter map[string]interface{}) ([]model.Class, error) {

	utils.Log(model.LogLevelInfo, model.StorePackage, model.GetClassByFilter, "fetching records of Class from db", nil)
	var classes []model.Class
	query := store.DB    // 

	for key, value := range filter {
		if key == model.PageNumber || key == model.DataPerPage || key == model.StartDate || key == model.EndDate {
			continue
		}
		utils.Log(model.LogLevelInfo, model.StorePackage, model.GetClassByFilter,
			"filters key", key+" value = "+fmt.Sprintf("%v", value))
		query = query.Where(fmt.Sprintf("%s = ?", key), value)
	}
	SetLimitAndPage(filter, query)
	SetDateRangeFilter(filter, query)

	err := query.Find(&classes).Error
	if err != nil {
		utils.Log(model.LogLevelError, model.StorePackage, model.GetClassByFilter,
			"error while reading Class data", err)
		return nil, fmt.Errorf("error while fetching Class records from DB, err = %v", err)
	}

	utils.Log(model.LogLevelInfo, model.StorePackage, model.GetClassByFilter, "records of classes from db", classes)
	return classes, nil
}




