// Copyright 2018 Andrey Z. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package test

import (
	. "testing"
	. "math"

	. "github.com/AndreyZWorkAccount/XIRR/numMethods"
	"github.com/AndreyZWorkAccount/XIRR/xirrCalclulation"
)

func TestIrr(t *T){
	methodParams := NumericMethodParams{MaxIterationsCount: 1000,Epsilon:0.0000001}

	for _,testCase := range TestCases{
		var xirrCalcMethod = xirrCalclulation.NewXIRRMethod( 0.00000001, 365, &methodParams )

		var orderedPayments = xirrCalclulation.OrderPayments(testCase.Payments)
		res, resType, err := xirrCalcMethod.Calculate(orderedPayments)

		if !resType.IsSolution(){
			t.Error("Successful solution is expected.")
		}
		if err != nil{
			t.Error(err)
		}
		if Abs(res - testCase.ExpectedValue) > 0.0000000001{
			t.Errorf("Expected: %v\n. Actual: %v\n", testCase.ExpectedValue, res)
		}
	}
}
