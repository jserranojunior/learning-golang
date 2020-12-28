package cci

import (
	"fmt"
)

//Money Value
type Money struct {
	Value float64
}

//MoneyMethods interface
type MoneyMethods interface {
	GetDolar() (float64, error)
	GetValue() float64
}

//GetValue return value default
func (money Money) GetValue() float64 {
	return money.Value
}

//RealToDolar converter dolar in real
func RealToDolar(m MoneyMethods) (float64, error) {
	dolar, err := m.GetDolar()

	if err != nil {
		fmt.Println(err)
		return float64(dolar), err
	}
	value := m.GetValue()
	RealInDolar := value * float64(dolar)
	return RealInDolar, nil
}

//MoneyContructor define values default
func MoneyContructor() Money {
	return Money{
		Value: 18.0,
	}
}
