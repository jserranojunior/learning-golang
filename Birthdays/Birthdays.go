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
