package response

type (
	// Holidays represents the struct of all holidays for a location in HTTP response.
	Holidays struct {
		Location string    `json:"location"`
		Holidays []Holiday `json:"holidays"`
	}

	// Holiday represents the struct of each holiday in HTTP response.
	Holiday struct {
		Name  string `json:"name"`
		Type  string `json:"type"`
		Month int    `json:"month"`
		Day   int    `json:"day"`
	}

	// HolidayCheck represents the struct of each holiday check result in HTTP response.
	HolidayCheck struct {
		Month     int  `json:"month"`
		Day       int  `json:"day"`
		IsHoliday bool `json:"is_holiday"`
	}
)
