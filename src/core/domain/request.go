package domain

import "time"

type DateRangeRequest struct {
	StartDate time.Time `json:"startDate" binding:"required"`
	EndDate   time.Time `json:"endDate" binding:"required,gtfield=StartDate"`
}
