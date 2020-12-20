package cvdt

import (
	"strings"
	"time"
)

//ConvertDateToUs return converted date in US
func ConvertDateToUs(date string) string {
	var layout string
	var dateFormated string

	if strings.Contains(date, "/") && !strings.Contains(date, "m=") && !strings.Contains(date, "UTC") {
		layout = "01/02/2006T15:04:05.000Z"
		if !strings.Contains(date, "T") {
			date = date + "T11:45:26.371Z"
		}
		NewDate, _ := time.Parse(layout, date)
		dateFormated = NewDate.Format("2006-02-01")

	} else if strings.Contains(date, "-") && !strings.Contains(date, "m=") && !strings.Contains(date, "UTC") {
		layout = "2006-02-01T15:04:05.000Z"
		if !strings.Contains(date, "T") {
			date = date + "T11:45:26.371Z"
		}
		NewDate, _ := time.Parse(layout, date)
		dateFormated = NewDate.Format("2006-02-01")

	} else if strings.Contains(date, " ") {
		layout = "2006-01-02"
		newRune := []rune(date)
		sliceDate := string(newRune[0:10])
		dateFormated = sliceDate
	}
	return dateFormated
}
