package secantAuto

import (
	. "github.com/AndreyZWorkAccount/XIRR/numMethods"
	"github.com/AndreyZWorkAccount/XIRR/secantModified"
	)

type Method struct {

	paymentsSumIsPositive bool

	bordersSearchAlg IBordersSearchAlgorithm

	//Min allowed value of dX[i] - dX[i-1]. Allows to indicate when iterations aren't effective anymore.
	minimumRateOfXDecrease float64
}

func NewMethod(paymentsSumIsPositive bool, searchAlg IBordersSearchAlgorithm, minRateOfXDecrease float64) Method{
	return Method{paymentsSumIsPositive:paymentsSumIsPositive, bordersSearchAlg:searchAlg, minimumRateOfXDecrease:minRateOfXDecrease}
}




// NumericMethodUsingSecondDerivative interface implementation
func (method *Method) Calculate(F NumericFunc, derivativeF NumericFunc, secondDerivativeF NumericFunc,
	methodParams *NumericMethodParams) (float64, NumericResultType, *NumericMethodError){

		borders := method.bordersSearchAlg.FindInitialBorders(method.paymentsSumIsPositive)

		//iterate over borders and return first successful result
		for _, border := range borders{
			secantModified := secantModified.NewMethod(border.Left(), border.Right(), method.minimumRateOfXDecrease)
			ans, resType, _ := secantModified.Calculate(F,derivativeF,secondDerivativeF,methodParams)

			if resType == NumericResultType_HasSolution{
				return SolutionFound(ans)
			}
		}
		return NoSolutionFound()
}
