package secantModifiedMethod

import . "math"

type ValidationOptions struct {
	useBorders bool
	validate bool
	minLeft float64
	maxRight float64
}

func NoValidation(useBorders bool) ValidationOptions {
	return ValidationOptions{validate:false, useBorders:useBorders, minLeft:Inf(-1), maxRight:Inf(1)}
}

func UseValidation(useBorders bool, minLeft float64, maxRight float64) ValidationOptions  {
	return ValidationOptions{validate:true, useBorders:useBorders, minLeft:minLeft, maxRight:maxRight}
}
