package numMethods

import (
	"math"
	)

type SecantMethod struct {
	XLeftInit, XRightInit float64
}

// NumericMethod interface implementation
func (s *SecantMethod) Calculate(F NumericFunc, methodParams *NumericMethodParams) (float64, NumericResultType, *NumericMethodError){
	xLeft := s.XLeftInit
	xRight := s.XRightInit

	var iterationPassed  uint64 = 0
	for iterationPassed < methodParams.MaxIterationsCount {

		//check if we reach necessary precision
		dx := math.Abs(xRight - xLeft)
		if dx < methodParams.Epsilon{
			return average(xLeft,xRight), NumericResultType_HasSolution, nil
		}

		xRightOld := xRight

		fxRight := F(xRight)
		fxLeft := F(xLeft)

		if isNanOrInfinity(fxLeft,fxRight){
			return xRight, NumericResultType_NoSolution, error(FunctionValueIsNanOrInfinity)
		}

		deltaF := fxRight - fxLeft
		if deltaF == 0 {
			return xRight, NumericResultType_NoSolution, error(FunctionsDeltaIsZero)
		}

		xRight = (xLeft*F(xRight) - xRight*F(xLeft))/deltaF
		xLeft = xRightOld

		iterationPassed++
	}

	return xRight, NumericResultType_NoSolution, nil
}

func error(description string) *NumericMethodError{
	return &NumericMethodError{"Secant", description}
}

func average(a float64,b float64) float64{
	return a + (b - a)/2
}

func isNanOrInfinity(numbers ...float64) bool{
	for _,num := range numbers{
		if math.IsInf(num,0) || math.IsNaN(num){
			return true
		}
	}
	return false
}

