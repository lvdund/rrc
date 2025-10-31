package ies

import "rrc/utils"

// DMRS-DownlinkConfig-dmrs-Type ::= ENUMERATED
type DmrsDownlinkconfigDmrsType struct {
	Value utils.ENUMERATED
}

const (
	DmrsDownlinkconfigDmrsTypeEnumeratedNothing = iota
	DmrsDownlinkconfigDmrsTypeEnumeratedType2
)
