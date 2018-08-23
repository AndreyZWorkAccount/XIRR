// Copyright 2018 Andrey Z. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package secantAuto

type IBordersSearchAlgorithm interface {
	//Returns a slice that contains optimal initial borders for 'secant' method applying
	FindInitialBorders(paymentsSumIsPositive bool) []IBorder
}

type IBorder interface {
	Left() float64
	Right() float64
}
