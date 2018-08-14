package secantModifiedMethod

import . "XIRR/netPresentValue"

type IBordersSearchAlgorithm interface {
	//Returns a slice that contains optimal initial borders for 'secant' method applying
	FindInitialBorders(payments []IPayment) []IBorder
}

type IBorder interface {
	Left() float64
	Right() float64
}
