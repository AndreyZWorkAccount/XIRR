package secantMethod

import (
	. "math"
	. "XIRR/numMethods"
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
		dx := Abs(xRight - xLeft)
		if dx < methodParams.Epsilon{
			return average(xLeft,xRight), NumericResultType_HasSolution, nil
		}

		xRightOld := xRight

		fxRight := F(xRight)
		fxLeft := F(xLeft)

		if AnyNanOrInfinity(fxLeft,fxRight){
			return xRight, NumericResultType_NoSolution, FunctionValueIsNanOrInfinityErr
		}

		deltaF := fxRight - fxLeft
		if deltaF == 0 {
			return xRight, NumericResultType_NoSolution, FunctionsDeltaIsZeroErr
		}

		xRight = (xLeft*F(xRight) - xRight*F(xLeft))/deltaF
		xLeft = xRightOld

		iterationPassed++
	}

	return xRight, NumericResultType_NoSolution, nil
}