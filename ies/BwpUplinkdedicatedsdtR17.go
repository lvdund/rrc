package ies

import "rrc/utils"

// BWP-UplinkDedicatedSDT-r17 ::= SEQUENCE
// Extensible
type BwpUplinkdedicatedsdtR17 struct {
	PuschConfigR17                        *utils.Setuprelease[PuschConfig]
	ConfiguredgrantconfigtoaddmodlistR17  *ConfiguredgrantconfigtoaddmodlistR16
	ConfiguredgrantconfigtoreleaselistR17 *ConfiguredgrantconfigtoreleaselistR16
}
