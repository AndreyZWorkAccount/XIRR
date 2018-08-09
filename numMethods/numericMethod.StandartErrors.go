package numMethods

var FunctionsDeltaIsZeroErr = error(functionsDeltaIsZero)

var FunctionValueIsNanOrInfinityErr = error(functionValueIsNanOrInfinity)

var FunctionHasNoSolutionInIntervalErr = error(functionHasNoSolutionInInterval)



func error(description string) *NumericMethodError{
	return &NumericMethodError{description}
}
