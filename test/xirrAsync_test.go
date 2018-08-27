package test


import (
	. "testing"
	. "math"

	"github.com/AndreyZWorkAccount/XIRR/xirrAsync"
	"github.com/AndreyZWorkAccount/XIRR/numMethods"
	)

func TestAsyncIrr_ExecuteAsSync(t *T){
	var coresCount int = 5;

	var processor xirrAsync.IProcessor = xirrAsync.NewProcessor()
	processor.Start(coresCount)

	requests := processor.Requests()
	responses := processor.Responses()

	for id,testCase := range TestCases{
		requests <- xirrAsync.NewRequest(int64(id), testCase.Payments)
		res := <- responses

		verifyTestResult(res.Result(), testCase.ExpectedValue, t)
	}
}


func verifyTestResult(res numMethods.IResult, expectedValue float64, t *T) {
	if !res.IsSolution() {
		t.Error("Successful solution is expected.")
	}
	if res.Error() != nil {
		t.Error(res.Error())
	}
	if Abs(res.Value() - expectedValue) > 0.0000000001 {
		t.Errorf("Expected: %v\n. Actual: %v\n", expectedValue, res.Value())
	}
}

