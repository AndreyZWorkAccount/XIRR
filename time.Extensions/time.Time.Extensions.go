// Copyright 2018 Andrey Z. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package time_Extensions

import (
	"time"
	)

//Returns difference between two dates (in days)
func DiffInDays(startDate *time.Time, endDate *time.Time) float64{
	if startDate == nil || endDate == nil{
		return 0
	}
	return endDate.Sub(*startDate).Hours()/24
}
