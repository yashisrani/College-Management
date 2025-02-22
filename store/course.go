package store

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/yashisrani/Go-Backend/model"
	"github.com/yashisrani/Go-Backend/utils"
	"gorm.io/gorm"
)

func (store Postgress) CreateCourse(Course *model.Course) error {
	utils.Log(model.LogLevelInfo, model.Store, model.CreateCourse, "creating new Course", nil)
	response := store.DB.Create(Course)
	if response.Error != nil {
		utils.Log(model.LogLevelError, model.Store, model.CreateCourse, "error while creating Course", response.Error)
	}

	utils.Log(model.LogLevelInfo, model.Store, model.CreateCourse, "Course created successfully", Course)
	return nil
}

func (store Postgress) GetCourse() ([]model.Course, error) {

	// create db layer
	Coursees := []model.Course{}

	response := store.DB.Find(&Coursees)
	if response.Error != nil {
		utils.Log(model.LogLevelError, model.Store, model.GetCourse, "error while fetching Course", response.Error)
	}

	utils.Log(model.LogLevelInfo, model.Store, model.GetCourse, "Course fetched successfully", Coursees)
	return Coursees, nil
}

func (store Postgress) GetCourseByID(CourseID uuid.UUID) (*model.Course, error) {
	// Create a Course variable
	var Course model.Course

	// Fetch Course by ID
	response := store.DB.First(&Course, CourseID)
	if response.Error != nil {
		if response.Error == gorm.ErrRecordNotFound {
			utils.Log(model.LogLevelError, model.Store, model.GetCourseByID, "Course record is not found", response.Error)
		} else {
			utils.Log(model.LogLevelError, model.Store, model.GetCourseByID, "error while fetching Course by ID", response.Error)
			return nil, response.Error
		}
	}

	utils.Log(model.LogLevelInfo, model.Store, model.GetCourseByID, "Course fetched successfully", Course)
	return &Course, nil
}

func (store Postgress) UpdateCourse(Course *model.Course) error {

	utils.Log(model.LogLevelInfo, model.StorePackage, model.UpdateCourse, "updating Course data", *Course)
	resp := store.DB.Save(Course)
	if resp.Error != nil {
		utils.Log(model.LogLevelError, model.StorePackage, model.UpdateCourse,
			"error while updating Course data", resp.Error)
		return fmt.Errorf("error while updating Course record, err = %v", resp.Error)
	}
	utils.Log(model.LogLevelInfo, model.StorePackage, model.UpdateCourse,
		"successfully updated Course", nil)
	return nil
}

func (store Postgress) DeleteCourse(CourseID string) error {

	var Course model.Course

	utils.Log(model.LogLevelInfo, model.StorePackage, model.DeleteCourse, "deleting Course data", CourseID)
	if err := store.DB.First(&Course, `"id" = '`+CourseID+`'`).Error; err != nil {
		utils.Log(model.LogLevelError, model.StorePackage, model.DeleteCourse,
			"error while deleting Course data", err)
		return fmt.Errorf("Course not found for given id, ID = %v", CourseID)
	}
	resp := store.DB.Delete(Course)
	if resp.Error != nil {
		return fmt.Errorf("error while deleting Course record from DB, err = %v", resp.Error)
	}
	utils.Log(model.LogLevelInfo, model.StorePackage, model.DeleteCourse,
		"successfully deleted Course", nil)
	return nil
}

// soft delete: not actually deleting record, but change the status of active (eg. active status is true then make it flase)
// hard delete: actually deleting record.

// func (store Postgress) SoftDeleteCourse(CourseID string) error {
// 	var Course model.Course

// 	utils.Log(model.LogLevelInfo, model.StorePackage, model.DeleteCourse, "soft deleting Course data", CourseID)

// 	// Find the Course by ID
// 	if err := store.DB.First(&Course, `"id" = '`+CourseID+`'`).Error; err != nil {
// 		utils.Log(model.LogLevelError, model.StorePackage, model.DeleteCourse,
// 			"error while finding Course data", err)
// 		return fmt.Errorf("Course not found for given ID, ID = %v", CourseID)
// 	}

// 	// Set active field to false (Soft delete)
// 	Course.Active = false
// 	if err := store.DB.Save(&Course).Error; err != nil {
// 		utils.Log(model.LogLevelError, model.StorePackage, model.DeleteCourse,
// 			"error while updating Course active status", err)
// 		return fmt.Errorf("error while updating Course record in DB, err = %v", err)
// 	}

// 	utils.Log(model.LogLevelInfo, model.StorePackage, model.DeleteCourse,
// 		"successfully soft deleted Course", CourseID)
// 	return nil
// }


func (store Postgress) GetCourseByFilter(filter map[string]interface{}) ([]model.Course, error) {

	utils.Log(model.LogLevelInfo, model.StorePackage, model.GetCourseByFilter, "fetching records of Course from db", nil)
	var Coursees []model.Course
	query := store.DB    // 

	for key, value := range filter {
		if key == model.PageNumber || key == model.DataPerPage || key == model.StartDate || key == model.EndDate {
			continue
		}
		utils.Log(model.LogLevelInfo, model.StorePackage, model.GetCourseByFilter,
			"filters key", key+" value = "+fmt.Sprintf("%v", value))
		query = query.Where(fmt.Sprintf("%s = ?", key), value)
	}
	SetLimitAndPage(filter, query)
	SetDateRangeFilter(filter, query)

	err := query.Find(&Coursees).Error
	if err != nil {
		utils.Log(model.LogLevelError, model.StorePackage, model.GetCourseByFilter,
			"error while reading Course data", err)
		return nil, fmt.Errorf("error while fetching Course records from DB, err = %v", err)
	}

	utils.Log(model.LogLevelInfo, model.StorePackage, model.GetCourseByFilter, "records of Coursees from db", Coursees)
	return Coursees, nil
}




