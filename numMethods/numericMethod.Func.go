// Copyright 2018 Andrey Z. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package numMethods

// type to represent a function
type INumericFunc interface{
	ApplyTo(arg float64) float64
}

//Default implementation
type NumFunc func(arg float64) float64
func (f NumFunc) ApplyTo(arg float64) float64{
	return f(arg)
}
