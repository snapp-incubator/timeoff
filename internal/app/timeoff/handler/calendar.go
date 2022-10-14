package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/rickar/cal/v2"
)

// CalendarHandler represents Calendar HTTP handler.
type CalendarHandler struct {
	baseCalendar *cal.BusinessCalendar
}

// NewCalendarHandler returns a new instance of CalendarHandler.
func NewCalendarHandler(baseCal *cal.BusinessCalendar) *CalendarHandler {
	return &CalendarHandler{
		baseCalendar: baseCal,
	}
}

// GetHolidays returns calendar holidays.
func (h *CalendarHandler) GetHolidays(echoContext echo.Context) error {
	return echoContext.JSON(http.StatusOK, "")
}

// AddHolidays adds a date to current calendar holidays.
func (h *CalendarHandler) AddHolidays(echoContext echo.Context) error {
	return echoContext.JSON(http.StatusOK, "")
}

// RemoveHolidays removes a date from current calendar holidays.
func (h *CalendarHandler) RemoveHolidays(echoContext echo.Context) error {
	return echoContext.JSON(http.StatusOK, "")
}
