package main

import  (
	. "../numMethods"
	. "../xirr"
	. "fmt"
	. "sort"
	. "time"
	)

func main() {

	testValues := []IPayment{
		NewPayment(-55506, OnDate(1, January,2000)),
		NewPayment(8340, OnDate(6, February,2001)),
		NewPayment(-293, OnDate(28, March,2001)),
	}

	XIRR(testValues, 365)
}

// IRR returns the Internal Rate of Return (IRR).
func XIRR(payments []IPayment, daysInYear uint16) (float64, *NumericMethodError) {
	if len(payments) == 0 {
		return 0, nil
	}

	//order payments by date
	Slice(payments, func(i,j int) bool { return payments[i].Before(payments[j]) })

	startPaymentDate := payments[0].Date()

	//NPV function
	f := func(irr float64) float64{
		return NetPresentValue(irr, payments, startPaymentDate, daysInYear)
	}

	//NPV derivative
	df := func(irr float64) float64{
		return NetPresentValueDerivative(irr, payments, startPaymentDate, daysInYear)
	}

	methodParams := NumericMethodParams{MaxIterationsCount: 1000,Epsilon:0.0000001}

	//newton
	var newtonMethod = &NewtonMethod{InitialGuess: -0.999900}
	res, _, err := newtonMethod.Calculate(f, df, &methodParams)
	PrintResult("Newton", err, res)

	//secant
	var secantMethod = &SecantMethod{XLeftInit: -0.9,XRightInit:-0.2}
	res, _, err = secantMethod.Calculate(f, &methodParams)
	PrintResult("Secant", err, res)

	return res, err
}

func PrintResult(methodName string,  err *NumericMethodError, res float64) {
	if err == nil {
		Printf("%v method gives root at x=%f\n", methodName, res)
	} else {
		Printf("%s", err.Error())
	}
}

func OnDate(day int, month Month, year int ) Time {
	return Date(year,month,day,12,0,0,0, UTC);
}


