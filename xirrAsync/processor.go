package xirrAsync

type IProcessor interface {

	Requests() chan IRequest

	Responses() <-chan IResponse

	Start(coresCount int)

	Stop() <- chan bool
}


