package model

import (
	"time"

	"github.com/google/uuid"
)

type Teacher struct {
	ID          uuid.UUID `gorm:"primarykey" json:"id"`
	Active      bool      `json:"active" example:"true"` // active is used to see user is active or not active.
	FirstName   string    `json:"first_name"`
	LastName    string    `json:"last_name"`
	CreatedBy   string    `json:"created_by"  example:"user"`
	CreatedAt   time.Time `json:"created_at"`
	DeleteBy    string    `json:"delete_by" example:"admin"`
	DeletedAt   time.Time `json:"deleted_at"`
	UpdatedBy   string    `json:"updated_by" example:"admin"`
	UpdatedAt   time.Time `json:"updated_at"`
	TeacherID   uuid.UUID `json:"teacher_id" binding:"required" example:"e9138023-1301081301"`
	SubjectCode string    `json:"subject_code" binding:"required" example:"e9138023-1301081301"`
	TeacherType string    `json:"teacher_type" binding:"required" example:"assistant"`
	Salary      int       `json:"salary" binding:"required"`
	Skills      string    `json:"skills"`
	JoiningDate time.Time `json:"join_date" binding:"required"`
}
