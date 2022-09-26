// Copyright 2018 Andrey Z. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package xirrAsync

import (
	. "github.com/krazybee/XIRR/netPresentValue"
	. "github.com/krazybee/XIRR/numMethods"
	. "github.com/krazybee/XIRR/xirr"
)

type RequestsProcessor struct {
	requests      chan IRequest
	responses     chan IResponse
	cancellations []chan interface{}
}

func (p *RequestsProcessor) Requests() chan IRequest {
	return p.requests
}

func (p *RequestsProcessor) Responses() <-chan IResponse {
	return p.responses
}

func (p *RequestsProcessor) Stop() <-chan bool {
	ch := make(chan bool)
	go stopCores(p, ch)
	return ch
}

func (p *RequestsProcessor) Start(coresCount int) {
	p.cancellations = startCores(coresCount, p.requests, p.responses)
}

//Private
func stopCores(p *RequestsProcessor, cancellation chan bool) {
	for _, c := range p.cancellations {
		c <- struct{}{}
		<-c
	}
	p.cancellations = make([]chan interface{}, 0)
	cancellation <- true
}

func startCores(coresCount int, requests chan IRequest, responses chan IResponse) []chan interface{} {
	var cancellations []chan interface{}
	for i := 0; i < coresCount; i++ {
		cancellation := make(chan interface{})
		go runProcessorCore(requests, responses, cancellation)
		cancellations = append(cancellations, cancellation)
	}
	return cancellations
}

func runProcessorCore(requests chan IRequest, responses chan IResponse, cancellation chan interface{}) {
	for {
		select {

		case request := <-requests:
			res := calculateXirr(request.Payments())
			responses <- NewResponse(request.RequestId(), res)
			continue

		case <-cancellation:
			close(cancellation)
			return

		}
	}
}

func calculateXirr(payments []IPayment) IResult {
	methodParams := Params{MaxIterationsCount: 1000, Epsilon: 0.0000001}
	var method CalcMethod = NewXIRRMethod(0.00000001, 365, &methodParams)

	var orderedPayments = OrderPayments(payments)
	return method.Calculate(orderedPayments)
}
