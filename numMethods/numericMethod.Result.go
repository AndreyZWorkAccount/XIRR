// Copyright 2018 Andrey Z. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package numMethods

type IResult interface {
	Value() float64
	IsSolution()  bool
	Error() *NumericMethodError
}

func ErrorFound(err *NumericMethodError) IResult{
	return &resultImpl{0, NumericResultType_NoSolution, err}
}

func NoSolutionFound() IResult{
	return &resultImpl{0, NumericResultType_NoSolution, nil}
}

func SolutionFound(res float64) IResult{
	return &resultImpl{res, NumericResultType_HasSolution, nil }
}


//Result impl
type resultImpl struct {
	value float64
	valueType  NumericResultType
	error *NumericMethodError
}
func (r *resultImpl) Value() float64{
	return r.value
}
func (r *resultImpl) IsSolution()  bool{
	return r.valueType == NumericResultType_HasSolution
}
func (r *resultImpl) Error() *NumericMethodError{
	return r.error
}