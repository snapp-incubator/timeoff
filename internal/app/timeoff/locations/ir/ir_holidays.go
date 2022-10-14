// Package ir provides holiday definitions for the Iran.
package ir

import (
	"time"

	"github.com/rickar/cal/v2"
)

var (
	// NewYear represents New Year's Days.
	// 1-Farvardin to 5-Farvardin for national holidays and 6-Farvardin to 11-Farvardin for school holidays.
	NewYear = []*cal.Holiday{
		{
			Name:  "New Year's Day",
			Type:  cal.ObservancePublic,
			Month: time.March,
			Day:   21,
			Func:  cal.CalcDayOfMonth,
		},
		{
			Name:  "New Year's Day",
			Type:  cal.ObservancePublic,
			Month: time.March,
			Day:   22,
			Func:  cal.CalcDayOfMonth,
		},
		{
			Name:  "New Year's Day",
			Type:  cal.ObservancePublic,
			Month: time.March,
			Day:   23,
			Func:  cal.CalcDayOfMonth,
		},
		{
			Name:  "New Year's Day",
			Type:  cal.ObservancePublic,
			Month: time.March,
			Day:   24,
			Func:  cal.CalcDayOfMonth,
		},
		{
			Name:  "New Year's Day",
			Type:  cal.ObservancePublic,
			Month: time.March,
			Day:   25,
			Func:  cal.CalcDayOfMonth,
		},
		{
			Name:  "New Year's Day",
			Type:  cal.ObservanceOther,
			Month: time.March,
			Day:   26,
			Func:  cal.CalcDayOfMonth,
		},
		{
			Name:  "New Year's Day",
			Type:  cal.ObservanceOther,
			Month: time.March,
			Day:   27,
			Func:  cal.CalcDayOfMonth,
		},
		{
			Name:  "New Year's Day",
			Type:  cal.ObservanceOther,
			Month: time.March,
			Day:   28,
			Func:  cal.CalcDayOfMonth,
		},
		{
			Name:  "New Year's Day",
			Type:  cal.ObservanceOther,
			Month: time.March,
			Day:   29,
			Func:  cal.CalcDayOfMonth,
		},
		{
			Name:  "New Year's Day",
			Type:  cal.ObservanceOther,
			Month: time.March,
			Day:   30,
			Func:  cal.CalcDayOfMonth,
		},
		{
			Name:  "New Year's Day",
			Type:  cal.ObservanceOther,
			Month: time.March,
			Day:   31,
			Func:  cal.CalcDayOfMonth,
		},
	}

	// IslamicRepublicDay represents Islamic Republic Day holiday on 12-Farvardin.
	IslamicRepublicDay = &cal.Holiday{
		Name:  "Islamic Republic Day",
		Type:  cal.ObservancePublic,
		Month: time.April,
		Day:   1,
		Func:  cal.CalcWeekdayOffset,
	}

	// NatureDay represents Nature Day holiday on 13-Farvardin.
	NatureDay = &cal.Holiday{
		Name:  "Nature Day",
		Type:  cal.ObservancePublic,
		Month: time.April,
		Day:   2,
		Func:  cal.CalcWeekdayOffset,
	}

	// ImamKhomeyniDecease represents Imam Khomeyni Decease holiday on 14-Khordad.
	ImamKhomeyniDecease = &cal.Holiday{
		Name:  "Imam Khomeyni Decease",
		Type:  cal.ObservancePublic,
		Month: time.June,
		Day:   4,
		Func:  cal.CalcWeekdayOffset,
	}

	// KhordadFifteenthHonor represents Khordad Fifteenth Honor holiday on 15-Khordad.
	KhordadFifteenthHonor = &cal.Holiday{
		Name:  "Khordad Fifteenth Honor",
		Type:  cal.ObservancePublic,
		Month: time.June,
		Day:   5,
		Func:  cal.CalcWeekdayOffset,
	}

	// RevolutionDay represents Revolution Day holiday on 22-Bahman.
	RevolutionDay = &cal.Holiday{
		Name:  "Khordad Fifteenth Honor",
		Type:  cal.ObservancePublic,
		Month: time.February,
		Day:   11,
		Func:  cal.CalcWeekdayOffset,
	}

	// OilNationalDay represents Oil National Day holiday on 29-Esfand.
	OilNationalDay = &cal.Holiday{
		Name:  "Oil National Day",
		Type:  cal.ObservancePublic,
		Month: time.March,
		Day:   20,
		Func:  cal.CalcWeekdayOffset,
	}

	// Holidays provides a list of the standard national holidays.
	Holidays = append(
		NewYear,
		IslamicRepublicDay,
		NatureDay,
		ImamKhomeyniDecease,
		KhordadFifteenthHonor,
		RevolutionDay,
		OilNationalDay,
	)
)
