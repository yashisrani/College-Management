package model

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID        uuid.UUID `gorm:"primarykey" json:"id"`
	Name      Name      `gorm:"embedded" json:"name" binding:"required"` // gorm:embedded to add name and address struct to main struct.
	Address   Address   `gorm:"embedded" json:"address" binding:"required"`
	Active    bool      `json:"active" example:"true"` // active is used to see user is active or not active.
	// Email     string    `json:"email" gorm:"unique;notnull" binding:"required"`
	// Password  string    `json:"password" gorm:"notnull" binding:"required"`
	CreatedBy string    `json:"created_by"  example:"user"`
	CreatedAt time.Time `json:"created_at"`
	DeleteBy  string    `json:"delete_by" example:"admin"`
	DeletedAt time.Time `json:"deleted_at"`
	UpdatedBy string    `json:"updated_by" example:"admin"`
	UpdatedAt time.Time `json:"updated_at"`
}

type Name struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

type Address struct {
	Lane     string `json:"lane"`
	Village  string `json:"village"`
	City     string `json:"city"`
	District string `json:"district"`
	Pincode  int    `json:"pincode"`
	State    string `json:"state"`
}
