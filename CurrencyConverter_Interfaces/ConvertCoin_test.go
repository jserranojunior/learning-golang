package cci

import "testing"

func TestMockRealInDolar(t *testing.T) {

	money := MockMoney{
		value: 15.0,
	}

	test := RealToDolar(money)
	result := 15.0
	if test != result {
		t.Error("Expected:", result, "Got:", test)
	}
}

func TestRealInDolar(t *testing.T) {
	money := MoneyContructor()
	test := RealToDolar(money)
	result := 17.0
	if test != result {
		t.Error("Expected:", result, "Got:", test)
	}
}
