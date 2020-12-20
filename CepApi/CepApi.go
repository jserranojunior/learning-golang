package cepapi

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// Cep interface
type Cep struct {
	cep         string
	logradouro  string
	complemento string
	bairro      string
	localidade  string
	uf          string
	ibge        string
	gia         string
	ddd         string
	siafi       string
}

//GetAddress return address with cep passed
func GetAddress(cep string) Cep {
	response, err := http.Get("http://viacep.com.br/ws/01001000/json")

	if err != nil {
		fmt.Println(err.Error())
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println(err.Error())
	}

	value := string(responseData)
	var Address Cep
	json.Unmarshal([]byte(value), &Address)
	return Address
}
