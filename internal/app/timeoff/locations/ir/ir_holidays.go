// Package ir provides holiday definitions for the Iran.
package ir

import (
	"time"

	"github.com/hablullah/go-hijri"
	"github.com/rickar/cal/v2"
	"github.com/sirupsen/logrus"
	ptime "github.com/yaa110/go-persian-calendar"
)

func Holidays() []*cal.Holiday {
	return append(nationalHolidays(), religiousHolidays()...)
}

//nolint:funlen
func nationalHolidays() []*cal.Holiday {
	holidays := make([]*cal.Holiday, 0)
	currentPersianYear := ptime.New(time.Now()).Year()

	// Add New Year public holidays.
	for i := 1; i < 5; i++ {
		curDate := ptime.Date(currentPersianYear, ptime.Farvardin, i, 0, 0, 0, 0, ptime.Iran()).Time()

		holidays = append(holidays,
			&cal.Holiday{
				Name:  "New Year Public Holiday",
				Type:  cal.ObservancePublic,
				Month: curDate.Month(),
				Day:   curDate.Day(),
				Func:  cal.CalcDayOfMonth,
			},
		)
	}

	// Add New Year school holidays.
	for i := 5; i < 11; i++ {
		curDate := ptime.Date(currentPersianYear, ptime.Farvardin, i, 0, 0, 0, 0, ptime.Iran()).Time()

		holidays = append(holidays,
			&cal.Holiday{
				Name:  "New Year School Holiday",
				Type:  cal.ObservanceOther,
				Month: curDate.Month(),
				Day:   curDate.Day(),
				Func:  cal.CalcDayOfMonth,
			},
		)
	}

	islamicRepublicDay := ptime.Date(currentPersianYear, ptime.Farvardin, 12, 0, 0, 0, 0, ptime.Iran()).Time()
	natureDay := ptime.Date(currentPersianYear, ptime.Farvardin, 13, 0, 0, 0, 0, ptime.Iran()).Time()
	imamKhomeyniDecease := ptime.Date(currentPersianYear, ptime.Khordad, 14, 0, 0, 0, 0, ptime.Iran()).Time()
	khordadFifteenthHonor := ptime.Date(currentPersianYear, ptime.Khordad, 15, 0, 0, 0, 0, ptime.Iran()).Time()
	revolutionDay := ptime.Date(currentPersianYear, ptime.Bahman, 22, 0, 0, 0, 0, ptime.Iran()).Time()
	oilNationalDay := ptime.Date(currentPersianYear, ptime.Esfand, 29, 0, 0, 0, 0, ptime.Iran()).Time()

	holidays = append(holidays,
		&cal.Holiday{
			Name:  "Islamic Republic Day",
			Type:  cal.ObservancePublic,
			Month: islamicRepublicDay.Month(),
			Day:   islamicRepublicDay.Day(),
			Func:  cal.CalcDayOfMonth,
		},
		&cal.Holiday{
			Name:  "Nature Day",
			Type:  cal.ObservancePublic,
			Month: natureDay.Month(),
			Day:   natureDay.Day(),
			Func:  cal.CalcDayOfMonth,
		},
		&cal.Holiday{
			Name:  "Imam Khomeyni Decease",
			Type:  cal.ObservancePublic,
			Month: imamKhomeyniDecease.Month(),
			Day:   imamKhomeyniDecease.Day(),
			Func:  cal.CalcDayOfMonth,
		},
		&cal.Holiday{
			Name:  "Khordad Fifteenth Honor",
			Type:  cal.ObservancePublic,
			Month: khordadFifteenthHonor.Month(),
			Day:   khordadFifteenthHonor.Day(),
			Func:  cal.CalcDayOfMonth,
		},
		&cal.Holiday{
			Name:  "Revolution Day",
			Type:  cal.ObservancePublic,
			Month: revolutionDay.Month(),
			Day:   revolutionDay.Day(),
			Func:  cal.CalcDayOfMonth,
		},
		&cal.Holiday{
			Name:  "Oil National Day",
			Type:  cal.ObservancePublic,
			Month: oilNationalDay.Month(),
			Day:   oilNationalDay.Day(),
			Func:  cal.CalcDayOfMonth,
		},
	)

	return holidays
}

