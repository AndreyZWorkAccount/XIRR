package xirrCalclulation

import (
	. "../netPresentValue"
	. "../numMethods"
)


type XIRRCalcMethod interface {
	Calculate(payments []IPayment) (result float64, resultType NumericResultType, error *NumericMethodError)
}