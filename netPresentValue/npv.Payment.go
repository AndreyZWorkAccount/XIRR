// Copyright 2018 Andrey Z. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package netPresentValue

import . "time"

//Basic implementation of Payment interface
type Payment struct {

	//amount of payment
	_amount float64

	//date of payment
	_date Time
}

func NewPayment(amount float64, date Time) *Payment{
	return &Payment{_amount:amount, _date:date}
}


func (p *Payment) Amount() float64{
	return p._amount
}

func (p *Payment) Date() *Time {
	return &p._date
}

func (p *Payment) Before( s IPayment) bool{
	return p._date.Before(*s.Date())
}

func (p *Payment) After( s IPayment) bool{
	return p._date.After(*s.Date())
}
