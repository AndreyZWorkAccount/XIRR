package test


import (
	. "testing"

	"github.com/AndreyZWorkAccount/XIRR/xirrAsync"
	"runtime"
	"time"
)


func TestRequestProcessor(t *T){
	var coresCount int = 50000;
	var timeout time.Duration = 1*time.Second

	initialGoroutinesCount := runtime.NumGoroutine()

	var processor xirrAsync.IProcessor = xirrAsync.NewProcessor()

	processor.Start(coresCount)

	if runtime.NumGoroutine() - initialGoroutinesCount != coresCount{
		t.Errorf("Number of created goroutines should be %v, but it's %v", coresCount + initialGoroutinesCount, runtime.NumGoroutine())
	}

	for{
		select {

		case success := <- processor.Stop():
			if success{
				if runtime.NumGoroutine() != initialGoroutinesCount{
					t.Error("XIRR processor has been stopped, but goroutines are still running.")
				}
			}else {
				t.Error("Can't stop XIRR processor.")
			}
		return

		case <- time.After(timeout):
			t.Error("Processor termination timeout.")
		return

		}
	}

}
