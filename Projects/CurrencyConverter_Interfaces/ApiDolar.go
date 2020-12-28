package cci

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type apiFinance struct {
	By            string  `json:"by"`
	ValidKey      bool    `json:"valid_key"`
	Results       Results `json:"results"`
	ExecutionTime float64 `json:"execution_time"`
	FromCache     bool    `json:"from_cache"`
}

//Results of apiDolar
type Results struct {
	Currencies Currencies `json:"currencies"`
}

//Currencies of ApiDolar
type Currencies struct {
	USD Currence `json:"USD"`
	EUR Currence `json:"EUR"`
}

//Currence of apiDolar
type Currence struct {
	Name string  `json:"name"`
	Buy  float64 `json:"buy"`
}

//GetDolar return api Dolar valeu
func (m Money) GetDolar() (float64, error) {
	var apiFinance apiFinance

	url := "https://api.hgbrasil.com/finance"
	response, err := http.Get(url)
	responseData, err := ioutil.ReadAll(response.Body)

	value := string(responseData)

	err = json.Unmarshal([]byte(value), &apiFinance)

	if err != nil {
		return 1, err
	}
	return apiFinance.Results.Currencies.USD.Buy, err
}
