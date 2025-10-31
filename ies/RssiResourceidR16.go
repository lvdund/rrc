package ies

import "rrc/utils"

// RSSI-ResourceId-r16 ::= utils.INTEGER (0.. maxNrofCLI-RSSI-Resources-1-r16)
type RssiResourceidR16 struct {
	Value utils.INTEGER `lb:0,ub:maxNrofCLIRssiResources1R16`
}
