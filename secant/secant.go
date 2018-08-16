package secant

import (
	. "math"
	. "github.com/AndreyZWorkAccount/XIRR/numMethods"
	. "github.com/AndreyZWorkAccount/XIRR/float.Extensions"
	)

type Method struct {
	xLeftInit, xRightInit float64
}

func NewMethod(xLeft, xRight float64) Method{
	return Method{xLeft, xRight}
}


// NumericMethod interface implementation
func (s *Method) Calculate(F NumericFunc, methodParams *NumericMethodParams) (float64, NumericResultType, *NumericMethodError){
	xLeft := s.xLeftInit
	xRight := s.xRightInit

	var iterationPassed  uint64 = 0
	for iterationPassed < methodParams.MaxIterationsCount {

		//check if we reach necessary precision
		dx := Abs(xRight - xLeft)
		if dx < methodParams.Epsilon{
			return SolutionFound(Average(xLeft,xRight))
		}

		xRightOld := xRight

		fxRight := F(xRight)
		fxLeft := F(xLeft)

		if AnyNanOrInfinity(fxLeft,fxRight){
			return ErrorFound(FunctionValueIsNanOrInfinityErr)
		}

		deltaF := fxRight - fxLeft
		if deltaF == 0 {
			return ErrorFound(FunctionsDeltaIsZeroErr)
		}

		xRight = (xLeft*F(xRight) - xRight*F(xLeft))/deltaF
		xLeft = xRightOld

		iterationPassed++
	}

	return NoSolutionFound()
}