package ies

import "rrc/utils"

// DMRS-UplinkConfig-maxLength ::= ENUMERATED
type DmrsUplinkconfigMaxlength struct {
	Value utils.ENUMERATED
}

const (
	DmrsUplinkconfigMaxlengthEnumeratedNothing = iota
	DmrsUplinkconfigMaxlengthEnumeratedLen2
)
