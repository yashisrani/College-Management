package model

import (
	"time"

	"github.com/google/uuid"
)

type Class struct {
	ID              uuid.UUID `gorm:"primarykey" json:"id"`
	Active          bool      `json:"active" example:"true"` // active is used to see user is active or not active.
	NumberOfStudent int       `json:"number_of_student" gorm:"not null" binding:"required" example:"10"`
	NumberofBench   int       `json:"number_of_bench" gorm:"not null" binding:"required" example:"10"`
	CreatedBy       string    `json:"created_by"  example:"user"`
	CreatedAt       time.Time `json:"created_at"`
	DeleteBy        string    `json:"delete_by" example:"admin"`
	DeletedAt       time.Time `json:"deleted_at"`
	UpdatedBy       string    `json:"updated_by" example:"admin"`
	UpdatedAt       time.Time `json:"updated_at"`
	TeacherID       uuid.UUID `json:"teacher_id" binding:"required" example:"e9138023-1301081301"`
	ClassMonitor    string    `json:"class_monitor"`
	FloorNumber     int       `json:"floor_number" binding:"required" example:"10"`
	ClassName       string    `json:"class_name" binding:"required" example:"class-10"`
}
