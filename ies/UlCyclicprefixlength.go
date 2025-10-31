package ies

import "rrc/utils"

// UL-CyclicPrefixLength ::= ENUMERATED
type UlCyclicprefixlength struct {
	Value utils.ENUMERATED
}

const (
	UlCyclicprefixlengthEnumeratedNothing = iota
	UlCyclicprefixlengthEnumeratedLen1
	UlCyclicprefixlengthEnumeratedLen2
)
