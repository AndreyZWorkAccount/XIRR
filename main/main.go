package main

import  (
	. "XIRR/numMethods"
	. "XIRR/newtonMethod"
	. "XIRR/secantMethod"
	. "XIRR/netPresentValue"
	. "fmt"
	"sort"
	"XIRR/test"
	)

func main() {

    for _,testCase := range test.TestCases{
    	Println("-------------New test case-----------")

		XIRR(testCase.Payments, 365)

    	Printf("Expected: %v", testCase.ExpectedValue)
    	Println()
	}
}

// IRR returns the Internal Rate of Return (IRR).
func XIRR(payments []IPayment, daysInYear uint16) {
	if len(payments) == 0 {
		return
	}

	//order payments by date
	sort.Slice(payments, func(i,j int) bool { return payments[i].Before(payments[j]) })

	startPaymentDate := payments[0].Date()

	//NPV function
	f := func(irr float64) float64{
		return NetPresentValue(irr, payments, startPaymentDate, daysInYear)
	}
	//NPV derivative
	df := func(irr float64) float64{
		return NetPresentValueDerivative(irr, payments, startPaymentDate, daysInYear)
	}
	//NPV second derivative
	ddf := func(irr float64) float64{
		return NetPresentValueSecondDerivative(irr, payments, startPaymentDate, daysInYear)
	}

	RunNumericMethods(f, df, ddf)

	return
}

func RunNumericMethods(f NumericFunc, df NumericFunc, ddf NumericFunc){

	methodParams := NumericMethodParams{MaxIterationsCount: 1000,Epsilon:0.0000001}

	//newton
	var newtonMethod = &NewtonMethod{InitialGuess: 0.1}
	res, _, err := newtonMethod.Calculate(f, df, &methodParams)
	PrintResult("Newton", err, res)

	//secant
	var secantMethod = &SecantMethod{XLeftInit: -0.2,XRightInit:0.0001}
	res, _, err = secantMethod.Calculate(f, &methodParams)
	PrintResult("Secant", err, res)

	//modified secant
	var modifiedSecantMethod = &SecantModifiedMethod{XLeftInit: -0.2,XRightInit:0.0001, MinimumRateOfXDecrease: 0.00000001}
	res, _, err = modifiedSecantMethod.Calculate(f, df, ddf, &methodParams)
	PrintResult("Modified Secant", err, res)
}

func PrintResult(methodName string,  err *NumericMethodError, res float64) {
	if err == nil {
		Printf("%v method gives root at x=%f\n", methodName, res)
	} else {
		Printf("%s", err.Error())
	}
}



