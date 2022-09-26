// Copyright 2018 Andrey Z. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package test

import (
	. "math"
	. "testing"

	"sync"

	"github.com/krazybee/XIRR/numMethods"
	"github.com/krazybee/XIRR/xirrAsync"
)

func TestAsyncIrr_OneRequestInChan(t *T) {
	var coresCount int = 5
	var processor xirrAsync.IProcessor = xirrAsync.NewProcessor()
	processor.Start(coresCount)

	requests := processor.Requests()
	responses := processor.Responses()

	for id, testCase := range TestCases {
		requests <- xirrAsync.NewRequest(int64(id), testCase.Payments)
		resp := <-responses
		verifyTestResult(resp.Result(), testCase.ExpectedValue, t)
	}
}

func TestAsyncIrr_ManyRequestsInChan(t *T) {
	var coresCount int = 5
	var processor xirrAsync.IProcessor = xirrAsync.NewProcessor()
	processor.Start(coresCount)

	requests := processor.Requests()
	responses := processor.Responses()

	testCases := make(map[int]TestCase)
	for i, t := range TestCases {
		testCases[i] = t
	}

	wg := sync.WaitGroup{}

	for id, testCase := range testCases {
		requests <- xirrAsync.NewRequest(int64(id), testCase.Payments)

		go func() {
			wg.Add(1)
			defer wg.Done()

			resp := <-responses
			expected := testCases[int(resp.RequestId())].ExpectedValue
			verifyTestResult(resp.Result(), expected, t)
		}()
	}

	wg.Wait()
}

func verifyTestResult(res numMethods.IResult, expectedValue float64, t *T) {

	if !res.IsSolution() {
		t.Error("Successful solution is expected.")
	}
	if res.Error() != nil {
		t.Error(res.Error())
	}
	if Abs(res.Value()-expectedValue) > 0.0000000001 {
		t.Errorf("Expected: %v\n. Actual: %v\n", expectedValue, res.Value())
	}
}
