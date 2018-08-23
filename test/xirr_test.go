package test

import (
	. "testing"
	. "math"

	. "github.com/AndreyZWorkAccount/XIRR/numMethods"
	"github.com/AndreyZWorkAccount/XIRR/xirrCalclulation"
)

func TestIrr(t *T){
	methodParams := NumericMethodParams{MaxIterationsCount: 1000,Epsilon:0.0000001}

	for _,testCase := range TestCases{
		var xirrCalcMethod = xirrCalclulation.NewMethod( 0.00000001, 365, &methodParams )
		res, resType, err := xirrCalcMethod.Calculate(testCase.Payments)

		if !resType.IsSolution(){
			t.Error("Successful solution is expected.")
		}
		if err != nil{
			t.Error(err.Description)
		}
		if Abs(res - testCase.ExpectedValue) > 0.0000001{
			t.Errorf("Expected: %v\n. Actual: %v\n", testCase.ExpectedValue, res)
		}
	}
}
