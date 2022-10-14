package calendar

import (
	"fmt"
	"time"

	"github.com/hablullah/go-hijri"
	"github.com/rickar/cal/v2"
	"github.com/robfig/cron"
	"github.com/sirupsen/logrus"
	"github.com/zeinababbasi/timeoff/internal/app/timeoff/locations/ir"
)

type HijriEngine struct {
	baseCalendar *cal.BusinessCalendar
	latestYear   int64
}

// NewHijriEngine returns a new instance of HijriEngine.
func NewHijriEngine(baseCal *cal.BusinessCalendar) *HijriEngine {
	return &HijriEngine{
		baseCalendar: baseCal,
	}
}

func (hijriEngine *HijriEngine) updateHijriHolidays() {
	var currentHijriYear int64

	currentDate := time.Now()

	if hijriNow, err := hijri.CreateHijriDate(currentDate, hijri.Default); err == nil {
		currentHijriYear = hijriNow.Year
	} else {
		logrus.Infof("failed to get current hijri year. using zero: %v", err)
	}

	if currentHijriYear == hijriEngine.latestYear {
		logrus.Infof("still the same hijri year: %d", currentHijriYear)
		return
	}

	for _, holiday := range hijriEngine.baseCalendar.Holidays {
		if holiday.Type == cal.ObservanceReligious {
			curHoliday := time.Date(currentDate.Year(), holiday.Month, holiday.Day, 0, 0, 0, 0, time.Local)

			curHoliday = curHoliday.Add(-10 * 24 * time.Hour)
			holiday.Day = curHoliday.Day()
			holiday.Month = curHoliday.Month()
		}
	}

}

//nolint:funlen
func (hijriEngine *HijriEngine) AddReligiousHolidays() {
	var currentHijriYear int64

	currentDate := time.Now()

	if hijriNow, err := hijri.CreateHijriDate(currentDate, hijri.Default); err == nil {
		currentHijriYear = hijriNow.Year
	} else {
		logrus.Infof("failed to add religious holidays: failed to get current hijri year: %v", err)
		return
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

	hijriEngine.baseCalendar.AddHoliday(
		&cal.Holiday{
			Name:  "Tasua",
			Type:  cal.ObservanceReligious,
			Month: tasua.Month(),
			Day:   tasua.Year(),
			Func:  cal.CalcWeekdayOffset,
		},
		&cal.Holiday{
			Name:  "Ashura",
			Type:  cal.ObservanceReligious,
			Month: ashura.Month(),
			Day:   ashura.Year(),
			Func:  cal.CalcWeekdayOffset,
		},
		&cal.Holiday{
			Name:  "Arbaeen",
			Type:  cal.ObservanceReligious,
			Month: arbaeen.Month(),
			Day:   arbaeen.Year(),
			Func:  cal.CalcWeekdayOffset,
		},
		&cal.Holiday{
			Name:  "Prophet Martyrdom",
			Type:  cal.ObservanceReligious,
			Month: prophetMartyrdom.Month(),
			Day:   prophetMartyrdom.Year(),
			Func:  cal.CalcWeekdayOffset,
		},
		&cal.Holiday{
			Name:  "Imam Reza Martyrdom",
			Type:  cal.ObservanceReligious,
			Month: imamRezaMartyrdom.Month(),
			Day:   imamRezaMartyrdom.Year(),
			Func:  cal.CalcWeekdayOffset,
		},
		&cal.Holiday{
			Name:  "Imam Askari Martyrdom",
			Type:  cal.ObservanceReligious,
			Month: imamAskariMartyrdom.Month(),
			Day:   imamAskariMartyrdom.Year(),
			Func:  cal.CalcWeekdayOffset,
		},
		&cal.Holiday{
			Name:  "Prophet Birthday",
			Type:  cal.ObservanceReligious,
			Month: prophetBirthday.Month(),
			Day:   prophetBirthday.Year(),
			Func:  cal.CalcWeekdayOffset,
		},
		&cal.Holiday{
			Name:  "Holiness Zahra Martyrdom",
			Type:  cal.ObservanceReligious,
			Month: holinessZahraMartyrdom.Month(),
			Day:   holinessZahraMartyrdom.Year(),
			Func:  cal.CalcWeekdayOffset,
		},
		&cal.Holiday{
			Name:  "Imam Ali Birthday",
			Type:  cal.ObservanceReligious,
			Month: imamAliBirthday.Month(),
			Day:   imamAliBirthday.Year(),
			Func:  cal.CalcWeekdayOffset,
		},
		&cal.Holiday{
			Name:  "Eid Mubarath",
			Type:  cal.ObservanceReligious,
			Month: eidMubarath.Month(),
			Day:   eidMubarath.Year(),
			Func:  cal.CalcWeekdayOffset,
		},
		&cal.Holiday{
			Name:  "Imam Mahdi Birthday",
			Type:  cal.ObservanceReligious,
			Month: imamMahdiBirthday.Month(),
			Day:   imamMahdiBirthday.Year(),
			Func:  cal.CalcWeekdayOffset,
		},
		&cal.Holiday{
			Name:  "Imam Ali Martyrdom",
			Type:  cal.ObservanceReligious,
			Month: imamAliMartyrdom.Month(),
			Day:   imamAliMartyrdom.Year(),
			Func:  cal.CalcWeekdayOffset,
		},
		&cal.Holiday{
			Name:  "Eid Al-Fitr",
			Type:  cal.ObservanceReligious,
			Month: eidAlFitr.Month(),
			Day:   eidAlFitr.Year(),
			Func:  cal.CalcWeekdayOffset,
		},
		&cal.Holiday{
			Name:  "Eid Al-Fitr Next",
			Type:  cal.ObservanceReligious,
			Month: eidAlFitrNext.Month(),
			Day:   eidAlFitrNext.Year(),
			Func:  cal.CalcWeekdayOffset,
		},
		&cal.Holiday{
			Name:  "Imam Sadegh Martyrdom",
			Type:  cal.ObservanceReligious,
			Month: imamSadeghMartyrdom.Month(),
			Day:   imamSadeghMartyrdom.Year(),
			Func:  cal.CalcWeekdayOffset,
		},
		&cal.Holiday{
			Name:  "Eid Al-Adha",
			Type:  cal.ObservanceReligious,
			Month: eidAlAdha.Month(),
			Day:   eidAlAdha.Year(),
			Func:  cal.CalcWeekdayOffset,
		},
		&cal.Holiday{
			Name:  "Eid Ghadir",
			Type:  cal.ObservanceReligious,
			Month: eidGhadir.Month(),
			Day:   eidGhadir.Year(),
			Func:  cal.CalcWeekdayOffset,
		},
	)
}

func (hijriEngine *HijriEngine) Start() error {
	cronJob := cron.New()
	if err := cronJob.AddFunc("0 0 * * *", func() {
		hijriEngine.updateHijriHolidays()
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

	businessCal.AddHoliday(ir.Holidays...)

	return businessCal
}
