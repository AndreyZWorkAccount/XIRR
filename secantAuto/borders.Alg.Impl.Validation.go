package secantAuto

import . "math"

type ValidationOptions struct {
	useBorders      bool

	validateBordersMinMax bool
	minLeft               float64
	maxRight              float64
}


func getBordersSearchParams(paymentsSumIsPositive bool, useValidation bool ) (left float64, right float64, options ValidationOptions){
	if paymentsSumIsPositive {
		if useValidation{
			left, right, options = bordersWithValidation(PositiveSum_Initial_Guess, PositiveSum_MinLeft, PositiveSum_MaxRight)
			return
		}
		left, right, options = bordersWithoutValidation(PositiveSum_LeftBoundary, PositiveSum_RightBoundary)
		return
	}

	if useValidation{
		left, right, options = bordersWithValidation(NegativeOrZeroSum_Initial_Guess, NegativeOrZeroSum_MinLeft, NegativeOrZeroSum_MaxRight)
		return
	}
	left, right, options = bordersWithoutValidation(NegativeOrZeroSum_LeftBoundary, NegativeOrZeroSum_RightBoundary)
	return
}

func bordersWithValidation(guess float64, minLeft float64, maxRight float64)(left float64, right float64, options ValidationOptions){
	left = guess - Delta_For_Borders
	right = guess + Delta_For_Borders
	options = useValidationOptions(minLeft, maxRight)
	return
}

func bordersWithoutValidation(leftBoundary float64, rightBoundary float64) (left float64, right float64, options ValidationOptions){
	left = leftBoundary
	right = rightBoundary
	options = noValidationOptions()
	return
}

func noValidationOptions() ValidationOptions {
	return ValidationOptions{validateBordersMinMax: false, useBorders:true, minLeft:Inf(-1), maxRight:Inf(1)}
}

func useValidationOptions(minLeft float64, maxRight float64) ValidationOptions  {
	return ValidationOptions{validateBordersMinMax: true, useBorders:false, minLeft:minLeft, maxRight:maxRight}
}
