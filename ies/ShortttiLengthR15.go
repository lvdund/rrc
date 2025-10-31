package ies

import "rrc/utils"

// ShortTTI-Length-r15 ::= ENUMERATED
type ShortttiLengthR15 struct {
	Value utils.ENUMERATED
}

const (
	ShortttiLengthR15EnumeratedNothing = iota
	ShortttiLengthR15EnumeratedSlot
	ShortttiLengthR15EnumeratedSubslot
)
