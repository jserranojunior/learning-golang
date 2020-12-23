package cci

//MockMoney Value
type MockMoney struct {
	Value float64
}

//GetDolar mock return API
func (mock MockMoney) GetDolar() (float64, error) {
	return float64(5), nil
}

//GetValue return value default
func (mock MockMoney) GetValue() float64 {
	return mock.Value
}
