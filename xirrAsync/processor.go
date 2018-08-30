// Copyright 2018 Andrey Z. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package xirrAsync

type IProcessor interface {

	Requests() chan IRequest

	Responses() <-chan IResponse

	Start(coresCount int)

	Stop() <- chan bool
}


