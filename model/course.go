package model

import (
	"time"

	"github.com/google/uuid"
)

// Course struct
type Course struct {
	ID            uuid.UUID `gorm:"primarykey" json:"id"`
	CourseName    string    `json:"course_name" binding:"required" example:"Full Stack Development"`
	SchoolID      uuid.UUID `binding:"required" json:"school_id"`
	Active        bool      `json:"active" binding:"required" example:"true"`
	Instructor    string    `json:"instructor" binding:"required" example:"Jane Doe"`
	DurationWeeks int       `json:"duration_weeks" binding:"required" example:"12"` // Duration in weeks
	Credits       int       `json:"credits" binding:"required" example:"3"` // Course credits
	Fee           float64   `json:"fee" binding:"required" example:"15000.00"`
	Description   string    `json:"description" binding:"required" example:"An advanced course covering both front-end and back-end development."`
	Syllabus      string    `json:"syllabus" binding:"required" example:"HTML, CSS, JavaScript, React, Node.js, Databases"`
	StartDate     time.Time `json:"start_date" binding:"required"`
	EndDate       time.Time `json:"end_date" binding:"required"`
	CreatedBy     string    `json:"created_by" binding:"required" example:"admin"`
	DeletedBy     string    `json:"deleted_by"`
	UpdatedBy     string    `json:"updated_by"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
	DeletedAt     time.Time `json:"deleted_at"`
}
