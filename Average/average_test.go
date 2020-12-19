package avg

import "testing"

func TestCalcAvg(t *testing.T ){
	test := CalcAvg(2,2,3)
	result := 2.3

	if test != result{
		t.Error("Expected:", result, "Got:", test)
	}
}

func TestCalcAvgFullValues(t *testing.T ){
	test := CalcAvg(2,2,3,4,5)
	result := 3.2

	if test != result{
		t.Error("Expected:", result, "Got:", test)
	}
}