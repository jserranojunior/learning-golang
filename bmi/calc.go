package bmi
import (
		"math"
	)
func Calc(height, weight float64)float64{
	var doubleHeight float64 = height * height
	var bodyMassIndex float64 = (weight / doubleHeight) * 10000
	formaterBodyMassIndex := math.Round(bodyMassIndex*100)/100
	return formaterBodyMassIndex
}