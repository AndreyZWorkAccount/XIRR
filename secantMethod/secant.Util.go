package secantMethod

import . "math"

func average(a float64, b float64) float64{
	return a + (b - a)/2
}

func AnyNanOrInfinity(numbers ...float64) bool{
	for _,num := range numbers{
		if IsInf(num,0) || IsNaN(num){
			return true
		}
	}
	return false
}
