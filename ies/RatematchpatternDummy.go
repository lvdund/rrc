package ies

import "rrc/utils"

// RateMatchPattern-dummy ::= ENUMERATED
type RatematchpatternDummy struct {
	Value utils.ENUMERATED
}

const (
	RatematchpatternDummyEnumeratedNothing = iota
	RatematchpatternDummyEnumeratedDynamic
	RatematchpatternDummyEnumeratedSemistatic
)
