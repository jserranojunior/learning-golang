package cci

//Money Value
type Money struct {
	value float64
}

//MoneyMethods interface
type MoneyMethods interface {
	GetDolar() float64
}

//GetDolar return api Dolar valeu
func (m Money) GetDolar() float64 {
	return m.value
}

//RealToDolar converter dolar in real
func RealToDolar(m MoneyMethods) float64 {
	return float64(m.GetDolar())
}

//MoneyContructor define values default
func MoneyContructor() Money {
	return Money{
		value: 18.0,
	}
}
