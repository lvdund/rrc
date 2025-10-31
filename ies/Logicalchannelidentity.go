package ies

import "rrc/utils"

// LogicalChannelIdentity ::= utils.INTEGER (1..maxLC-ID)
type Logicalchannelidentity struct {
	Value utils.INTEGER `lb:0,ub:maxLCId`
}
