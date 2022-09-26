// Copyright 2018 Andrey Z. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package secant

import (
	. "math"

	. "github.com/krazybee/XIRR/float.Extensions"
	. "github.com/krazybee/XIRR/numMethods"
)

type Method struct {
	xLeftInit, xRightInit float64
}

func NewMethod(xLeft, xRight float64) Method {
	return Method{xLeft, xRight}
}

// NumericMethod interface implementation
func (s *Method) Calculate(F INumericFunc, methodParams *Params) Result {
	xLeft := s.xLeftInit
	xRight := s.xRightInit

	var iterationPassed uint64 = 0
	for iterationPassed < methodParams.MaxIterationsCount {

		//check if we reach necessary precision
		dx := Abs(xRight - xLeft)
		if dx < methodParams.Epsilon {
			return SolutionFound(Average(xLeft, xRight))
		}

		xRightOld := xRight

		fxRight := F.ApplyTo(xRight)
		fxLeft := F.ApplyTo(xLeft)

		if AnyNanOrInfinity(fxLeft, fxRight) {
			return ErrorFound(FunctionValueIsNanOrInfinityErr)
		}

		deltaF := fxRight - fxLeft
		if deltaF == 0 {
			return ErrorFound(FunctionsDeltaIsZeroErr)
		}

		xRight = (xLeft*F.ApplyTo(xRight) - xRight*F.ApplyTo(xLeft)) / deltaF
		xLeft = xRightOld

		iterationPassed++
	}

	return NoSolutionFound()
}
