package store

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/yashisrani/Go-Backend/model"
	"github.com/yashisrani/Go-Backend/utils"
	"gorm.io/gorm"
)

func (store Postgress) CreateCollege(college *model.College) error {
	utils.Log(model.LogLevelInfo, model.Store, model.CreateCollege, "creating new college", nil)
	response := store.DB.Create(college)
	if response.Error != nil {
		utils.Log(model.LogLevelError, model.Store, model.CreateCollege, "error while creating college", response.Error)
	}

	utils.Log(model.LogLevelInfo, model.Store, model.CreateCollege, "college created successfully", college)
	return nil
}

func (store Postgress) GetColleges() ([]model.College, error) {

	// create db layer
	colleges := []model.College{}

	response := store.DB.Find(&colleges)
	if response.Error != nil {
		utils.Log(model.LogLevelError, model.Store, model.GetCollege, "error while fetching college", response.Error)
	}

	utils.Log(model.LogLevelInfo, model.Store, model.GetCollege, "college fetched successfully", colleges)
	return colleges, nil
}

func (store Postgress) GetCollegeByID(collegeID uuid.UUID) (*model.College, error) {
	// Create a college variable
	var college model.College

	// Fetch college by ID
	response := store.DB.First(&college, collegeID)
	if response.Error != nil {
		if response.Error == gorm.ErrRecordNotFound {
			utils.Log(model.LogLevelError, model.Store, model.GetCollegeByID, "college record is not found", response.Error)
		} else {
			utils.Log(model.LogLevelError, model.Store, model.GetCollegeByID, "error while fetching college by ID", response.Error)
			return nil, response.Error
		}
	}

	utils.Log(model.LogLevelInfo, model.Store, model.GetCollegeByID, "college fetched successfully", college)
	return &college, nil
}

func (store Postgress) UpdateCollege(college *model.College) error {

	utils.Log(model.LogLevelInfo, model.StorePackage, model.UpdateCollege, "updating college data", *college)
	resp := store.DB.Save(college)
	if resp.Error != nil {
		utils.Log(model.LogLevelError, model.StorePackage, model.UpdateCollege,
			"error while updating college data", resp.Error)
		return fmt.Errorf("error while updating college record, err = %v", resp.Error)
	}
	utils.Log(model.LogLevelInfo, model.StorePackage, model.UpdateCollege,
		"successfully updated college", nil)
	return nil
}

func (store Postgress) DeleteCollege(collegeID string) error {

	var college model.College

	utils.Log(model.LogLevelInfo, model.StorePackage, model.DeleteCollege, "deleting college data", collegeID)
	if err := store.DB.First(&college, `"id" = '`+collegeID+`'`).Error; err != nil {
		utils.Log(model.LogLevelError, model.StorePackage, model.DeleteCollege,
			"error while deleting college data", err)
		return fmt.Errorf("college not found for given id, ID = %v", collegeID)
	}
	resp := store.DB.Delete(college)
	if resp.Error != nil {
		return fmt.Errorf("error while deleting college record from DB, err = %v", resp.Error)
	}
	utils.Log(model.LogLevelInfo, model.StorePackage, model.DeleteCollege,
		"successfully deleted college", nil)
	return nil
}

// soft delete: not actually deleting record, but change the status of active (eg. active status is true then make it flase)
// hard delete: actually deleting record.

// func (store Postgress) SoftDeletecollege(collegeID string) error {
// 	var college model.college

// 	utils.Log(model.LogLevelInfo, model.StorePackage, model.Deletecollege, "soft deleting college data", collegeID)

// 	// Find the college by ID
// 	if err := store.DB.First(&college, `"id" = '`+collegeID+`'`).Error; err != nil {
// 		utils.Log(model.LogLevelError, model.StorePackage, model.Deletecollege,
// 			"error while finding college data", err)
// 		return fmt.Errorf("college not found for given ID, ID = %v", collegeID)
// 	}

// 	// Set active field to false (Soft delete)
// 	college.Active = false
// 	if err := store.DB.Save(&college).Error; err != nil {
// 		utils.Log(model.LogLevelError, model.StorePackage, model.Deletecollege,
// 			"error while updating college active status", err)
// 		return fmt.Errorf("error while updating college record in DB, err = %v", err)
// 	}

// 	utils.Log(model.LogLevelInfo, model.StorePackage, model.Deletecollege,
// 		"successfully soft deleted college", collegeID)
// 	return nil
// }


func (store Postgress) GetCollegeByFilter(filter map[string]interface{}) ([]model.College, error) {

	utils.Log(model.LogLevelInfo, model.StorePackage, model.GetCollegeByFilter, "fetching records of college from db", nil)
	var colleges []model.College
	query := store.DB    // 

	for key, value := range filter {
		if key == model.PageNumber || key == model.DataPerPage || key == model.StartDate || key == model.EndDate {
			continue
		}
		utils.Log(model.LogLevelInfo, model.StorePackage, model.GetCollegeByFilter,
			"filters key", key+" value = "+fmt.Sprintf("%v", value))
		query = query.Where(fmt.Sprintf("%s = ?", key), value)
	}
	// setLimitAndPage(filter, query)
	// setDateRangeFilter(filter, query)

	err := query.Find(&colleges).Error
	if err != nil {
		utils.Log(model.LogLevelError, model.StorePackage, model.GetCollegeByFilter,
			"error while reading college data", err)
		return nil, fmt.Errorf("error while fetching college records from DB, err = %v", err)
	}

	utils.Log(model.LogLevelInfo, model.StorePackage, model.GetCollegeByFilter, "records of colleges from db", colleges)
	return colleges, nil
}


