// Copyright 2018 Andrey Z. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package newton

import (
	. "math"

	. "github.com/krazybee/XIRR/numMethods"
)

type Method struct {
	initialGuess float64
}

func NewMethod(guess float64) Method {
	return Method{initialGuess: guess}
}

// NumericMethodUsingDerivative interface implementation
func (nm *Method) Calculate(F INumericFunc, derivativeF INumericFunc, methodParams *Params) IResult {

	xCurrent := nm.initialGuess

	var iterationPassed uint64 = 0
	for iterationPassed < methodParams.MaxIterationsCount {

		fValue := F.ApplyTo(xCurrent)
		fDerivativeValue := derivativeF.ApplyTo(xCurrent)

		xNext := xCurrent - fValue/fDerivativeValue

		dx := Abs(xNext - xCurrent)
		if dx <= methodParams.Epsilon {
			return SolutionFound(xNext)
		}
		xCurrent = xNext

		iterationPassed++
	}

	return NoSolutionFound()
}
