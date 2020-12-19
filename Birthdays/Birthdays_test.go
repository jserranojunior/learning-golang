package birt

import (
	"testing"
	"time"
)

// Checks whether the date greater than the current month returns false
func TestNextMonthBirthdays(t *testing.T) {
	today := time.Now()
	nextMonth := Newbirthday(today.AddDate(0, 1, 0))
	test := nextMonth.Month()
	result := false
	if test != result {
		t.Error("Expected:", result, "Got:", test)
	}
}

// Checks whether the date less than the current month returns false
func TestLessMonthBirthdays(t *testing.T) {
	today := time.Now()
	nextMonth := Newbirthday(today.AddDate(0, -1, 0))
	test := nextMonth.Month()
	result := false
	if test != result {
		t.Error("Expected:", result, "Got:", test)
	}
}

// Checks whether the current date returns true
func TestCurrentMonthBirthdays(t *testing.T) {
	today := time.Now()
	currentMonth := Newbirthday(today)
	test := currentMonth.Month()
	result := true
	if test != result {
		t.Error("Expected:", result, "Got:", test)
	}
}

// Checks whether the day of the month return true
func TestFirstDayMonthBirthdays(t *testing.T) {
	today := time.Now()
	firstDayOfmonth := time.Date(today.Year(), today.Month(), 1, 0, 0, 0, 0, today.Location())
	nextMonth := Newbirthday(firstDayOfmonth)
	test := nextMonth.Month()
	result := true
	if test != result {
		t.Error("Expected:", result, "Got:", test)
	}
}

// Checks whether the second day of the month return true
func TestSecondDayMonthBirthdays(t *testing.T) {
	today := time.Now()
	secondDayOfmonth := time.Date(today.Year(), today.Month(), 2, 0, 0, 0, 0, today.Location())
	nextMonth := Newbirthday(secondDayOfmonth)
	test := nextMonth.Month()
	result := true
	if test != result {
		t.Error("Expected:", result, "Got:", test)
	}
}

// Checks whether year diferent of birth return true
func TestAnotherYearMonthBirthdays(t *testing.T) {
	today := time.Now()
	firstDayOfmonth := time.Date(2015, today.Month(), 2, 0, 0, 0, 0, today.Location())
	nextMonth := Newbirthday(firstDayOfmonth)
	test := nextMonth.Month()
	result := true
	if test != result {
		t.Error("Expected:", result, "Got:", test)
	}
}

// Checks whether last day of the week return week birthday true
func TestLastDayWeekBirthdays(t *testing.T) {
	today := time.Now()
	lastDayOfWeek := today

	for {
		if int(lastDayOfWeek.Weekday()) != 6 {
			lastDayOfWeek = lastDayOfWeek.AddDate(0, 0, 1)
		} else {
			break
		}
	}
	birth := Newbirthday(lastDayOfWeek)
	test := birth.Week()
	result := true
	if test != result {
		t.Error("Expected:", result, "Got:", test)
	}
}

// Checks whether first day of the week return week birthday true
func TestFirstDayWeekBirthdays(t *testing.T) {
	today := time.Now()
	firstDayOfWeek := today

	for {
		if int(firstDayOfWeek.Weekday()) != 0 {
			firstDayOfWeek = firstDayOfWeek.AddDate(0, 0, -1)
		} else {
			break
		}
	}
	birth := Newbirthday(firstDayOfWeek)
	test := birth.Week()
	result := true
	if test != result {
		t.Error("Expected:", result, "Got:", test)
	}
}

// Checks whether the birthday in the next week birthday false
func TestDayNextWeekBirthdays(t *testing.T) {
	today := time.Now()
	nextWeek := today

	for {
		if int(nextWeek.Weekday()) != 6 {
			nextWeek = nextWeek.AddDate(0, 0, 1)
		} else {
			break
		}
	}
	birth := Newbirthday(nextWeek.AddDate(0, 0, 1))
	test := birth.Week()
	result := false
	if test != result {
		t.Error("Expected:", result, "Got:", test)
	}
}

// Checks whether the birthday in the previous week birthday false
func TestDayPreviousWeekBirthdays(t *testing.T) {
	today := time.Now()
	previousWeek := today

	for {
		if int(previousWeek.Weekday()) != 0 {
			previousWeek = previousWeek.AddDate(0, 0, -1)
		} else {
			break
		}
	}
	birth := Newbirthday(previousWeek.AddDate(0, 0, -1))
	test := birth.Week()
	result := false
	if test != result {
		t.Error("Expected:", result, "Got:", test, "date:", previousWeek)
	}
}

// Checks tomorrow birth return day birthday false
func TestTomorrowBirthday(t *testing.T) {
	today := time.Now()
	tomorrow := today.AddDate(0, 0, 1)
	birth := Newbirthday(tomorrow)
	test := birth.Day()
	result := false
	if test != result {
		t.Error("Expected:", result, "Got:", test)
	}
}

// Checks Yesterday birth return day birthday false
func TestYesterdayBirthday(t *testing.T) {
	today := time.Now()
	yesterday := today.AddDate(0, 0, 1)
	birth := Newbirthday(yesterday)
	test := birth.Day()
	result := false
	if test != result {
		t.Error("Expected:", result, "Got:", test)
	}
}

// Checks today birth return day birthday false
func TestTodayBirthday(t *testing.T) {
	today := time.Now()
	birth := Newbirthday(today)
	test := birth.Day()
	result := true
	if test != result {
		t.Error("Expected:", result, "Got:", test)
	}
}
