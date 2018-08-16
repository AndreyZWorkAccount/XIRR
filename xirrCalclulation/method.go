package xirrCalclulation

import (
	. "github.com/AndreyZWorkAccount/XIRR/netPresentValue"
	. "github.com/AndreyZWorkAccount/XIRR/numMethods"
)


type XIRRCalcMethod interface {
	Calculate(payments []IPayment) (result float64, resultType NumericResultType, error *NumericMethodError)
}