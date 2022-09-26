// Copyright 2018 Andrey Z. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package test

import (
	. "testing"

	"github.com/krazybee/XIRR/xirr"
)

func TestNewOrderedPayments(t *T) {

	for _, testCase := range TestCases {
		orderedPayments := xirr.OrderPayments(testCase.Payments)

		allPayments := orderedPayments.GetAll()

		for i := 1; i < len(allPayments); i++ {
			currentPayment := allPayments[i]
			prevPayment := allPayments[i-1]

			afterPrevPayment := currentPayment.After(prevPayment)
			atTheSameDate := currentPayment.SameDateAs(prevPayment)

			if !(afterPrevPayment || atTheSameDate) {
				t.Errorf("Payment with date %v should be after payment with date %v", prevPayment.Date(), currentPayment.Date())
			}
		}
	}
}
