// Copyright 2018 Andrey Z. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

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
