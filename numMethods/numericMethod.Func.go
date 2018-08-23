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
