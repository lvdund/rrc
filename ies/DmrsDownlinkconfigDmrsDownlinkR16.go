package ies

import "rrc/utils"

// DMRS-DownlinkConfig-dmrs-Downlink-r16 ::= ENUMERATED
type DmrsDownlinkconfigDmrsDownlinkR16 struct {
	Value utils.ENUMERATED
}

const (
	DmrsDownlinkconfigDmrsDownlinkR16EnumeratedNothing = iota
	DmrsDownlinkconfigDmrsDownlinkR16EnumeratedEnabled
)
