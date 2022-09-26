// Copyright 2018 Andrey Z. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package netPresentValue

import (
	. "math"
	. "time"

	. "github.com/krazybee/XIRR/time.Extensions"
)

//NetPresentValue
func NPV(irrValue float64, payments []IPayment, firstPaymentDate *Time, daysInYear uint16) float64 {
	if irrValue <= IrrMinValue {
		irrValue = IrrDefaultValue
	}

	daysInYearF := float64(daysInYear)

	var npv = 0.0
	for _, payment := range payments {
		npv += netPresentValue(irrValue, payment, firstPaymentDate, daysInYearF)
	}
	return npv
}

//NetPresentValue derivative
func NPVDerivative(irrValue float64, payments []IPayment, firstPaymentDate *Time, daysInYear uint16) float64 {
	if irrValue <= IrrMinValue {
		irrValue = IrrDefaultValue
	}

	daysInYearF := float64(daysInYear)

	var npv = 0.0
	for _, payment := range payments {
		npv += netPresentValueDerivative(irrValue, payment, firstPaymentDate, daysInYearF)
	}
	return npv
}

//NetPresentValue second derivative
func NPVSecondDerivative(irrValue float64, payments []IPayment, firstPaymentDate *Time, daysInYear uint16) float64 {
	if irrValue <= IrrMinValue {
		irrValue = IrrDefaultValue
	}

	daysInYearF := float64(daysInYear)

	var npv = 0.0
	for _, payment := range payments {
		npv += netPresentValueSecondDerivative(irrValue, payment, firstPaymentDate, daysInYearF)
	}
	return npv
}

func netPresentValue(irrValue float64, payment IPayment, startDate *Time, daysInYear float64) float64 {
	var diffInDays = DiffInDays(payment.Date(), startDate)
	return payment.Amount() * Pow(1.0+irrValue, diffInDays/daysInYear)
}
func netPresentValueDerivative(irrValue float64, payment IPayment, startDate *Time, daysInYear float64) float64 {
	diffInDays := DiffInDays(payment.Date(), startDate)
	return (payment.Amount() * Pow(1.0+irrValue, (diffInDays/daysInYear)-1.0)) * (diffInDays / daysInYear)
}
func netPresentValueSecondDerivative(irrValue float64, payment IPayment, startDate *Time, daysInYear float64) float64 {
	if irrValue <= IrrMinValue {
		irrValue = IrrDefaultValue
	}
	diffInDays := DiffInDays(payment.Date(), startDate)

	return (payment.Amount() * Pow(1.0+irrValue, (diffInDays/daysInYear)-2.0)) * (diffInDays / daysInYear) * (diffInDays/daysInYear - 1.0)
}
