package xirr

import (
	. "math"

	. "github.com/krazybee/XIRR/numMethods"
	. "github.com/krazybee/XIRR/time.Extensions"
)

//XIRR numeric method with deannualisation support
type XIRRDeAnnualizeMethod struct {
	XIRRMethod
}

func NewXIRRDeAnnualizeMethod(minRateOfIrr float64, daysInYear uint16, methodParams *Params) XIRRDeAnnualizeMethod {
	return XIRRDeAnnualizeMethod{XIRRMethod{daysInYear, methodParams, minRateOfIrr}}
}

//XIRRCalcMethod implementation
func (method XIRRDeAnnualizeMethod) Calculate(payments IOrderedPayments) IResult {

	if payments.Count() == 0 {
		return NoSolutionFound()
	}

	res := method.XIRRMethod.Calculate(payments)
	if res.Error != nil || !res.IsSolution() {
		return res
	}

	//deannualize if solution found
	return SolutionFound(method.deAnnualize(res.Value(), payments))
}

func (method XIRRDeAnnualizeMethod) deAnnualize(res float64, payments IOrderedPayments) (result float64) {
	allPayments := payments.GetAll()

	lastDate := allPayments[len(allPayments)-1].Date()
	firstDate := allPayments[0].Date()

	maxDiffInDays := DiffInDays(lastDate, firstDate)
	daysInYearF := float64(method.daysInYear)

	if maxDiffInDays < daysInYearF {
		return Pow(1.0+res, maxDiffInDays/daysInYearF) - 1.0
	}

	return res
}
