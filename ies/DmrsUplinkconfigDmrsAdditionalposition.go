package ies

import "rrc/utils"

// DMRS-UplinkConfig-dmrs-AdditionalPosition ::= ENUMERATED
type DmrsUplinkconfigDmrsAdditionalposition struct {
	Value utils.ENUMERATED
}

const (
	DmrsUplinkconfigDmrsAdditionalpositionEnumeratedNothing = iota
	DmrsUplinkconfigDmrsAdditionalpositionEnumeratedPos0
	DmrsUplinkconfigDmrsAdditionalpositionEnumeratedPos1
	DmrsUplinkconfigDmrsAdditionalpositionEnumeratedPos3
)
