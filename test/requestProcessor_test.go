// Copyright 2018 Andrey Z. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package test

import (
	. "testing"

	"runtime"
	"time"

	"github.com/krazybee/XIRR/xirrAsync"
)

func TestRequestProcessorStartAndStop(t *T) {
	var coresCount int = 50000
	var timeout time.Duration = 1 * time.Second

	initialGoroutinesCount := runtime.NumGoroutine()

	var processor xirrAsync.IProcessor = xirrAsync.NewProcessor()
	processor.Start(coresCount)

	if runtime.NumGoroutine()-initialGoroutinesCount != coresCount {
		t.Errorf("Number of created goroutines should be %v, but it's %v", coresCount+initialGoroutinesCount, runtime.NumGoroutine())
	}

	for {
		select {

		case isOk := <-processor.Stop():
			if !isOk {
				t.Error("Can't stop XIRR processor.")
				return
			}
			if runtime.NumGoroutine() != initialGoroutinesCount {
				t.Error("XIRR processor has been stopped, but goroutines are still running.")
				return
			}
			return

		case <-time.After(timeout):
			t.Error("Processor termination timeout.")
			return

		}
	}
}
