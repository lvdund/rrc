package ies

import "rrc/utils"

// CG-SDT-ConfigLCH-Restriction-r17 ::= SEQUENCE
type CgSdtConfiglchRestrictionR17 struct {
	LogicalchannelidentityR17      Logicalchannelidentity
	Configuredgranttype1allowedR17 *utils.BOOLEAN
	AllowedcgListR17               *[]ConfiguredgrantconfigindexmacR16 `lb:0,ub:maxNrofConfiguredGrantConfigMAC1R16`
}
