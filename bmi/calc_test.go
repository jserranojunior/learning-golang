package bmi

import "testing"

func TestCalc(t *testing.T){
	test := Calc(170,70)
	result := 24.22

	if test != result{
		t.Error("Expected:", result, "got:", test)
	}
}