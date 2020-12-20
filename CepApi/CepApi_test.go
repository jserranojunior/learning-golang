package cepapi

import (
	"encoding/json"
	"testing"
)

func TestGetAddress(t *testing.T) {
	test := GetAddress("01001-000")
	var Address Cep
	value := `{
		"cep": "01001-000",
		"logradouro": "Praça da Sé",
		"complemento": "lado ímpar",
		"bairro": "Sé",
		"localidade": "São Paulo",
		"uf": "SP",
		"ibge": "3550308",
		"gia": "1004",
		"ddd": "11",
		"siafi": "7107"
	  }`
	json.Unmarshal([]byte(value), &Address)
	result := Address

	if test != result {
		t.Error("Expected:", result, "Got:", test)
	}
}
