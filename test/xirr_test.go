package test

import (
	. "testing"
	. "github.com/AndreyZWorkAccount/XIRR/numMethods"
	. "github.com/AndreyZWorkAccount/XIRR/netPresentValue"
	. "time"
	. "math"
	"github.com/AndreyZWorkAccount/XIRR/xirrCalclulation"
)

type TestCase struct {
	Payments []IPayment
	ExpectedValue float64
}


func TestXIRR(t *T) {

	testCases := []TestCase{

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

		TestCase{
			Payments: []IPayment{
				NewPayment(-46, OnDate(1,January,2000)),
				NewPayment(668, OnDate(31,August,2000)),
				NewPayment(1453, OnDate(11,August,2001)),
				NewPayment(1225, OnDate(22,January,2003)),
				NewPayment(-282, OnDate(6,October,2003)),
				NewPayment(1155, OnDate(26,January,2004)),
				NewPayment(1570, OnDate(25,August,2004)),
				NewPayment(1225, OnDate(1,December,2005)),
				NewPayment(1376, OnDate(16,April,2006)),
				NewPayment(358, OnDate(2,May,2006)),
				NewPayment(-200, OnDate(21,March,2007)),
				NewPayment(921, OnDate(25,April,2007)),
				NewPayment(302, OnDate(21,March,2008)),
				NewPayment(-39, OnDate(24,March,2008)),
				NewPayment(-80, OnDate(28,February,2010)),
			},
			ExpectedValue: 58.5175807952881,
		},

		TestCase{
			Payments: []IPayment{
				NewPayment(-123400,OnDate(1,January,2012)),
				NewPayment(36200,OnDate(1,January,2013)),
				NewPayment(54800,OnDate(1,January,2014)),
				NewPayment(48100,OnDate(1,January,2015)),
			},
			ExpectedValue: 0.05953,
		},

		TestCase{
			Payments: []IPayment{
				NewPayment(-17139,OnDate(1,January,2000)),
				NewPayment(795,OnDate(31,August,2000)),
				NewPayment(-344,OnDate(1,October,2000)),
			},
			ExpectedValue: -0.999913135099877,
		},

		TestCase{
			Payments: []IPayment{
				NewPayment(-5100,OnDate(25,June,2015)),
				NewPayment(-800,OnDate(9,September,2015)),
				NewPayment(2500,OnDate(10,September,2015)),
				NewPayment(500,OnDate(11,September,2015)),
				NewPayment(-200,OnDate(12,September,2015)),
				NewPayment(1800,OnDate(13,September,2015)),
				NewPayment(500,OnDate(14,September,2015)),
				NewPayment(100,OnDate(21,September,2015)),
				NewPayment(100,OnDate(24,September,2015)),
				NewPayment(-5100,OnDate(25,September,2015)),
				NewPayment(5100,OnDate(15,October,2015)),
				NewPayment(-100,OnDate(17,October,2015)),
				NewPayment(800,OnDate(18,October,2015)),
				NewPayment(500,OnDate(19,October,2015)),
				NewPayment(200,OnDate(20,October,2015)),
				NewPayment(500,OnDate(22,October,2015)),
				NewPayment(-800,OnDate(23,October,2015)),
			},
			ExpectedValue: 0.38751143,
		},

		TestCase{
			Payments: []IPayment{
				NewPayment(-13463.0,OnDate(1,January,2000)),
				NewPayment(-111,OnDate(11,September,2000)),
				NewPayment(1859,OnDate(30,June,2001)),
			},
			ExpectedValue: -0.73747226819396,
		},

		TestCase{
			Payments: []IPayment{
				NewPayment(-2769,OnDate(1,January,2000)),
				NewPayment(1642,OnDate(14,June,2000)),
				NewPayment(1811,OnDate(21,September,2000)),
				NewPayment(1898,OnDate(30,December,2001)),
			},
			ExpectedValue: 1.00320521593094,
		},

		TestCase{
			Payments: []IPayment{
				NewPayment(-75,OnDate(1,January,2000)),
				NewPayment(-425,OnDate(23,October,2001)),
				NewPayment(-249,OnDate(21,May,2002)),
				NewPayment(1191,OnDate(2,May,2003)),
				NewPayment(1172,OnDate(26,September,2004)),
				NewPayment(785,OnDate(29,October,2004)),
				NewPayment(869,OnDate(2,August,2006)),
				NewPayment(879,OnDate(8,September,2007)),
				NewPayment(-306,OnDate(27,October,2007)),
				NewPayment(83,OnDate(4,July,2009)),
				NewPayment(443,OnDate(6,January,2011)),
				NewPayment(462,OnDate(15,August,2012)),
				NewPayment(194,OnDate(3,July,2014)),
			},
			ExpectedValue: 0.865619027614594,
		},

		TestCase{
			Payments: []IPayment{
				NewPayment(-10773,OnDate(1,January,2000)),
				NewPayment(1920,OnDate(1,February,2000)),
				NewPayment(1198,OnDate(22,February,2001)),
			},
			ExpectedValue: -0.82026396356523,
		},

		TestCase{
			Payments: []IPayment{
				NewPayment(-1285,OnDate(1,January,2000)),
				NewPayment(556,OnDate(17,July,2000)),},
			ExpectedValue: -0.786544985324144,
		},

	}

	_testXIRR(t, testCases)
}


func _testXIRR(t *T, testCases []TestCase) {

	methodParams := NumericMethodParams{MaxIterationsCount: 1000,Epsilon:0.0000001}

	for _,testCase := range testCases{

		var xirrCalcMethod = xirrCalclulation.NewMethod( 0.00000001, 365, &methodParams )
		res, resType, err := xirrCalcMethod.Calculate(testCase.Payments)

		if resType != NumericResultType_HasSolution{
			t.Error("Successful solution is expected.")
		}

		if err != nil{
			t.Error(err.Description)
		}

		if Abs(res - testCase.ExpectedValue) > 0.000001{
			t.Errorf("Expected: %v\n. Actual: %v\n", testCase.ExpectedValue, res)
		}
	}
}



func OnDate(day int, month Month, year int ) Time {
	return Date(year,month,day,12,0,0,0, UTC);
}
