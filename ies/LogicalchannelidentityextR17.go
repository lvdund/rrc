package ies

import "rrc/utils"

// LogicalChannelIdentityExt-r17 ::= utils.INTEGER (320..65855)
type LogicalchannelidentityextR17 struct {
	Value utils.INTEGER `lb:0,ub:65855`
}
