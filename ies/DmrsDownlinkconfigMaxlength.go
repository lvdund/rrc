package ies

import "rrc/utils"

// DMRS-DownlinkConfig-maxLength ::= ENUMERATED
type DmrsDownlinkconfigMaxlength struct {
	Value utils.ENUMERATED
}

const (
	DmrsDownlinkconfigMaxlengthEnumeratedNothing = iota
	DmrsDownlinkconfigMaxlengthEnumeratedLen2
)
