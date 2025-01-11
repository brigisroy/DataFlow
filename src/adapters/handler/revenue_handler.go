package handler

import (
	"fmt"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"sample.code/dataflow/src/core/domain"
	"sample.code/dataflow/src/core/ports"
)

type revenueHandler struct {
	revenueService ports.RevenueService
	validate       *validator.Validate
}

func NewRevenueHandler(rs ports.RevenueService) ports.RevenueHandler {
	return &revenueHandler{
		revenueService: rs,
		validate:       validator.New(),
	}
}

func (rh *revenueHandler) FetchRevenue(e echo.Context) error {
	var request domain.DateRangeRequest
	ctx := e.Request().Context()
	if err := e.Bind(&request); err != nil {
		return e.JSON(http.StatusBadRequest, fmt.Sprintf("Invalid date range: %v", err))
	}
	if err := rh.validate.Struct(request); err != nil {
		// Check for validation errors and handle them
		if _, ok := err.(*validator.InvalidValidationError); ok {
			return e.JSON(http.StatusBadRequest, fmt.Sprintf("Invalid date range: %v", err))
		}

		// Additional validation for date range: StartDate must be before EndDate
		for _, err := range err.(validator.ValidationErrors) {
			if err.Field() == "StartDate" && err.Tag() == "ltefield" {
				return e.JSON(http.StatusBadRequest, "Start date must be before or equal to End date.")
			}
			if err.Field() == "EndDate" && err.Tag() == "gtefield" {
				return e.JSON(http.StatusBadRequest, "End date must be after Start date.")
			}
		}
	}

	revenue, err := rh.revenueService.GetRevenueDetails(ctx, request)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, fmt.Sprintf("Error getting revenue: %v", err))
	}
	return e.JSON(http.StatusOK, revenue)
}
