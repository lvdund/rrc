package ies

import "rrc/utils"

// UL-CyclicPrefixLength ::= ENUMERATED
type UlCyclicprefixlength struct {
	Value utils.ENUMERATED
}

const (
	UlCyclicprefixlengthLen1 = 0
	UlCyclicprefixlengthLen2 = 1
)
