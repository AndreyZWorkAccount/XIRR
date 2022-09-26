// Copyright 2018 Andrey Z. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package test

import (
	. "time"

	. "github.com/krazybee/XIRR/netPresentValue"
)

type TestCase struct {
	Payments      []IPayment
	ExpectedValue float64
}

var TestCases = []TestCase{

	TestCase{
		Payments: []IPayment{
			NewPayment(1000000.00, OnDate(30, April, 2016)),
			NewPayment(-2937480.00, OnDate(1, January, 2016)),
			NewPayment(-28000000.00, OnDate(1, April, 2016)),
			NewPayment(28837484.00, OnDate(15, April, 2016)),
			NewPayment(1029877.00, OnDate(23, April, 2016)),
			NewPayment(372384.00, OnDate(20, April, 2016)),
		},
		ExpectedValue: 0.1623411031457017},

	/*TestCase{
	Payments: []IPayment{
		NewPayment(-63823, OnDate(1,January,2000 )),
		NewPayment(710, OnDate(26,August,2000 )),
		NewPayment(693, OnDate(11,July,2001 )),
	},
	ExpectedValue:-0.9456970110868315},*/

	TestCase{
		Payments: []IPayment{
			NewPayment(-2937480.00, OnDate(1, January, 2016)),
			NewPayment(-28000000.00, OnDate(1, April, 2016)),
			NewPayment(28837484.00, OnDate(15, April, 2016)),
			NewPayment(372384.00, OnDate(20, April, 2016)),
			NewPayment(1029877.00, OnDate(23, April, 2016)),
			NewPayment(1000000.00, OnDate(30, April, 2016)),
		},
		ExpectedValue: 0.1623411031457017},

	TestCase{
		Payments: []IPayment{
			NewPayment(20, OnDate(1, January, 2014)),
			NewPayment(-100, OnDate(1, January, 2012)),
			NewPayment(10, OnDate(1, January, 2013)),
			NewPayment(30, OnDate(1, January, 2015)),
		},
		ExpectedValue: -0.1922001575},

	TestCase{
		Payments: []IPayment{
			NewPayment(-2937480.00, OnDate(1, January, 2016)),
			NewPayment(-28000000.00, OnDate(1, April, 2016)),
			NewPayment(28837484.00, OnDate(15, April, 2016)),
			NewPayment(372384.00, OnDate(20, May, 2016)),
			NewPayment(1029877.00, OnDate(23, July, 2016)),
			NewPayment(1000000.00, OnDate(30, July, 2016)),
		},
		ExpectedValue: 0.12679876439643994},

	TestCase{
		Payments: []IPayment{
			NewPayment(-46, OnDate(1, January, 2000)),
			NewPayment(668, OnDate(31, August, 2000)),
			NewPayment(1453, OnDate(11, August, 2001)),
			NewPayment(1225, OnDate(22, January, 2003)),
			NewPayment(-282, OnDate(6, October, 2003)),
			NewPayment(1155, OnDate(26, January, 2004)),
			NewPayment(1570, OnDate(25, August, 2004)),
			NewPayment(1225, OnDate(1, December, 2005)),
			NewPayment(1376, OnDate(16, April, 2006)),
			NewPayment(358, OnDate(2, May, 2006)),
			NewPayment(-200, OnDate(21, March, 2007)),
			NewPayment(921, OnDate(25, April, 2007)),
			NewPayment(302, OnDate(21, March, 2008)),
			NewPayment(-39, OnDate(24, March, 2008)),
			NewPayment(-80, OnDate(28, February, 2010)),
		},
		ExpectedValue: 58.5175804368881},

	TestCase{
		Payments: []IPayment{
			NewPayment(-123400, OnDate(1, January, 2012)),
			NewPayment(48100, OnDate(1, January, 2015)),
			NewPayment(36200, OnDate(1, January, 2013)),
			NewPayment(54800, OnDate(1, January, 2014)),
		},
		ExpectedValue: 0.0595345222},

	TestCase{
		Payments: []IPayment{
			NewPayment(-17139, OnDate(1, January, 2000)),
			NewPayment(795, OnDate(31, August, 2000)),
			NewPayment(-344, OnDate(1, October, 2000)),
		},
		ExpectedValue: -0.999912992599877},

	TestCase{
		Payments: []IPayment{
			NewPayment(-5100, OnDate(25, June, 2015)),
			NewPayment(-800, OnDate(9, September, 2015)),
			NewPayment(2500, OnDate(10, September, 2015)),
			NewPayment(500, OnDate(11, September, 2015)),
			NewPayment(-200, OnDate(12, September, 2015)),
			NewPayment(1800, OnDate(13, September, 2015)),
			NewPayment(500, OnDate(14, September, 2015)),
			NewPayment(100, OnDate(21, September, 2015)),
			NewPayment(100, OnDate(24, September, 2015)),
			NewPayment(-5100, OnDate(25, September, 2015)),
			NewPayment(5100, OnDate(15, October, 2015)),
			NewPayment(-100, OnDate(17, October, 2015)),
			NewPayment(800, OnDate(18, October, 2015)),
			NewPayment(500, OnDate(19, October, 2015)),
			NewPayment(200, OnDate(20, October, 2015)),
			NewPayment(500, OnDate(22, October, 2015)),
			NewPayment(-800, OnDate(23, October, 2015)),
		},
		ExpectedValue: 0.3875114316258743},

	TestCase{
		Payments: []IPayment{
			NewPayment(-13463.0, OnDate(1, January, 2000)),
			NewPayment(-111, OnDate(11, September, 2000)),
			NewPayment(1859, OnDate(30, June, 2001)),
		},
		ExpectedValue: -0.7374722671574048},

	TestCase{
		Payments: []IPayment{
			NewPayment(-14123, OnDate(1, January, 2000)),
			NewPayment(1243, OnDate(19, September, 2000)),
			NewPayment(1578, OnDate(10, November, 2001)),
			NewPayment(1268, OnDate(10, May, 2002)),
			NewPayment(159, OnDate(9, March, 2003)),
			NewPayment(-48, OnDate(26, February, 2005)),
			NewPayment(1217, OnDate(28, June, 2006)),
			NewPayment(-289, OnDate(1, October, 2007)),
			NewPayment(998, OnDate(6, December, 2008)),
			NewPayment(214, OnDate(16, January, 2009)),
			NewPayment(965, OnDate(3, July, 2009)),
			NewPayment(-114, OnDate(27, April, 2010)),
			NewPayment(-118, OnDate(24, March, 2011)),
			NewPayment(-106, OnDate(4, September, 2011)),
			NewPayment(630, OnDate(1, October, 2011)),
			NewPayment(-489, OnDate(22, August, 2012)),
		},
		ExpectedValue: -0.12982362249643103},

	TestCase{
		Payments: []IPayment{
			NewPayment(-2769, OnDate(1, January, 2000)),
			NewPayment(1811, OnDate(21, September, 2000)),
			NewPayment(1642, OnDate(14, June, 2000)),
			NewPayment(1898, OnDate(30, December, 2001)),
		},
		ExpectedValue: 1.0032052076959754},

	TestCase{
		Payments: []IPayment{
			NewPayment(-75, OnDate(1, January, 2000)),
			NewPayment(-425, OnDate(23, October, 2001)),
			NewPayment(-249, OnDate(21, May, 2002)),
			NewPayment(1191, OnDate(2, May, 2003)),
			NewPayment(1172, OnDate(26, September, 2004)),
			NewPayment(785, OnDate(29, October, 2004)),
			NewPayment(869, OnDate(2, August, 2006)),
			NewPayment(879, OnDate(8, September, 2007)),
			NewPayment(-306, OnDate(27, October, 2007)),
			NewPayment(83, OnDate(4, July, 2009)),
			NewPayment(443, OnDate(6, January, 2011)),
			NewPayment(462, OnDate(15, August, 2012)),
			NewPayment(194, OnDate(3, July, 2014)),
		},
		ExpectedValue: 0.8656190240806798},

	TestCase{
		Payments: []IPayment{
			NewPayment(1920, OnDate(1, February, 2000)),
			NewPayment(-10773, OnDate(1, January, 2000)),
			NewPayment(1198, OnDate(22, February, 2001)),
		},
		ExpectedValue: -0.8202639662898871},

	TestCase{
		Payments: []IPayment{
			NewPayment(-1285, OnDate(1, January, 2000)),
			NewPayment(556, OnDate(17, July, 2000))},
		ExpectedValue: -0.7865449910261024},

	TestCase{
		Payments: []IPayment{
			NewPayment(-9915, OnDate(1, January, 2000)),
			NewPayment(255, OnDate(25, December, 2000)),
			NewPayment(550, OnDate(31, July, 2002)),
			NewPayment(1294, OnDate(13, April, 2003)),
			NewPayment(1848, OnDate(20, March, 2005)),
			NewPayment(433, OnDate(29, September, 2005)),
		},
		ExpectedValue: -0.17347618547013388},

	TestCase{
		Payments: []IPayment{
			NewPayment(-9779, OnDate(1, January, 2000)),
			NewPayment(342, OnDate(6, September, 2001)),
			NewPayment(-123, OnDate(28, December, 2001)),
			NewPayment(1024, OnDate(29, March, 2002)),
			NewPayment(1406, OnDate(5, October, 2003)),
			NewPayment(-228, OnDate(9, August, 2005)),
			NewPayment(1541, OnDate(14, July, 2007)),
		},
		ExpectedValue: -0.16221514453356106},

	TestCase{
		Payments: []IPayment{
			NewPayment(-8950, OnDate(1, January, 2000)),
			NewPayment(1103, OnDate(10, September, 2001)),
			NewPayment(1733, OnDate(31, October, 2002)),
			NewPayment(1828, OnDate(18, January, 2004)),
			NewPayment(660, OnDate(30, April, 2005)),
		},
		ExpectedValue: -0.14075459394958345},

	TestCase{
		Payments: []IPayment{
			NewPayment(-8866, OnDate(1, January, 2000)),
			NewPayment(72, OnDate(22, May, 2002)),
			NewPayment(757, OnDate(10, October, 2000)),
			NewPayment(560, OnDate(10, August, 2001)),
			NewPayment(-270, OnDate(27, November, 2004)),
			NewPayment(203, OnDate(23, March, 2003)),
			NewPayment(1714, OnDate(15, June, 2006)),
			NewPayment(715, OnDate(14, September, 2007)),
		},
		ExpectedValue: -0.15055389256088803},

	TestCase{
		Payments: []IPayment{
			NewPayment(-2477, OnDate(1, January, 2000)),
			NewPayment(951, OnDate(13, June, 2000)),
			NewPayment(726, OnDate(15, December, 2000)),
			NewPayment(645, OnDate(14, March, 2002)),
			NewPayment(523, OnDate(10, April, 2002)),
			NewPayment(133, OnDate(18, October, 2003)),
			NewPayment(500, OnDate(6, June, 2005)),
			NewPayment(-2, OnDate(18, August, 2006)),
		},
		ExpectedValue: 0.21172166842038695,
	},
}

func OnDate(day int, month Month, year int) Time {
	return Date(year, month, day, 12, 0, 0, 0, UTC)
}
