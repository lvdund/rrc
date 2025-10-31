package ies

import "rrc/utils"

// DMRS-DownlinkConfig-dmrs-AdditionalPosition ::= ENUMERATED
type DmrsDownlinkconfigDmrsAdditionalposition struct {
	Value utils.ENUMERATED
}

const (
	DmrsDownlinkconfigDmrsAdditionalpositionEnumeratedNothing = iota
	DmrsDownlinkconfigDmrsAdditionalpositionEnumeratedPos0
	DmrsDownlinkconfigDmrsAdditionalpositionEnumeratedPos1
	DmrsDownlinkconfigDmrsAdditionalpositionEnumeratedPos3
)
