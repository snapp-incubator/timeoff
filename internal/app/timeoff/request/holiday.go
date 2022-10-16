package request

import validation "github.com/go-ozzo/ozzo-validation"

const (
	minNameLen = 1
	maxNameLen = 128

	minMonthValue = 1
	maxMonthValue = 12

	minDayValue = 1
	maxDayValue = 12

	minHolidayCheckListLen = 1
	maxHolidayCheckListLen = 100

	publicObservanceType    = "public"
	religiousObservanceType = "religious"
	bankObservanceType      = "bank"
	otherObservanceType     = "other"
	unknownObservanceType   = "unknown"
)

// Holiday represents the struct of an HTTP holiday request.
type Holiday struct {
	Name  string `json:"name"`
	Type  string `json:"type"`
	Month int    `json:"month"`
	Day   int    `json:"day"`
}

// Validate validates a Holiday struct to be as expected.
func (h Holiday) Validate() error {
	return validation.ValidateStruct(&h,
		validation.Field(&h.Name, validation.Required, validation.Length(minNameLen, maxNameLen)),
		validation.Field(&h.Type, validation.Required, validation.In(
			publicObservanceType,
			religiousObservanceType,
			bankObservanceType,
			otherObservanceType,
			unknownObservanceType,
		)),
		validation.Field(&h.Month, validation.Required, validation.Min(minMonthValue), validation.Max(maxMonthValue)),
		validation.Field(&h.Day, validation.Required, validation.Min(minDayValue), validation.Max(maxDayValue)),
	)
}

// HolidayCheck represents the struct of an HTTP holiday check request.
type HolidayCheck struct {
	Dates []HolidayCheckItem `json:"dates"`
}

// Validate validates a HolidayCheck struct to be as expected.
func (hc HolidayCheck) Validate() error {
	return validation.ValidateStruct(&hc,
		validation.Field(&hc.Dates,
			validation.Required,
			validation.Length(minHolidayCheckListLen, maxHolidayCheckListLen),
		),
	)
}

// HolidayCheckItem represents the struct of an HTTP holiday check item.
type HolidayCheckItem struct {
	Month int `json:"month"`
	Day   int `json:"day"`
}

// Validate validates a HolidayCheckItem struct to be as expected.
func (hci HolidayCheckItem) Validate() error {
	return validation.ValidateStruct(&hci,
		validation.Field(&hci.Month, validation.Required, validation.Min(minMonthValue), validation.Max(maxMonthValue)),
		validation.Field(&hci.Day, validation.Required, validation.Min(minDayValue), validation.Max(maxDayValue)),
	)
}
