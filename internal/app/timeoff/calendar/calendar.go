package calendar

import (
	"fmt"
	"time"

	"github.com/hablullah/go-hijri"
	"github.com/rickar/cal/v2"
	"github.com/robfig/cron"
	"github.com/sirupsen/logrus"
	"github.com/snapp-incubator/timeoff/internal/app/timeoff/locations/ir"
	ptime "github.com/yaa110/go-persian-calendar"
)

type HolidayEngine struct {
	baseCalendar *cal.BusinessCalendar
}

// NewHolidayEngine returns a new instance of HolidayEngine.
func NewHolidayEngine(baseCal *cal.BusinessCalendar) *HolidayEngine {
	return &HolidayEngine{
		baseCalendar: baseCal,
	}
}

func (holidayEngine *HolidayEngine) updateHijriHolidays() {
	var todayHijriYear, yesterdayHijriYear int64

	todayDate := time.Now()

	if hijriDate, err := hijri.CreateHijriDate(todayDate, hijri.Default); err == nil {
		todayHijriYear = hijriDate.Year
	} else {
		logrus.Infof("failed to get today's hijri year. using zero: %v", err)
	}

	yesterdayDate := time.Now()

	if hijriDate, err := hijri.CreateHijriDate(yesterdayDate, hijri.Default); err == nil {
		yesterdayHijriYear = hijriDate.Year
	} else {
		logrus.Infof("failed to get yesterday's hijri year. using zero: %v", err)
	}

	if todayHijriYear == yesterdayHijriYear {
		logrus.Infof("still the same hijri year: %d", todayHijriYear)
		return
	}

	for _, holiday := range holidayEngine.baseCalendar.Holidays {
		if holiday.Type == cal.ObservanceReligious {
			curHoliday := time.Date(todayDate.Year(), holiday.Month, holiday.Day, 0, 0, 0, 0, ptime.Iran())

			curHoliday = curHoliday.Add(-10 * 24 * time.Hour)
			holiday.Day = curHoliday.Day()
			holiday.Month = curHoliday.Month()
		}
	}
}

func (holidayEngine *HolidayEngine) updatePersianHolidays() {
	todayPersianDate := ptime.New(time.Now())
	yesterdayPersianDate := ptime.New(time.Now().Add(-24 * time.Hour))

	if todayPersianDate.Year() != yesterdayPersianDate.Year() {
		var leapDiff int

		if yesterdayPersianDate.IsLeap() {
			leapDiff = 1
		} else if todayPersianDate.IsLeap() {
			leapDiff = -1
		}

		for _, holiday := range holidayEngine.baseCalendar.Holidays {
			if holiday.Type != cal.ObservanceReligious {
				newHoliday := time.Date(time.Now().Year(), holiday.Month, holiday.Day,
					0, 0, 0, 0, ptime.Iran()).Add(time.Duration(leapDiff*24) * time.Hour)

				holiday.Day = newHoliday.Day()
				holiday.Month = newHoliday.Month()
			}
		}
	}
}

func (holidayEngine *HolidayEngine) isGeorgianLeapYear(year int) bool {
	return year%4 == 0 && year%100 != 0 || year%400 == 0
}

func (holidayEngine *HolidayEngine) Start() error {
	cronJob := cron.New()
	if err := cronJob.AddFunc("0 0 * * *", func() {
		holidayEngine.updateHijriHolidays()
		holidayEngine.updatePersianHolidays()
	}); err != nil {
		return fmt.Errorf("failed to start hijri engine: %w", err)
	}

	cronJob.Start()

	return nil
}

// SetupBaseCalendar creates and returns an iranian calendar with its default holidays.
func SetupBaseCalendar() *cal.BusinessCalendar {
	businessCal := cal.NewBusinessCalendar()

	businessCal.SetWorkday(time.Saturday, true)
	businessCal.SetWorkday(time.Sunday, true)
	businessCal.SetWorkday(time.Monday, true)
	businessCal.SetWorkday(time.Tuesday, true)
	businessCal.SetWorkday(time.Wednesday, true)
	businessCal.SetWorkday(time.Thursday, true)
	businessCal.SetWorkday(time.Friday, false)

	businessCal.AddHoliday(ir.Holidays()...)

	return businessCal
}
