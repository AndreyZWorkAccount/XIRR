package main

import  (
	. "../numMethods"
	. "fmt"
	"../xirrCalclulation"
	"../test"
)

func main() {

    for _,testCase := range test.TestCases{
    	Println("*********************New test case********************")

		methodParams := NumericMethodParams{MaxIterationsCount: 1000,Epsilon:0.0000001}

		var xirrCalcMethod = xirrCalclulation.NewMethod( 0.00000001, 365, &methodParams )
		res, resType, err := xirrCalcMethod.Calculate(testCase.Payments)
		PrintResult("XIRR method", err, res, resType)

    	Printf("Expected:                          %v", testCase.ExpectedValue)
    	Println()
	}
}


func PrintResult(methodName string,  err *NumericMethodError, res float64, resType NumericResultType) {

	if resType == NumericResultType_NoSolution{
		Printf("%v method                   No solution\n", methodName)
		return
	}

	if err != nil {
		Printf("%s", err.Error())
		return
	}

	Printf("%v method gives root at x=%f\n", methodName, res)
}



