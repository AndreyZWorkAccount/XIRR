// Copyright 2018 Andrey Z. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package secantAuto

import "github.com/krazybee/XIRR/netPresentValue"

const PositiveSum_LeftBoundary = -0.0001
const PositiveSum_RightBoundary = 0.20001

const NegativeOrZeroSum_LeftBoundary = -0.2001
const NegativeOrZeroSum_RightBoundary = 0.0001

const Delta_For_Borders = 0.1
const Delta_For_Borders_Multiplexer = 5

const PositiveSum_Initial_Guess = 0.4
const NegativeOrZeroSum_Initial_Guess = -0.4

const PositiveSum_MinLeft = 0.2
const PositiveSum_MaxRight = 100

const NegativeOrZeroSum_MinLeft = netPresentValue.IrrDefaultValue
const NegativeOrZeroSum_MaxRight = -0.2

const Max_Iterations = 100

const Iteration_Step = 0.02
