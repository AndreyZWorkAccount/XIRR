// Copyright 2018 Andrey Z. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package float_Extensions

import . "math"

func Average(a float64, b float64) float64{
	return a + (b - a)/2
}

func AnyNanOrInfinity(numbers ...float64) bool{
	for _,num := range numbers{
		if IsInf(num,0) || IsNaN(num){
			return true
		}
	}
	return false
}
