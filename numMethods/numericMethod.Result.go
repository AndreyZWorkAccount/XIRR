package numMethods



func ErrorFound(err *NumericMethodError) (result float64, resultType NumericResultType, error *NumericMethodError){
	result = 0
	resultType = NumericResultType_NoSolution
	error = err
	return
}

func NoSolutionFound() (result float64, resultType NumericResultType, error *NumericMethodError){
	result = 0
	resultType = NumericResultType_NoSolution
	error = nil
	return
}

func SolutionFound(res float64) (result float64, resultType NumericResultType, error *NumericMethodError){
	result = res
	resultType = NumericResultType_HasSolution
	error = nil
	return
}
