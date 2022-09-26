// Copyright 2018 Andrey Z. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package test

import (
	. "math"
	. "testing"

	. "github.com/krazybee/XIRR/numMethods"
	"github.com/krazybee/XIRR/xirr"
)

func TestIrr(t *T) {
	methodParams := Params{MaxIterationsCount: 1000, Epsilon: 0.0000001}

	for _, testCase := range TestCases {
		var xirrCalcMethod xirr.CalcMethod = xirr.NewXIRRMethod(0.00000001, 365, &methodParams)

		var orderedPayments = xirr.OrderPayments(testCase.Payments)
		res := xirrCalcMethod.Calculate(orderedPayments)

		if !res.IsSolution() {
			t.Error("Successful solution is expected.")
		}
		if res.Error() != nil {
			t.Error(res.Error())
		}
		if Abs(res.Value()-testCase.ExpectedValue) > 0.0000000001 {
			t.Errorf("Expected: %v\n. Actual: %v\n", testCase.ExpectedValue, res)
		}
	}
}
