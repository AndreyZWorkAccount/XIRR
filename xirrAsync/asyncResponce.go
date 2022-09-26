package xirrAsync

import (
	"github.com/krazybee/XIRR/numMethods"
)

type IResponse interface {
	RequestId() int64

	Result() numMethods.IResult
}

func NewResponse(id int64, res numMethods.IResult) IResponse {
	return &Response{id, res}
}

//implementation

type Response struct {
	requestId int64
	result    numMethods.IResult
}

func (r *Response) RequestId() int64 {
	return r.requestId
}
func (r *Response) Result() numMethods.IResult {
	return r.result
}
