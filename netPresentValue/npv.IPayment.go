// Copyright 2018 Andrey Z. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package netPresentValue

import . "time"

//Payment interface
type IPayment interface {

	//amount of payment
	Amount() float64

	//date of payment
	Date() *Time

	//True if current payment is before p
	Before(p IPayment) bool

	//True if current payment is after p
	After(p IPayment) bool

	//True if current payment is at the sane date as p
	SameDateAs(p IPayment) bool
}





