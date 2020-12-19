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
