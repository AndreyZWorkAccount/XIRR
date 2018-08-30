// Copyright 2018 Andrey Z. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package xirrAsync


//Factory method
func NewProcessor() IProcessor{

	requests := make(chan IRequest)
	responses := make(chan IResponse)
	cancellations := make([] chan interface{}, 0)

	return &RequestsProcessor{ requests,  responses, cancellations}
}