package secantAuto

type IBordersSearchAlgorithm interface {
	//Returns a slice that contains optimal initial borders for 'secant' method applying
	FindInitialBorders(paymentsSumIsPositive bool) []IBorder
}

type IBorder interface {
	Left() float64
	Right() float64
}
