package model

import (
	"time"

	"github.com/google/uuid"
)

type College struct {
	ID             uuid.UUID `gorm:"primarykey" json:"id"`
	Active         bool      `json:"active" example:"true"` // active is used to see user is active or not active.
	Domain         string    `json:"domain" gorm:"not null" binding:"required" example:"noble"`
	CreatedBy      string    `json:"created_by"  example:"user"`
	CreatedAt      time.Time `json:"created_at"`
	DeleteBy       string    `json:"delete_by" example:"admin"`
	DeletedAt      time.Time `json:"deleted_at"`
	UpdatedBy      string    `json:"updated_by" example:"admin"`
	UpdatedAt      time.Time `json:"updated_at"`
	Lane           string    `json:"lane"`
	Village        string    `json:"village"`
	City           string    `json:"city" binding:"required"`
	District       string    `json:"district" binding:"required"`
	Pincode        int       `json:"pincode" binding:"required"`
	State          string    `json:"state" binding:"required"`
	PrincipleID    uuid.UUID `json:"principle_id"`
	AdminID        uuid.UUID `json:"admin_id"`
	DirectorID     uuid.UUID `json:"director_id" binding:"required" example:"e9138023-1301081301"`
	UniversityType string    `json:"university_type" gorm:"not null" example:"private college" binding:"required"`
	Hostel         bool      `json:"hostel" gorm:"not null" binding:"required"`
}
