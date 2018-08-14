package newtonMethod

import (
	. "math"
	. "XIRR/numMethods"
)

type NewtonMethod struct {
	InitialGuess float64
}

// NumericMethodUsingDerivative interface implementation
func (nm *NewtonMethod) Calculate(F NumericFunc, derivativeF NumericFunc, methodParams *NumericMethodParams) (float64, NumericResultType, *NumericMethodError) {

	xCurrent := nm.InitialGuess

	var iterationPassed  uint64 = 0
	for iterationPassed < methodParams.MaxIterationsCount {

		fValue := F(xCurrent)
		fDerivativeValue := derivativeF(xCurrent)

		xNext := xCurrent - fValue/fDerivativeValue

		dx := Abs(xNext-xCurrent)
		if dx <= methodParams.Epsilon {
			return xNext, NumericResultType_HasSolution, nil
		}
		xCurrent = xNext

		iterationPassed++
	}

	return xCurrent, NumericResultType_NoSolution, nil
}