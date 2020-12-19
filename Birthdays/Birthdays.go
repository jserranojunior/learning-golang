package birt

import "time"

// Birthday is ...
type Birthday struct {
	date time.Time
}

// Newbirthday Implements Birthday ...
func Newbirthday(date time.Time) *Birthday {
	birth := Birthday{date: date}
	return &birth
}

// Month verify month of birth in Birthday
func (birth Birthday) Month() bool {
	today := time.Now()
	currentMonth := today.Month()
	monthOfBirth := birth.date.Month()
	if currentMonth != monthOfBirth {
		return false
	}
	return true
}

// Day verify day of birth in Birthday
func (birth Birthday) Day() bool {
	today := time.Now()
	if today.Day() != birth.date.Day() {
		return false
	}
	return true
}

// Week verify month and week of birth in Birthday
func (birth Birthday) Week() bool {
	if birth.Month() {
		today := time.Now()
		firstDayOfWeek := today

		for {
			if int(firstDayOfWeek.Weekday()) > 0 {
				firstDayOfWeek = firstDayOfWeek.AddDate(0, 0, -1)
			} else {
				break
			}
		}
		if birth.date.Day() == firstDayOfWeek.Day() {
			return true
		}
		for {
			if int(firstDayOfWeek.Weekday()) < 6 {
				firstDayOfWeek = firstDayOfWeek.AddDate(0, 0, 1)
				if birth.date.Day() == firstDayOfWeek.Day() {
					return true
				}
			} else {
				return false
			}
		}
	}
	return false
}
