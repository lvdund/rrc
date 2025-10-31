package ies

import "rrc/utils"

// BWP-cyclicPrefix ::= ENUMERATED
type BwpCyclicprefix struct {
	Value utils.ENUMERATED
}

const (
	BwpCyclicprefixEnumeratedNothing = iota
	BwpCyclicprefixEnumeratedExtended
)
