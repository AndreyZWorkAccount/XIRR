package numMethods

//type to represent a numeric method
type NumericMethod interface {
	Calculate(F INumericFunc, methodParams *NumericMethodParams) (float64, NumericResultType, *NumericMethodError)
}

//type to represent a numeric method using derivative of function
type NumericMethodUsingDerivative interface {
	Calculate(F INumericFunc, derivativeF INumericFunc, methodParams *NumericMethodParams) (float64, NumericResultType, *NumericMethodError)
}

//type to a numeric method using second derivative of function
type NumericMethodUsingSecondDerivative interface {
	Calculate(F INumericFunc, derivativeF INumericFunc, secondDerivativeF INumericFunc, methodParams *NumericMethodParams) (float64, NumericResultType, *NumericMethodError)
}

// parameters of numeric methods
type NumericMethodParams struct {

	// max count of method iterations
	MaxIterationsCount uint64

	// acceptable calculation error
	Epsilon float64
}


// type to represent status of calculations
type NumericResultType int
const (
	NumericResultType_HasSolution NumericResultType = 1 + iota
	NumericResultType_NoSolution
)






