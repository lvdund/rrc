package ies

import "rrc/utils"

// SL-SRAP-Config-r17 ::= SEQUENCE
// Extensible
type SlSrapConfigR17 struct {
	SlLocalidentityR17        *utils.INTEGER             `lb:0,ub:255`
	SlMappingtoaddmodlistR17  *[]SlMappingtoaddmodR17    `lb:1,ub:maxLCId`
	SlMappingtoreleaselistR17 *[]SlRemoteueRbIdentityR17 `lb:1,ub:maxLCId`
}
