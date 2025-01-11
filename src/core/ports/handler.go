package ports

import "github.com/labstack/echo/v4"

type RevenueHandler interface {
	FetchRevenue(e echo.Context) error
}
