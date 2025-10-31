package ies

import "rrc/utils"

// MIMO-NonPrecodedCapabilities-r13-config3-r13 ::= ENUMERATED
type MimoNonprecodedcapabilitiesR13Config3R13 struct {
	Value utils.ENUMERATED
}

const (
	MimoNonprecodedcapabilitiesR13Config3R13EnumeratedNothing = iota
	MimoNonprecodedcapabilitiesR13Config3R13EnumeratedSupported
)
