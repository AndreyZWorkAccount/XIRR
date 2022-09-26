package xirrAsync

import (
	"github.com/krazybee/XIRR/netPresentValue"
)

type IRequest interface {
	RequestId() int64

	Payments() []netPresentValue.IPayment
}

func NewRequest(id int64, payments []netPresentValue.IPayment) IRequest {
	return &Request{id, payments}
}

//implementation

type Request struct {
	requestId int64
	payments  []netPresentValue.IPayment
}

func (r *Request) RequestId() int64 {
	return r.requestId
}

func (r *Request) Payments() []netPresentValue.IPayment {
	return r.payments
}
