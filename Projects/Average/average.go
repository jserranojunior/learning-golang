package avg
import (
"math")

func CalcAvg(nums ...float64) float64{
	var total float64 = 0
	for _, num := range nums{
		total += float64(num)
	}
	var formaterAverage float64 = (total / float64(len(nums)))
	formaterAverage = math.Round(formaterAverage*10)/10

	return formaterAverage

}

