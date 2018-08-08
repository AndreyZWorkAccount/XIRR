package numMethods

import (
	"fmt"
)

type NumericMethodError struct {
	NumericMethodName string
	Description string
}
func (e *NumericMethodError) Error() string {
	return fmt.Sprintf("The numerical method %v has been terminated.", e.NumericMethodName)
}