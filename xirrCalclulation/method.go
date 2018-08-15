package xirrCalclulation

import (
	. "XIRR/netPresentValue"
	. "XIRR/numMethods"
)


type XIRRCalcMethod interface {
	Calculate(payments []IPayment) (result float64, resultType NumericResultType, error *NumericMethodError)
}