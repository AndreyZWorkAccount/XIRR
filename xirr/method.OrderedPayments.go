package xirr

import (
	"sort"

	. "github.com/krazybee/XIRR/netPresentValue"
)

type IOrderedPayments interface {
	GetAll() []IPayment

	Count() int
}

func OrderPayments(payments []IPayment) IOrderedPayments {
	//order payments by date
	sort.Slice(payments, func(i, j int) bool { return payments[i].Before(payments[j]) })
	return &OrderedPayments{payments}
}

//Default implementation
type OrderedPayments struct {
	payments []IPayment
}

func (o *OrderedPayments) GetAll() []IPayment {
	return o.payments
}

func (o *OrderedPayments) Count() int {
	return len(o.payments)
}
