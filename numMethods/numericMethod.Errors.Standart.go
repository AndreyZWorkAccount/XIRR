// Copyright 2018 Andrey Z. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package numMethods

var FunctionsDeltaIsZeroErr = error(functionsDeltaIsZero)

var FunctionValueIsNanOrInfinityErr = error(functionValueIsNanOrInfinity)

var FunctionHasNoSolutionInIntervalErr = error(functionHasNoSolutionInInterval)

var AllNumericMethodsHaveBeenFailed = error(allNumericMethodsHaveBeenFailed)



func error(description string) *NumericMethodError{
	return &NumericMethodError{description}
}
