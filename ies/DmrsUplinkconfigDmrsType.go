package ies

import "rrc/utils"

// DMRS-UplinkConfig-dmrs-Type ::= ENUMERATED
type DmrsUplinkconfigDmrsType struct {
	Value utils.ENUMERATED
}

const (
	DmrsUplinkconfigDmrsTypeEnumeratedNothing = iota
	DmrsUplinkconfigDmrsTypeEnumeratedType2
)
