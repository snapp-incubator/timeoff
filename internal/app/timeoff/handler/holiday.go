package handler

import (
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/rickar/cal/v2"
	"github.com/sirupsen/logrus"
	ptime "github.com/yaa110/go-persian-calendar"
	"github.com/zeinababbasi/timeoff/internal/app/timeoff/request"
	"github.com/zeinababbasi/timeoff/internal/app/timeoff/response"
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
func (holidayHandler *CalendarHandler) GetHolidays(echoContext echo.Context) error {
	irHolidays := make([]response.Holiday, len(holidayHandler.baseCalendar.Holidays))

	for index, holiday := range holidayHandler.baseCalendar.Holidays {
		irHolidays[index] = response.Holiday{
			Name:  holiday.Name,
			Type:  holidayHandler.observanceTypeDesc(holiday.Type),
			Month: int(holiday.Month),
			Day:   holiday.Day,
		}
	}

	resp := []response.Holidays{
		{
			Location: "ir",
			Holidays: irHolidays,
		},
	}

	return echoContext.JSON(http.StatusOK, resp)
}

// AddHolidays adds a date to current calendar holidays.
func (holidayHandler *CalendarHandler) AddHoliday(echoContext echo.Context) error {
	holidayReq := new(request.Holiday)

	if err := echoContext.Bind(holidayReq); err != nil {
		logrus.Errorf("add holiday handler bind: %v", err)
		return echo.ErrBadRequest
	}

	if err := holidayReq.Validate(); err != nil {
		logrus.Errorf("add holiday handler validate: %v", err)
		return echo.ErrBadRequest
	}

	holidayHandler.baseCalendar.AddHoliday(&cal.Holiday{
		Name:  holidayReq.Name,
		Type:  holidayHandler.observanceType(holidayReq.Type),
		Month: time.Month(holidayReq.Month),
		Day:   holidayReq.Day,
		Func:  cal.CalcDayOfMonth,
	})

	return echoContext.NoContent(http.StatusNoContent)
}

// UpdateHoliday updates a holiday of current calendar holidays.
func (holidayHandler *CalendarHandler) UpdateHoliday(echoContext echo.Context) error {
	return echoContext.JSON(http.StatusOK, "not implemented")
}

// RemoveHoliday removes a date from current calendar holidays.
func (holidayHandler *CalendarHandler) RemoveHoliday(echoContext echo.Context) error {
	return echoContext.JSON(http.StatusOK, "not implemented")
}

// IsHoliday checks if provided dates are holidays or not.
func (holidayHandler *CalendarHandler) IsHoliday(echoContext echo.Context) error {
	holidayReq := new(request.HolidayCheck)

	if err := echoContext.Bind(holidayReq); err != nil {
		logrus.Errorf("is holiday handler bind: %v", err)
		return echo.ErrBadRequest
	}

	if err := holidayReq.Validate(); err != nil {
		logrus.Errorf("is holiday handler validate: %v", err)
		return echo.ErrBadRequest
	}

	resp := make([]response.HolidayCheck, len(holidayReq.Dates))

	for index, date := range holidayReq.Dates {
		isHoliday, _, _ := holidayHandler.baseCalendar.IsHoliday(
			time.Date(time.Now().Year(), time.Month(date.Month), date.Day, 0, 0, 0, 0, ptime.Iran()),
		)

		resp[index] = response.HolidayCheck{
			Month:     date.Month,
			Day:       date.Day,
			IsHoliday: isHoliday,
		}
	}

	return echoContext.JSON(http.StatusOK, resp)
}

// observanceTypeDesc returns the string description of an observance type.
func (holidayHandler *CalendarHandler) observanceTypeDesc(observanceType cal.ObservanceType) string {
	switch observanceType {
	case cal.ObservancePublic:
		return "public"
	case cal.ObservanceReligious:
		return "religious"
	case cal.ObservanceBank:
		return "bank"
	case cal.ObservanceOther:
		return "other"
	case cal.ObservanceUnknown:
		return "unknown"
	default:
		return ""
	}
}

// observanceType returns the integer value of an observance type desc.
func (holidayHandler *CalendarHandler) observanceType(observanceTypeDesc string) cal.ObservanceType {
	switch observanceTypeDesc {
	case "public":
		return cal.ObservancePublic
	case "religious":
		return cal.ObservanceReligious
	case "bank":
		return cal.ObservanceBank
	case "other":
		return cal.ObservanceOther
	case "unknown":
		return cal.ObservanceUnknown
	default:
		return cal.ObservanceUnknown
	}
}
