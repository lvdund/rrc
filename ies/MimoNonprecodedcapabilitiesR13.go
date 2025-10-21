package ies

import "rrc/utils"

// MIMO-NonPrecodedCapabilities-r13 ::= SEQUENCE
type MimoNonprecodedcapabilitiesR13 struct {
	Config1R13 *utils.ENUMERATED
	Config2R13 *utils.ENUMERATED
	Config3R13 *utils.ENUMERATED
	Config4R13 *utils.ENUMERATED
}
