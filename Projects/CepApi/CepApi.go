package cepapi

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
)

// Cep interface
type Cep struct {
	Cep         string `json:"cep"`
	Logradouro  string `json:"logradouro"`
	Complemento string `json:"complemento"`
	Bairro      string `json:"bairro"`
	Localidade  string `json:"localidade"`
	Uf          string `json:"uf"`
	Ibge        string `json:"ibge"`
	Gia         string `json:"gia"`
	Ddd         string `json:"ddd"`
	Siafi       string `json:"siafi"`
}

//GetAddress return address with cep passed
func GetAddress(cep string) (Cep, error) {
	var Address Cep
	url := "http://viacep.com.br/ws/" + cep + "/json"
	response, err := http.Get(url)

	if err != nil {
		return Address, err
	}

	responseData, err := ioutil.ReadAll(response.Body)

	if err != nil {
		return Address, err
	}

	value := string(responseData)

	err = json.Unmarshal([]byte(value), &Address)

	if Address.Cep == "" {
		return Address, errors.New("There's not address with this cep")
	}

	return Address, nil
}
