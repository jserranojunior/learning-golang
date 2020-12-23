package cci

//MockMoney Value
type MockMoney struct {
	value float64
}

//GetDolar mock return API
func (mock MockMoney) GetDolar() float64 {
	return 15.0
}
