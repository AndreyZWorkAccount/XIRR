// Copyright 2018 Andrey Z. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package xirr

import (
	. "math"

	. "github.com/krazybee/XIRR/netPresentValue"
	"github.com/krazybee/XIRR/newton"
	. "github.com/krazybee/XIRR/numMethods"
	"github.com/krazybee/XIRR/secantAuto"
)

//XIRR numeric method
type XIRRMethod struct {
	daysInYear uint16

	params *Params

	minRateOfIrrDecrease float64
}

func NewXIRRMethod(minRateOfIrr float64, daysInYear uint16, methodParams *Params) XIRRMethod {
	return XIRRMethod{daysInYear, methodParams, minRateOfIrr}
}

//XIRRCalcMethod implementation
func (method XIRRMethod) Calculate(payments IOrderedPayments) IResult {

	if payments.Count() == 0 {
		return NoSolutionFound()
	}

	allPayments := payments.GetAll()
	startPaymentDate := allPayments[0].Date()

	//NPV function
	F := NumFunc(func(irr float64) float64 {
		return NPV(irr, allPayments, startPaymentDate, method.daysInYear)
	})

	//NPV derivative
	derivativeF := NumFunc(func(irr float64) float64 {
		return NPVDerivative(irr, allPayments, startPaymentDate, method.daysInYear)
	})

	//NPV second derivative
	secondDerivativeF := NumFunc(func(irr float64) float64 {
		return NPVSecondDerivative(irr, allPayments, startPaymentDate, method.daysInYear)
	})

	paymentsSumIsPositive := IsPaymentsSumPositive(allPayments)

	bestSolution := solution{x: Inf(1), fx: Inf(1)}

	//secant
	bordersSearchAlg := secantAuto.NewBordersSearchAlgorithm(F, derivativeF)
	secantMethod := secantAuto.NewMethod(paymentsSumIsPositive, bordersSearchAlg, method.minRateOfIrrDecrease)

	xSecant := secantMethod.Calculate(F, derivativeF, secondDerivativeF, method.params)
	canFinish, bestSolution := updateBestSolution(xSecant, F, bestSolution)
	if canFinish {
		return SolutionFound(bestSolution.x)
	}

	//newton
	guess := 0.1
	if !paymentsSumIsPositive {
		guess = -0.1
	}

	//try use guess
	canFinish, bestSolution = tryNewton(guess, F, derivativeF, method.params, bestSolution)
	if canFinish {
		return SolutionFound(bestSolution.x)
	}

	//try use negative guess
	canFinish, bestSolution = tryNewton(-guess, F, derivativeF, method.params, bestSolution)
	if canFinish {
		return SolutionFound(bestSolution.x)
	}

	// In case of we failed both guesses then try negative guess close to -1
	canFinish, bestSolution = tryNewton(IrrDefaultValue, F, derivativeF, method.params, bestSolution)
	if canFinish {
		return SolutionFound(bestSolution.x)
	}

	return ErrorFound(AllNumericMethodsHaveBeenFailed)
}

func tryNewton(guess float64, F, derivativeF INumericFunc, methodParams *Params, bestSolution solution) (canFinish bool, newBest solution) {
	newtonMethod := newton.NewMethod(guess)
	xNewton := newtonMethod.Calculate(F, derivativeF, methodParams)
	return updateBestSolution(xNewton, F, bestSolution)
}

func updateBestSolution(x IResult, F INumericFunc, currentBest solution) (canFinish bool, newBest solution) {
	if x.IsSolution() {
		bestSolution := bestSolutionOf(currentBest, solution{x.Value(), F.ApplyTo(x.Value())})
		if closeEnough(IdealNPV, bestSolution.fx) {
			return true, bestSolution
		}
	}
	return false, currentBest
}

func closeEnough(first, second float64) bool {
	return Abs(second-first) <= IrrEpsilon
}

func bestSolutionOf(old, new solution) solution {
	if Abs(new.fx) < Abs(old.fx) {
		return new
	}
	return old
}

func IsPaymentsSumPositive(payments []IPayment) bool {
	var totalSum = 0.0
	for _, payment := range payments {
		totalSum += payment.Amount()
	}
	return totalSum > 0.0
}

type solution struct {
	x, fx float64
}
