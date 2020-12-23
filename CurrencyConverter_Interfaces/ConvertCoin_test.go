package cci

import "testing"

func TestMockRealInDolar(t *testing.T) {

	money := MockMoney{
		value: 15.0,
	}

	test, err := RealToDolar(money)
	result := 15.0
	if test != result || err != nil {
		t.Error("Expected:", result, "Got:", test)
	}
}

func TestRealInDolar(t *testing.T) {
	money := MoneyContructor()
	test, err := RealToDolar(money)
	result := 17.0
	if test != result || err != nil {
		t.Error("Expected:", result, "Got:", test)
	}
}
