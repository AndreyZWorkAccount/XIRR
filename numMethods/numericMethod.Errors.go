package numMethods

import (
	. "fmt"
)

type NumericMethodError struct {
	Description string
}
func (e *NumericMethodError) Error() string {
	return Sprintf("The numerical method has been terminated: %v.", e.Description)
}
