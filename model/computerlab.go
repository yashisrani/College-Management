package model

import (
	"time"

	"github.com/google/uuid"
)

// ComputerClass struct
type ComputerLab struct {
	ID              uuid.UUID `gorm:"primarykey" json:"id"`
	ClassName       string    `json:"class_name" binding:"required" example:"Advanced Programming"`
	SchoolID        uuid.UUID `binding:"required" json:"school_id"`
	Active          bool      `json:"active" binding:"required" example:"true"`
	Instructor      string    `json:"instructor" binding:"required" example:"John Doe"`
	TotalStudents   int       `json:"total_students" binding:"required" example:"30"`
	ClassroomNumber string    `json:"classroom_number" binding:"required" example:"C-202"`
	Syllabus        string    `json:"syllabus" binding:"required" example:"Introduction to Python, Data Structures"`
	CreatedBy       string    `json:"created_by" binding:"required" example:"admin"`
	DeletedBy       string    `json:"deleted_by"`
	UpdatedBy       string    `json:"updated_by"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
	DeletedAt       time.Time `json:"deleted_at"`
}
