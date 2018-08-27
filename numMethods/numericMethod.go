// Copyright 2018 Andrey Z. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package numMethods

//type to represent a numeric method
type NumericMethod interface {
	Calculate(F INumericFunc, methodParams *Params) IResult
}

//type to represent a numeric method using derivative of function
type NumericMethodUsingDerivative interface {
	Calculate(F INumericFunc, derivativeF INumericFunc, methodParams *Params) IResult
}

//type to a numeric method using second derivative of function
type NumericMethodUsingSecondDerivative interface {
	Calculate(F INumericFunc, derivativeF INumericFunc, secondDerivativeF INumericFunc, methodParams *Params) IResult
}




// parameters of numeric methods
type Params struct {
	// max count of method iterations
	MaxIterationsCount uint64
	// acceptable calculation error
	Epsilon float64
}

// type to represent status of calculations
type NumericResultType int
const (
	NumericResultType_HasSolution NumericResultType = 1 + iota
	NumericResultType_NoSolution
)






