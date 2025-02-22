package model

import (
	"time"

	"github.com/google/uuid"
)

// Book struct
type Book struct {
	ID            uuid.UUID `gorm:"primarykey" json:"id"`
	BookName      string    `json:"book_name" binding:"required"  example:"panchtantra"`
	SchoolID      uuid.UUID `binding:"required" json:"school_id"`
	Active        bool      `json:"active"  binding:"required" example:"true"`
	Publisher     string    `json:"publisher"  binding:"required"`
	Author        string    `json:"author"  binding:"required"`
	EditionNumber int       `json:"edition_number" binding:"required"  example:"2"`
	CreatedBy     string    `json:"created_by" binding:"required"  example:"vishal"`
	DeletedBy     string    `json:"deleted_by" `
	UpdatedBy     string    `json:"updated_by" `
	CreatedAt     time.Time `json:"created_at" `
	UpdatedAt     time.Time `json:"updated_at" `
	DeletedAt     time.Time `json:"deleted_at" `
}
