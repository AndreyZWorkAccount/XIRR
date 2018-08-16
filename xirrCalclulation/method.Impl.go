package xirrCalclulation

import (
	. "../numMethods"
	. "../netPresentValue"
	. "math"
	"../newton"
	"../secantAuto"
	"sort"
)

//XIRR numeric method
type XIRRCalculationMethod struct {

	daysInYear uint16

	params *NumericMethodParams

	minRateOfIrrDecrease float64
}

func NewMethod( minRateOfIrr float64, daysInYear uint16, methodParams *NumericMethodParams) XIRRCalculationMethod{
	return XIRRCalculationMethod{  daysInYear, methodParams, minRateOfIrr}
}



//XIRRCalcMethod implementation
func (method XIRRCalculationMethod) Calculate(payments []IPayment) (result float64, resultType NumericResultType, error *NumericMethodError) {

	if len(payments) == 0 {
		return NoSolutionFound()
	}

	//order payments by date
	sort.Slice(payments, func(i,j int) bool { return payments[i].Before(payments[j]) })
	startPaymentDate := payments[0].Date()

	//NPV function
	F := func(irr float64) float64{
		return NetPresentValue(irr, payments, startPaymentDate, method.daysInYear)
	}
	//NPV derivative
	derivativeF := func(irr float64) float64{
		return NetPresentValueDerivative(irr, payments, startPaymentDate, method.daysInYear)
	}
	//NPV second derivative
	secondDerivativeF := func(irr float64) float64{
		return NetPresentValueSecondDerivative(irr, payments, startPaymentDate, method.daysInYear)
	}

	paymentsSumIsPositive := IsPaymentsSumPositive(payments)

	bestSolution := Solution{x:Inf(1), fx:Inf(1)}

	//secant
	bordersSearchAlg := secantAuto.NewBordersSearchAlgorithm(F, derivativeF)
	secantMethod := secantAuto.NewMethod(paymentsSumIsPositive, bordersSearchAlg , method.minRateOfIrrDecrease)
	xSecant, resType, _ := secantMethod.Calculate(F, derivativeF, secondDerivativeF, method.params)
	canFinish, bestSolution := updateBestSolution(xSecant, F, resType, bestSolution)
	if canFinish{
		return SolutionFound(bestSolution.x)
	}

	//newton
	guess := 0.1
	if !paymentsSumIsPositive{ guess = -0.1 }

	//try use guess
	canFinish, bestSolution = tryNewton(guess,F,derivativeF,method.params,bestSolution)
	if canFinish{
		return SolutionFound(bestSolution.x)
	}

	//try use negative guess
	canFinish, bestSolution = tryNewton(-guess,F,derivativeF,method.params,bestSolution)
	if canFinish{
		return SolutionFound(bestSolution.x)
	}

	// In case of we failed both guesses and total sum is negative then try negative guess close to -1
	if !paymentsSumIsPositive{
		canFinish, bestSolution = tryNewton(IrrDefaultValue,F,derivativeF,method.params,bestSolution)
		if canFinish{
			return SolutionFound(bestSolution.x)
		}
	}

	return ErrorFound(&AllNumericMethodsHaveBeenFailed)
}

func tryNewton(guess float64, F, derivativeF NumericFunc, methodParams *NumericMethodParams, bestSolution Solution ) (canFinish bool, newBest Solution) {
	newtonMethod := newton.NewMethod(guess)
	xNewton, resType, _ := newtonMethod.Calculate(F,derivativeF,methodParams)
	return updateBestSolution(xNewton, F, resType, bestSolution)
}

func updateBestSolution(x float64, f NumericFunc, resType NumericResultType, currentBest Solution) (canFinish bool, newBest Solution){
	if resType == NumericResultType_HasSolution{
		fx := f(x)
		bestSolution := BestSolutionOf(currentBest, Solution{x,fx })
		if Abs(IdealNPV - bestSolution.fx) <= IrrEpsilon {
			return true, bestSolution
		}
	}
	return false, currentBest
}

func BestSolutionOf(old, new Solution ) Solution{
	if Abs(new.fx) < Abs(old.fx){
		return new
	}
	return old
}

func IsPaymentsSumPositive(payments []IPayment) bool{
	var totalSum = 0.0
	for _,payment := range payments{
		totalSum += payment.Amount()
	}
	return totalSum > 0.0
}


type Solution struct {
	x, fx float64
}

