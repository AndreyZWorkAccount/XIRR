package main

import (
	"NumericMethods/numMethods"
	"fmt"
	"NumericMethods/xirr"
	"sort"
		"time"
	)

func main() {

	testValues := []xirr.IPayment{
		xirr.NewPayment(-55506, OnDate(1,time.January,2000)),
		xirr.NewPayment(8340, OnDate(6,time.February,2001)),
		xirr.NewPayment(-293, OnDate(28,time.March,2001)),
	}

	XIRR(testValues, 365)
}

// IRR returns the Internal Rate of Return (IRR).
func XIRR(payments []xirr.IPayment, daysInYear uint16) (float64, *numMethods.NumericMethodError) {
	if len(payments) == 0 {
		return 0, nil
	}

	//order payments by date
	sort.Slice(payments, func(i,j int) bool { return payments[i].Before(payments[j]) })

	startPaymentDate := payments[0].Date()

	//NPV function
	f := func(irr float64) float64{
		return xirr.NetPresentValue(irr, payments, startPaymentDate, daysInYear)
	}

	//NPV derivative
	df := func(irr float64) float64{
		return xirr.NetPresentValueDerivative(irr, payments, startPaymentDate, daysInYear)
	}

	methodParams := numMethods.NumericMethodParams{MaxIterationsCount:1000,Epsilon:0.0000001}

	//newton
	var newtonMethod = &numMethods.NewtonMethod{InitialGuess:-0.999900}
	res, _, err := newtonMethod.Calculate(f, df, &methodParams)
	PrintResult("Newton", err, res)

	//secant
	var secantMethod = &numMethods.SecantMethod{XLeftInit:-0.9,XRightInit:-0.2}
	res, _, err = secantMethod.Calculate(f, &methodParams)
	PrintResult("Secant", err, res)

	return res, err
}

func PrintResult(methodName string,  err *numMethods.NumericMethodError, res float64) {
	if err == nil {
		fmt.Printf("%v method gives root at x=%f\n", methodName, res)
	} else {
		fmt.Printf("%s", err.Error())
	}
}

func OnDate(day int, month time.Month, year int ) time.Time{
	return time.Date(year,month,day,12,0,0,0, time.UTC);
}


