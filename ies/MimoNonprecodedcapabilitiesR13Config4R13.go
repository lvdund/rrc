package ies

import "rrc/utils"

// MIMO-NonPrecodedCapabilities-r13-config4-r13 ::= ENUMERATED
type MimoNonprecodedcapabilitiesR13Config4R13 struct {
	Value utils.ENUMERATED
}

const (
	MimoNonprecodedcapabilitiesR13Config4R13EnumeratedNothing = iota
	MimoNonprecodedcapabilitiesR13Config4R13EnumeratedSupported
)
