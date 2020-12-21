package cepapi

import (
	"testing"
)

func TestGetAddress(t *testing.T) {
	test, err := GetAddress("01001-000")
	var result = Cep{
		Cep:         "01001-000",
		Logradouro:  "Praça da Sé",
		Complemento: "lado ímpar",
		Bairro:      "Sé",
		Localidade:  "São Paulo",
		Uf:          "SP",
		Ibge:        "3550308",
		Gia:         "1004",
		Ddd:         "11",
		Siafi:       "7107",
	}

	if test != result || err != nil {
		t.Error("Expected:", result.Cep, "Got:", test.Cep, err)
	}
}

func TestGetOutherAddress(t *testing.T) {
	test, err := GetAddress("08062190")
	var result = Cep{
		Cep:         "08062-190",
		Logradouro:  "Rua Francisco Monteiro",
		Complemento: "",
		Bairro:      "Vila Norma",
		Localidade:  "São Paulo",
		Uf:          "SP",
		Ibge:        "3550308",
		Gia:         "1004",
		Ddd:         "11",
		Siafi:       "7107",
	}

	if test != result || err != nil {
		t.Error("Expected:", result.Cep, "Got:", test.Cep, err)
	}
}

func TestGetNilAddress(t *testing.T) {
	test, err := GetAddress("00000000")
	var result = Cep{}

	if test != result && err == nil {
		t.Error("Expected:", result.Cep, "Got:", test.Cep, err)
	}

}
