package numMethods

//type to represent a numeric method
type NumericMethod interface {
	Calculate(F NumericFunc, methodParams *NumericMethodParams) (float64, NumericResultType, *NumericMethodError)
}

//type to represent a numeric method using derivative of function
type NumericMethodUsingDerivative interface {
	Calculate(F NumericFunc, derivativeF NumericFunc, methodParams *NumericMethodParams) (float64, NumericResultType, *NumericMethodError)
}

// type to represent a function
type NumericFunc func(float64) float64

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
	NumericResultType_HasSolution NumericResultType = 0
	NumericResultType_NoSolution  NumericResultType = 1
)
