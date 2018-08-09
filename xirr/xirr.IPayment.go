package xirr

import . "time"

//Payment interface
type IPayment interface {

	//amount of payment
	Amount() float64

	//date of payment
	Date() *Time

	//True if current payment is before p
	Before(p IPayment) bool

	//True if current payment is after p
	After(p IPayment) bool
}





