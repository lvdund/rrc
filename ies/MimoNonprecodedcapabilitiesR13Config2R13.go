package ies

import "rrc/utils"

// MIMO-NonPrecodedCapabilities-r13-config2-r13 ::= ENUMERATED
type MimoNonprecodedcapabilitiesR13Config2R13 struct {
	Value utils.ENUMERATED
}

const (
	MimoNonprecodedcapabilitiesR13Config2R13EnumeratedNothing = iota
	MimoNonprecodedcapabilitiesR13Config2R13EnumeratedSupported
)
