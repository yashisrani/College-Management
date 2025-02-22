package store

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/yashisrani/Go-Backend/model"
	"github.com/yashisrani/Go-Backend/utils"
	"gorm.io/gorm"
)

func (store Postgress) CreateComputerLab(ComputerLab *model.ComputerLab) error {
	utils.Log(model.LogLevelInfo, model.Store, model.CreateComputerLab, "creating new ComputerLab", nil)
	response := store.DB.Create(ComputerLab)
	if response.Error != nil {
		utils.Log(model.LogLevelError, model.Store, model.CreateComputerLab, "error while creating ComputerLab", response.Error)
	}

	utils.Log(model.LogLevelInfo, model.Store, model.CreateComputerLab, "ComputerLab created successfully", ComputerLab)
	return nil
}

func (store Postgress) GetComputerLab() ([]model.ComputerLab, error) {

	// create db layer
	ComputerLabes := []model.ComputerLab{}

	response := store.DB.Find(&ComputerLabes)
	if response.Error != nil {
		utils.Log(model.LogLevelError, model.Store, model.GetComputerLab, "error while fetching ComputerLab", response.Error)
	}

	utils.Log(model.LogLevelInfo, model.Store, model.GetComputerLab, "ComputerLab fetched successfully", ComputerLabes)
	return ComputerLabes, nil
}

func (store Postgress) GetComputerLabByID(ComputerLabID uuid.UUID) (*model.ComputerLab, error) {
	// Create a ComputerLab variable
	var ComputerLab model.ComputerLab

	// Fetch ComputerLab by ID
	response := store.DB.First(&ComputerLab, ComputerLabID)
	if response.Error != nil {
		if response.Error == gorm.ErrRecordNotFound {
			utils.Log(model.LogLevelError, model.Store, model.GetComputerLabByID, "ComputerLab record is not found", response.Error)
		} else {
			utils.Log(model.LogLevelError, model.Store, model.GetComputerLabByID, "error while fetching ComputerLab by ID", response.Error)
			return nil, response.Error
		}
	}

	utils.Log(model.LogLevelInfo, model.Store, model.GetComputerLabByID, "ComputerLab fetched successfully", ComputerLab)
	return &ComputerLab, nil
}

func (store Postgress) UpdateComputerLab(ComputerLab *model.ComputerLab) error {

	utils.Log(model.LogLevelInfo, model.StorePackage, model.UpdateComputerLab, "updating ComputerLab data", *ComputerLab)
	resp := store.DB.Save(ComputerLab)
	if resp.Error != nil {
		utils.Log(model.LogLevelError, model.StorePackage, model.UpdateComputerLab,
			"error while updating ComputerLab data", resp.Error)
		return fmt.Errorf("error while updating ComputerLab record, err = %v", resp.Error)
	}
	utils.Log(model.LogLevelInfo, model.StorePackage, model.UpdateComputerLab,
		"successfully updated ComputerLab", nil)
	return nil
}

func (store Postgress) DeleteComputerLab(ComputerLabID string) error {

	var ComputerLab model.ComputerLab

	utils.Log(model.LogLevelInfo, model.StorePackage, model.DeleteComputerLab, "deleting ComputerLab data", ComputerLabID)
	if err := store.DB.First(&ComputerLab, `"id" = '`+ComputerLabID+`'`).Error; err != nil {
		utils.Log(model.LogLevelError, model.StorePackage, model.DeleteComputerLab,
			"error while deleting ComputerLab data", err)
		return fmt.Errorf("ComputerLab not found for given id, ID = %v", ComputerLabID)
	}
	resp := store.DB.Delete(ComputerLab)
	if resp.Error != nil {
		return fmt.Errorf("error while deleting ComputerLab record from DB, err = %v", resp.Error)
	}
	utils.Log(model.LogLevelInfo, model.StorePackage, model.DeleteComputerLab,
		"successfully deleted ComputerLab", nil)
	return nil
}

// soft delete: not actually deleting record, but change the status of active (eg. active status is true then make it flase)
// hard delete: actually deleting record.

// func (store Postgress) SoftDeleteComputerLab(ComputerLabID string) error {
// 	var ComputerLab model.ComputerLab

// 	utils.Log(model.LogLevelInfo, model.StorePackage, model.DeleteComputerLab, "soft deleting ComputerLab data", ComputerLabID)

// 	// Find the ComputerLab by ID
// 	if err := store.DB.First(&ComputerLab, `"id" = '`+ComputerLabID+`'`).Error; err != nil {
// 		utils.Log(model.LogLevelError, model.StorePackage, model.DeleteComputerLab,
// 			"error while finding ComputerLab data", err)
// 		return fmt.Errorf("ComputerLab not found for given ID, ID = %v", ComputerLabID)
// 	}

// 	// Set active field to false (Soft delete)
// 	ComputerLab.Active = false
// 	if err := store.DB.Save(&ComputerLab).Error; err != nil {
// 		utils.Log(model.LogLevelError, model.StorePackage, model.DeleteComputerLab,
// 			"error while updating ComputerLab active status", err)
// 		return fmt.Errorf("error while updating ComputerLab record in DB, err = %v", err)
// 	}

// 	utils.Log(model.LogLevelInfo, model.StorePackage, model.DeleteComputerLab,
// 		"successfully soft deleted ComputerLab", ComputerLabID)
// 	return nil
// }


func (store Postgress) GetComputerLabByFilter(filter map[string]interface{}) ([]model.ComputerLab, error) {

	utils.Log(model.LogLevelInfo, model.StorePackage, model.GetComputerLabByFilter, "fetching records of ComputerLab from db", nil)
	var ComputerLabes []model.ComputerLab
	query := store.DB    // 

	for key, value := range filter {
		if key == model.PageNumber || key == model.DataPerPage || key == model.StartDate || key == model.EndDate {
			continue
		}
		utils.Log(model.LogLevelInfo, model.StorePackage, model.GetComputerLabByFilter,
			"filters key", key+" value = "+fmt.Sprintf("%v", value))
		query = query.Where(fmt.Sprintf("%s = ?", key), value)
	}
	SetLimitAndPage(filter, query)
	SetDateRangeFilter(filter, query)

	err := query.Find(&ComputerLabes).Error
	if err != nil {
		utils.Log(model.LogLevelError, model.StorePackage, model.GetComputerLabByFilter,
			"error while reading ComputerLab data", err)
		return nil, fmt.Errorf("error while fetching ComputerLab records from DB, err = %v", err)
	}

	utils.Log(model.LogLevelInfo, model.StorePackage, model.GetComputerLabByFilter, "records of ComputerLabes from db", ComputerLabes)
	return ComputerLabes, nil
}




