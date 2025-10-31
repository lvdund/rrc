package ies

import "rrc/utils"

// RateMatchPatternLTE-CRS-v-Shift ::= ENUMERATED
type RatematchpatternlteCrsVShift struct {
	Value utils.ENUMERATED
}

const (
	RatematchpatternlteCrsVShiftEnumeratedNothing = iota
	RatematchpatternlteCrsVShiftEnumeratedN0
	RatematchpatternlteCrsVShiftEnumeratedN1
	RatematchpatternlteCrsVShiftEnumeratedN2
	RatematchpatternlteCrsVShiftEnumeratedN3
	RatematchpatternlteCrsVShiftEnumeratedN4
	RatematchpatternlteCrsVShiftEnumeratedN5
)
