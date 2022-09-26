// Copyright 2018 Andrey Z. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package secantAuto

import (
	. "github.com/krazybee/XIRR/float.Extensions"
	. "github.com/krazybee/XIRR/netPresentValue"
	. "github.com/krazybee/XIRR/numMethods"
)

type BordersSearchAlgorithm struct {
	f INumericFunc

	derivativeF INumericFunc
}

func NewBordersSearchAlgorithm(F INumericFunc, dF INumericFunc) BordersSearchAlgorithm {
	return BordersSearchAlgorithm{f: F, derivativeF: dF}
}

//IBordersSearchAlgorithm implementation
func (alg BordersSearchAlgorithm) FindInitialBorders(paymentsSumIsPositive bool) []IBorder {

	//try to find borders with default guess
	borders := alg.tryGetBorders(paymentsSumIsPositive, false)
	if len(borders) > 0 {
		return borders
	}

	//try to use different guess
	borders = alg.tryGetBorders(paymentsSumIsPositive, true)
	if len(borders) > 0 {
		return borders
	}

	return []IBorder{}
}

func (alg *BordersSearchAlgorithm) tryGetBorders(paymentsSumIsPositive bool, useValidation bool) (borders []IBorder) {
	var left, right float64
	var options ValidationOptions

	left, right, options = getBordersSearchParams(paymentsSumIsPositive, useValidation)
	borders = alg.findInitialBorders(left, right, options)

	return
}

func (alg *BordersSearchAlgorithm) findInitialBorders(leftInitial float64,
	rightInitial float64,
	options ValidationOptions) (borders []IBorder) {

	res := make([]IBorder, 0)

	currentBorder := NewBorder(leftInitial, rightInitial)
	iterationsPassed := 0

	worthTryLeft, worthTryRight := true, true

	for iterationsPassed < Max_Iterations {

		iterationResult := alg.findInitialBorder(currentBorder, worthTryLeft, worthTryRight, options)

		if iterationResult.isSolution {
			res = append(res, &iterationResult.ansBorder)
		}

		if !iterationResult.canContinue {
			return res
		}

		currentBorder = iterationResult.nextBorder
		worthTryLeft = iterationResult.tryLeft
		worthTryRight = iterationResult.tryRight

		iterationsPassed++
	}

	return res
}

func (alg *BordersSearchAlgorithm) findInitialBorder(border Border, tryLeft, tryRight bool, options ValidationOptions) BorderIterationResult {

	xLeft := border.left
	xRight := border.right

	FxLeft := alg.f.ApplyTo(border.left)
	FxRight := alg.f.ApplyTo(border.right)

	if AnyNanOrInfinity(FxLeft, FxRight) {
		return NoSolutionAndBreak()
	}

	if (FxLeft > 0 && FxRight > 0) || (FxLeft < 0 && FxRight < 0) {
		if options.useBorders || !tryLeft && !tryRight {
			return NoSolutionAndBreak()
		}
		if tryLeft {
			xLeft, tryLeft = tryGoLeft(xLeft, options)
		}
		if tryRight {
			xRight, tryRight = tryGoRight(xRight, options)
		}
		return NoSolutionAndContinue(Border{xLeft, xRight}, tryLeft, tryRight)
	}

	dFxLeft := alg.derivativeF.ApplyTo(xLeft)
	dFxRight := alg.derivativeF.ApplyTo(xRight)

	if (dFxLeft > 0 && dFxRight > 0) || (dFxLeft < 0 && dFxRight < 0) {
		ansBorder := Border{xLeft, xRight}
		nextBorder := Border{xRight, xRight + Delta_For_Borders}
		return SolutionAndContinue(nextBorder, ansBorder, tryLeft, tryRight)
	}

	//try to iterate from left to right
	for left := xLeft + Iteration_Step; left < xRight; left += Iteration_Step {
		FLeft := alg.f.ApplyTo(left)

		if (FxLeft > 0 && FLeft > 0) || (FxLeft < 0 && FLeft < 0) {
			continue
		}

		ansBorder := Border{xLeft, xRight}
		nextBorder := Border{xRight, xRight + Delta_For_Borders}
		return SolutionAndContinue(nextBorder, ansBorder, tryLeft, tryRight)
	}

	return NoSolutionAndBreak()
}

func tryGoLeft(xLeft float64, options ValidationOptions) (newXLeft float64, worthTryLeft bool) {
	newVal := xLeft - Delta_For_Borders

	if options.validateBordersMinMax && newVal <= options.minLeft {
		return options.minLeft, false
	}
	if newVal <= IrrMinValue {
		return IrrDefaultValue, false
	}
	return newVal, true
}

func tryGoRight(xRight float64, options ValidationOptions) (newXRight float64, worthTryRight bool) {
	var newVal float64
	if xRight > 1.0 {
		newVal = xRight + Delta_For_Borders*Delta_For_Borders_Multiplexer
	} else {
		newVal = xRight + Delta_For_Borders
	}

	if options.validateBordersMinMax && newVal >= options.maxRight {
		return options.maxRight, false
	}
	return newVal, true
}
