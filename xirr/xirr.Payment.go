package xirr

import "time"

//Basic implementation of Payment interface
type Payment struct {

	//amount of payment
	_amount float64

	//date of payment
	_date time.Time
}

func NewPayment(amount float64, date time.Time) *Payment {
	return &Payment{_amount: amount, _date: date}
}

func (p *Payment) Amount() float64 {
	return p._amount
}

func (p *Payment) Date() *time.Time {
	return &p._date
}

func (p *Payment) Before(s IPayment) bool {
	return p._date.Before(*s.Date())
}

func (p *Payment) After(s IPayment) bool {
	return p._date.After(*s.Date())
}
