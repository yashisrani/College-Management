package store

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/yashisrani/Go-Backend/model"
	"github.com/yashisrani/Go-Backend/utils"
	"gorm.io/gorm"
)

func (store Postgress) CreateTeacher(Teacher *model.Teacher) error {
	utils.Log(model.LogLevelInfo, model.Store, model.CreateTeacher, "creating new Teacher", nil)
	response := store.DB.Create(Teacher)
	if response.Error != nil {
		utils.Log(model.LogLevelError, model.Store, model.CreateTeacher, "error while creating Teacher", response.Error)
	}

	utils.Log(model.LogLevelInfo, model.Store, model.CreateTeacher, "Teacher created successfully", Teacher)
	return nil
}

func (store Postgress) GetTeacher() ([]model.Teacher, error) {

	// create db layer
	Teacheres := []model.Teacher{}

	response := store.DB.Find(&Teacheres)
	if response.Error != nil {
		utils.Log(model.LogLevelError, model.Store, model.GetTeacher, "error while fetching Teacher", response.Error)
	}

	utils.Log(model.LogLevelInfo, model.Store, model.GetTeacher, "Teacher fetched successfully", Teacheres)
	return Teacheres, nil
}

func (store Postgress) GetTeacherByID(TeacherID uuid.UUID) (*model.Teacher, error) {
	// Create a Teacher variable
	var Teacher model.Teacher

	// Fetch Teacher by ID
	response := store.DB.First(&Teacher, TeacherID)
	if response.Error != nil {
		if response.Error == gorm.ErrRecordNotFound {
			utils.Log(model.LogLevelError, model.Store, model.GetTeacherByID, "Teacher record is not found", response.Error)
		} else {
			utils.Log(model.LogLevelError, model.Store, model.GetTeacherByID, "error while fetching Teacher by ID", response.Error)
			return nil, response.Error
		}
	}

	utils.Log(model.LogLevelInfo, model.Store, model.GetTeacherByID, "Teacher fetched successfully", Teacher)
	return &Teacher, nil
}

func (store Postgress) UpdateTeacher(Teacher *model.Teacher) error {

	utils.Log(model.LogLevelInfo, model.StorePackage, model.UpdateTeacher, "updating Teacher data", *Teacher)
	resp := store.DB.Save(Teacher)
	if resp.Error != nil {
		utils.Log(model.LogLevelError, model.StorePackage, model.UpdateTeacher,
			"error while updating Teacher data", resp.Error)
		return fmt.Errorf("error while updating Teacher record, err = %v", resp.Error)
	}
	utils.Log(model.LogLevelInfo, model.StorePackage, model.UpdateTeacher,
		"successfully updated Teacher", nil)
	return nil
}

func (store Postgress) DeleteTeacher(TeacherID string) error {

	var Teacher model.Teacher

	utils.Log(model.LogLevelInfo, model.StorePackage, model.DeleteTeacher, "deleting Teacher data", TeacherID)
	if err := store.DB.First(&Teacher, `"id" = '`+TeacherID+`'`).Error; err != nil {
		utils.Log(model.LogLevelError, model.StorePackage, model.DeleteTeacher,
			"error while deleting Teacher data", err)
		return fmt.Errorf("Teacher not found for given id, ID = %v", TeacherID)
	}
	resp := store.DB.Delete(Teacher)
	if resp.Error != nil {
		return fmt.Errorf("error while deleting Teacher record from DB, err = %v", resp.Error)
	}
	utils.Log(model.LogLevelInfo, model.StorePackage, model.DeleteTeacher,
		"successfully deleted Teacher", nil)
	return nil
}

// soft delete: not actually deleting record, but change the status of active (eg. active status is true then make it flase)
// hard delete: actually deleting record.

// func (store Postgress) SoftDeleteTeacher(TeacherID string) error {
// 	var Teacher model.Teacher

// 	utils.Log(model.LogLevelInfo, model.StorePackage, model.DeleteTeacher, "soft deleting Teacher data", TeacherID)

// 	// Find the Teacher by ID
// 	if err := store.DB.First(&Teacher, `"id" = '`+TeacherID+`'`).Error; err != nil {
// 		utils.Log(model.LogLevelError, model.StorePackage, model.DeleteTeacher,
// 			"error while finding Teacher data", err)
// 		return fmt.Errorf("Teacher not found for given ID, ID = %v", TeacherID)
// 	}

// 	// Set active field to false (Soft delete)
// 	Teacher.Active = false
// 	if err := store.DB.Save(&Teacher).Error; err != nil {
// 		utils.Log(model.LogLevelError, model.StorePackage, model.DeleteTeacher,
// 			"error while updating Teacher active status", err)
// 		return fmt.Errorf("error while updating Teacher record in DB, err = %v", err)
// 	}

// 	utils.Log(model.LogLevelInfo, model.StorePackage, model.DeleteTeacher,
// 		"successfully soft deleted Teacher", TeacherID)
// 	return nil
// }


func (store Postgress) GetTeacherByFilter(filter map[string]interface{}) ([]model.Teacher, error) {

	utils.Log(model.LogLevelInfo, model.StorePackage, model.GetTeacherByFilter, "fetching records of Teacher from db", nil)
	var Teacheres []model.Teacher
	query := store.DB    // 

	for key, value := range filter {
		if key == model.PageNumber || key == model.DataPerPage || key == model.StartDate || key == model.EndDate {
			continue
		}
		utils.Log(model.LogLevelInfo, model.StorePackage, model.GetTeacherByFilter,
			"filters key", key+" value = "+fmt.Sprintf("%v", value))
		query = query.Where(fmt.Sprintf("%s = ?", key), value)
	}
	SetLimitAndPage(filter, query)
	SetDateRangeFilter(filter, query)

	err := query.Find(&Teacheres).Error
	if err != nil {
		utils.Log(model.LogLevelError, model.StorePackage, model.GetTeacherByFilter,
			"error while reading Teacher data", err)
		return nil, fmt.Errorf("error while fetching Teacher records from DB, err = %v", err)
	}

	utils.Log(model.LogLevelInfo, model.StorePackage, model.GetTeacherByFilter, "records of Teacheres from db", Teacheres)
	return Teacheres, nil
}




