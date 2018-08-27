package xirrAsync


//Factory method
func NewProcessor() IProcessor{

	requests := make(chan IRequest)
	responses := make(chan IResponse)
	cancellations := make([] chan interface{}, 0)

	return &RequestsProcessor{ requests,  responses, cancellations}
}