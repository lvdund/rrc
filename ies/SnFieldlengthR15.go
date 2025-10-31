package ies

import "rrc/utils"

// SN-FieldLength-r15 ::= ENUMERATED
type SnFieldlengthR15 struct {
	Value utils.ENUMERATED
}

const (
	SnFieldlengthR15EnumeratedNothing = iota
	SnFieldlengthR15EnumeratedSize5
	SnFieldlengthR15EnumeratedSize10
	SnFieldlengthR15EnumeratedSize16_R15
)
