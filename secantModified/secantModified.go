package secantModified

import (
	. "math"
	. "../numMethods"
	. "../float.Extensions"
)


type Method struct {

	xLeftInit, xRightInit float64

	//Min allowed value of dX[i] - dX[i-1]. Allows to indicate when iterations aren't effective anymore.
	minimumRateOfXDecrease float64
}

func NewMethod(xLeft, xRight, minRateOfXDecrease float64) Method {
	return Method{xLeftInit: xLeft, xRightInit:xRight, minimumRateOfXDecrease:minRateOfXDecrease}
}





// NumericMethodUsingSecondDerivative interface implementation
func (method *Method) Calculate(F NumericFunc, derivativeF NumericFunc, secondDerivativeF NumericFunc,
	methodParams *NumericMethodParams) (float64, NumericResultType, *NumericMethodError){

	xLeft := method.xLeftInit
	xRight := method.xRightInit
	isConvergesOnLeft := isConvergesOnLeft(xLeft,F,secondDerivativeF)

	var iterationPassed  uint64 = 0
	var solutionFound = false
	var err *NumericMethodError = nil

	for iterationPassed < methodParams.MaxIterationsCount {
		xLeft, xRight, solutionFound, err = method.runIteration(xLeft, xRight, isConvergesOnLeft, methodParams, F, derivativeF, secondDerivativeF)
		if err != nil {
			return ErrorFound(err)
		}
		if solutionFound {
			return SolutionFound(Average(xLeft,xRight))
		}
		iterationPassed++
	}

	return NoSolutionFound()
}

func (method *Method) runIteration(xLeft float64, xRight float64, isConvergesOnLeft bool, methodParams *NumericMethodParams, F NumericFunc, derivativeF NumericFunc, secondDerivativeF NumericFunc)(xLeftOut float64, xRightOut float64, solutionFound bool,err *NumericMethodError ) {
	prevIterationDx := Abs(xRight - xLeft)

	FxRight := F(xRight)
	FxLeft := F(xLeft)

	if AnyNanOrInfinity(FxLeft,FxRight){
		return xLeft, xRight, false, FunctionValueIsNanOrInfinityErr
	}

	if isConvergesOnLeft {
		dFxLeft := derivativeF(xLeft)
		xRight = xRight - ((xRight - xLeft)/(FxRight - FxLeft))*FxRight
		xLeft = xLeft - FxLeft/dFxLeft
	} else {
		dFxRight := derivativeF(xRight)
		xRight = xRight - FxRight/dFxRight
		xLeft = xLeft - ((xRight - xLeft)/(FxRight - FxLeft))*FxLeft
	}

	dx := Abs(xRight - xLeft)
	if dx < methodParams.Epsilon {
		return xLeft,xRight, true, nil
	}

	if Abs(dx - prevIterationDx) < method.minimumRateOfXDecrease{
		xAvg := Average(xLeft,xRight)
		FxAvg := F(xAvg)
		if (FxLeft < 0 && FxAvg > 0) || (FxLeft > 0 && FxAvg < 0){
			xRight = xAvg
		} else {
			xLeft = xAvg
		}
	}

	return xLeft, xRight, false, nil
}

func isConvergesOnLeft(xLeft float64, F NumericFunc, secondDerivativeF NumericFunc ) (bool){
	FxLeft := F(xLeft)
	ddFxLeft := secondDerivativeF(xLeft)
	if (FxLeft <= 0 && ddFxLeft <= 0) || (FxLeft >= 0 && ddFxLeft >= 0) {
		return true
	}
	return false
}

