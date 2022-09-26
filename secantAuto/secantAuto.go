// Copyright 2018 Andrey Z. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package secantAuto

import (
	. "github.com/krazybee/XIRR/numMethods"
	"github.com/krazybee/XIRR/secantModified"
)

type Method struct {
	paymentsSumIsPositive bool

	bordersSearchAlg IBordersSearchAlgorithm

	//Min allowed value of dX[i] - dX[i-1]. Allows to indicate when iterations aren't effective anymore.
	minimumRateOfXDecrease float64
}

func NewMethod(paymentsSumIsPositive bool, searchAlg IBordersSearchAlgorithm, minRateOfXDecrease float64) Method {
	return Method{paymentsSumIsPositive: paymentsSumIsPositive, bordersSearchAlg: searchAlg, minimumRateOfXDecrease: minRateOfXDecrease}
}

// NumericMethodUsingSecondDerivative interface implementation
func (method *Method) Calculate(F INumericFunc, derivativeF INumericFunc, secondDerivativeF INumericFunc, methodParams *Params) IResult {

	borders := method.bordersSearchAlg.FindInitialBorders(method.paymentsSumIsPositive)

	//iterate over borders and return first successful result
	for _, border := range borders {
		secantModified := secantModified.NewMethod(border.Left(), border.Right(), method.minimumRateOfXDecrease)
		ans := secantModified.Calculate(F, derivativeF, secondDerivativeF, methodParams)

		if ans.IsSolution() {
			return ans
		}
	}
	return NoSolutionFound()
}