//nolint:funlen
func religiousHolidays() []*cal.Holiday {
	var currentHijriYear int64

	holidays := make([]*cal.Holiday, 0)
	currentDate := time.Now()

	if hijriNow, err := hijri.CreateHijriDate(currentDate, hijri.Default); err == nil {
		currentHijriYear = hijriNow.Year
	} else {
		logrus.Infof("failed to calculate religious holidays: failed to get current hijri year: %v", err)
		return holidays
	}

	tasua := hijri.HijriDate{Year: currentHijriYear, Month: 1, Day: 9}.ToGregorian()
	ashura := hijri.HijriDate{Year: currentHijriYear, Month: 1, Day: 10}.ToGregorian()
	arbaeen := hijri.HijriDate{Year: currentHijriYear, Month: 2, Day: 20}.ToGregorian()
	prophetMartyrdom := hijri.HijriDate{Year: currentHijriYear, Month: 2, Day: 28}.ToGregorian()
	imamRezaMartyrdom := hijri.HijriDate{Year: currentHijriYear, Month: 2, Day: 30}.ToGregorian()
	imamAskariMartyrdom := hijri.HijriDate{Year: currentHijriYear, Month: 3, Day: 8}.ToGregorian()
	prophetBirthday := hijri.HijriDate{Year: currentHijriYear, Month: 3, Day: 17}.ToGregorian()
	holinessZahraMartyrdom := hijri.HijriDate{Year: currentHijriYear, Month: 6, Day: 3}.ToGregorian()
	imamAliBirthday := hijri.HijriDate{Year: currentHijriYear, Month: 7, Day: 13}.ToGregorian()
	eidMubarath := hijri.HijriDate{Year: currentHijriYear, Month: 7, Day: 27}.ToGregorian()
	imamMahdiBirthday := hijri.HijriDate{Year: currentHijriYear, Month: 8, Day: 15}.ToGregorian()
	imamAliMartyrdom := hijri.HijriDate{Year: currentHijriYear, Month: 9, Day: 21}.ToGregorian()
	eidAlFitr := hijri.HijriDate{Year: currentHijriYear, Month: 10, Day: 1}.ToGregorian()
	eidAlFitrNext := hijri.HijriDate{Year: currentHijriYear, Month: 10, Day: 2}.ToGregorian()
	imamSadeghMartyrdom := hijri.HijriDate{Year: currentHijriYear, Month: 10, Day: 25}.ToGregorian()
	eidAlAdha := hijri.HijriDate{Year: currentHijriYear, Month: 12, Day: 10}.ToGregorian()
	eidGhadir := hijri.HijriDate{Year: currentHijriYear, Month: 12, Day: 18}.ToGregorian()

	holidays = append(holidays,
		&cal.Holiday{
			Name:  "Tasua",
			Type:  cal.ObservanceReligious,
			Month: tasua.Month(),
			Day:   tasua.Day(),
			Func:  cal.CalcDayOfMonth,
		},
		&cal.Holiday{
			Name:  "Ashura",
			Type:  cal.ObservanceReligious,
			Month: ashura.Month(),
			Day:   ashura.Day(),
			Func:  cal.CalcDayOfMonth,
		},
		&cal.Holiday{
			Name:  "Arbaeen",
			Type:  cal.ObservanceReligious,
			Month: arbaeen.Month(),
			Day:   arbaeen.Day(),
			Func:  cal.CalcDayOfMonth,
		},
		&cal.Holiday{
			Name:  "Prophet Martyrdom",
			Type:  cal.ObservanceReligious,
			Month: prophetMartyrdom.Month(),
			Day:   prophetMartyrdom.Day(),
			Func:  cal.CalcDayOfMonth,
		},
		&cal.Holiday{
			Name:  "Imam Reza Martyrdom",
			Type:  cal.ObservanceReligious,
			Month: imamRezaMartyrdom.Month(),
			Day:   imamRezaMartyrdom.Day(),
			Func:  cal.CalcDayOfMonth,
		},
		&cal.Holiday{
			Name:  "Imam Askari Martyrdom",
			Type:  cal.ObservanceReligious,
			Month: imamAskariMartyrdom.Month(),
			Day:   imamAskariMartyrdom.Day(),
			Func:  cal.CalcDayOfMonth,
		},
		&cal.Holiday{
			Name:  "Prophet Birthday",
			Type:  cal.ObservanceReligious,
			Month: prophetBirthday.Month(),
			Day:   prophetBirthday.Day(),
			Func:  cal.CalcDayOfMonth,
		},
		&cal.Holiday{
			Name:  "Holiness Zahra Martyrdom",
			Type:  cal.ObservanceReligious,
			Month: holinessZahraMartyrdom.Month(),
			Day:   holinessZahraMartyrdom.Day(),
			Func:  cal.CalcDayOfMonth,
		},
		&cal.Holiday{
			Name:  "Imam Ali Birthday",
			Type:  cal.ObservanceReligious,
			Month: imamAliBirthday.Month(),
			Day:   imamAliBirthday.Day(),
			Func:  cal.CalcDayOfMonth,
		},
		&cal.Holiday{
			Name:  "Eid Mubarath",
			Type:  cal.ObservanceReligious,
			Month: eidMubarath.Month(),
			Day:   eidMubarath.Day(),
			Func:  cal.CalcDayOfMonth,
		},
		&cal.Holiday{
			Name:  "Imam Mahdi Birthday",
			Type:  cal.ObservanceReligious,
			Month: imamMahdiBirthday.Month(),
			Day:   imamMahdiBirthday.Day(),
			Func:  cal.CalcDayOfMonth,
		},
		&cal.Holiday{
			Name:  "Imam Ali Martyrdom",
			Type:  cal.ObservanceReligious,
			Month: imamAliMartyrdom.Month(),
			Day:   imamAliMartyrdom.Day(),
			Func:  cal.CalcDayOfMonth,
		},
		&cal.Holiday{
			Name:  "Eid Al-Fitr",
			Type:  cal.ObservanceReligious,
			Month: eidAlFitr.Month(),
			Day:   eidAlFitr.Day(),
			Func:  cal.CalcDayOfMonth,
		},
		&cal.Holiday{
			Name:  "Eid Al-Fitr Next",
			Type:  cal.ObservanceReligious,
			Month: eidAlFitrNext.Month(),
			Day:   eidAlFitrNext.Day(),
			Func:  cal.CalcDayOfMonth,
		},
		&cal.Holiday{
			Name:  "Imam Sadegh Martyrdom",
			Type:  cal.ObservanceReligious,
			Month: imamSadeghMartyrdom.Month(),
			Day:   imamSadeghMartyrdom.Day(),
			Func:  cal.CalcDayOfMonth,
		},
		&cal.Holiday{
			Name:  "Eid Al-Adha",
			Type:  cal.ObservanceReligious,
			Month: eidAlAdha.Month(),
			Day:   eidAlAdha.Day(),
			Func:  cal.CalcDayOfMonth,
		},
		&cal.Holiday{
			Name:  "Eid Ghadir",
			Type:  cal.ObservanceReligious,
			Month: eidGhadir.Month(),
			Day:   eidGhadir.Day(),
			Func:  cal.CalcDayOfMonth,
		},
	)

	return holidays
}
