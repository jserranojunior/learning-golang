package cci

//MockMoney Value
type MockMoney struct {
	value float64
}

//GetDolar mock return API
func (mock MockMoney) GetDolar() (float64, error) {

	return 15.0, nil
}
