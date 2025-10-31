package ies

import "rrc/utils"

// MIB-dmrs-TypeA-Position ::= ENUMERATED
type MibDmrsTypeaPosition struct {
	Value utils.ENUMERATED
}

const (
	MibDmrsTypeaPositionEnumeratedNothing = iota
	MibDmrsTypeaPositionEnumeratedPos2
	MibDmrsTypeaPositionEnumeratedPos3
)
