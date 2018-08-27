package xirr

import (
	. "github.com/AndreyZWorkAccount/XIRR/numMethods"
	. "github.com/AndreyZWorkAccount/XIRR/time.Extensions"
	. "math"
)


//XIRR numeric method with deannualisation support
type XIRRDeAnnualizeMethod struct {
	XIRRMethod
}

func NewXIRRDeAnnualizeMethod( minRateOfIrr float64, daysInYear uint16, methodParams *NumericMethodParams) XIRRDeAnnualizeMethod{
	return XIRRDeAnnualizeMethod{XIRRMethod{daysInYear, methodParams, minRateOfIrr}}
}



//XIRRCalcMethod implementation
func (method XIRRDeAnnualizeMethod) Calculate(payments IOrderedPayments) (result float64, resultType NumericResultType, error *NumericMethodError) {

	if payments.Count() == 0 {
		return NoSolutionFound()
	}

	res, resType, err := method.XIRRMethod.Calculate(payments)

	if err != nil || !resType.IsSolution(){
		return res, resType, err
	}

	//deannualize if solution found
	return method.deAnnualize(res, payments), NumericResultType_HasSolution, nil
}

func (method XIRRDeAnnualizeMethod) deAnnualize(res float64, payments IOrderedPayments) (result float64) {
	allPayments := payments.GetAll()

	lastDate := allPayments[len(allPayments)-1].Date()
	firstDate := allPayments[0].Date()

	maxDiffInDays := DiffInDays(lastDate, firstDate)
	daysInYearF := float64(method.daysInYear)

	if maxDiffInDays < daysInYearF {
		return Pow(1.0 + res, maxDiffInDays/daysInYearF ) - 1.0
	}

	return res
}

