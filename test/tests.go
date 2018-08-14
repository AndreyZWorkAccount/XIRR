package test

import (
	. "XIRR/netPresentValue"
	. "time"
	)


type TestCase struct {
	Payments []IPayment
	ExpectedValue float64
}

var TestCases = []TestCase {

	TestCase{
		Payments:[]IPayment{
			NewPayment(-2937480.00, OnDate(1, January,2016)),
			NewPayment(-28000000.00, OnDate(1, April,2016)),
			NewPayment(28837484.00, OnDate(15, April,2016)),
			NewPayment(372384.00, OnDate(20, April,2016)),
			NewPayment(1029877.00, OnDate(23, April,2016)),
			NewPayment(1000000.00, OnDate(30, April,2016)),
		},
		ExpectedValue: 0.162341112},

	TestCase{
		Payments: []IPayment{
			NewPayment(-100, OnDate(1, January,2012)),
			NewPayment(10, OnDate(1, January,2013)),
			NewPayment(20, OnDate(1, January ,2014)),
			NewPayment(30, OnDate(1, January,2015)),
		},
		ExpectedValue: -0.1922},

	TestCase{
		Payments: []IPayment{
			NewPayment(-2937480.00, OnDate(1, January, 2016)),
			NewPayment(-28000000.00, OnDate(1, April,2016)),
			NewPayment(28837484.00, OnDate(15, April,2016)),
			NewPayment(372384.00, OnDate(20, May,2016)),
			NewPayment(1029877.00, OnDate(23, July,2016)),
			NewPayment(1000000.00, OnDate(30, July,2016)),
		},
		ExpectedValue: 0.12679876439643994},
}



func OnDate(day int, month Month, year int ) Time {
	return Date(year,month,day,12,0,0,0, UTC);
}



