package numMethods


type SecantModifiedMethod struct {
	XLeftInit, XRightInit float64
}

// NumericMethodUsingSecondDerivative interface implementation
func (s *SecantModifiedMethod) Calculate(F NumericFunc, derivativeF NumericFunc, secondDerivativeF NumericFunc, methodParams *NumericMethodParams) (float64, NumericResultType, *NumericMethodError){







	return 0, NumericResultType_NoSolution, nil
}
