package secantModifiedMethod


type BorderIterationResult struct {
	nextBorder, ansBorder  Border
	isSolution, canContinue, tryLeft, tryRight bool
}

func NoSolutionAndBreak() BorderIterationResult {
	return BorderIterationResult{isSolution: false, canContinue:false}
}

func NoSolutionAndContinue(next Border, tryLeft, tryRight bool) BorderIterationResult {
	return BorderIterationResult{ nextBorder: next, isSolution: false, canContinue: true, tryLeft: tryLeft, tryRight: tryRight }
}

func SolutionAndContinue(next Border, ans Border, tryLeft, tryRight bool ) BorderIterationResult {
	return BorderIterationResult{next, ans, true, true, tryLeft, tryRight }
}
