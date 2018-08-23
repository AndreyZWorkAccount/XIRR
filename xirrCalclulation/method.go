// Copyright 2018 Andrey Z. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package xirrCalclulation

import (
	. "github.com/AndreyZWorkAccount/XIRR/netPresentValue"
	. "github.com/AndreyZWorkAccount/XIRR/numMethods"
)


type XIRRCalcMethod interface {
	Calculate(payments []IPayment) (result float64, resultType NumericResultType, error *NumericMethodError)
}