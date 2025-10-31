package ies

import "rrc/utils"

// MIMO-NonPrecodedCapabilities-r13-config1-r13 ::= ENUMERATED
type MimoNonprecodedcapabilitiesR13Config1R13 struct {
	Value utils.ENUMERATED
}

const (
	MimoNonprecodedcapabilitiesR13Config1R13EnumeratedNothing = iota
	MimoNonprecodedcapabilitiesR13Config1R13EnumeratedSupported
)
