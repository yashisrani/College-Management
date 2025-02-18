package model

import (
	"time"
)

type CommonParamerers struct {
	Name      string    `json:"name"`
	StartTime time.Time `json:"start_time"`
	EndTime   time.Time `json:"end_time"`
}

// ErrorResponse struct
type ErrorResponse struct {
	Message string `json:"error"`
}

// SuccessResponse struct
type SuccessResponse struct {
	Message string `json:"message"`
}