package ies

import "rrc/utils"

// ShortTTI-Length-r15 ::= ENUMERATED
type ShortttiLengthR15 struct {
	Value utils.ENUMERATED
}

const (
	ShortttiLengthR15Slot    = 0
	ShortttiLengthR15Subslot = 1
)
