package ies

import "rrc/utils"

// FailedLogicalChannelIdentity-r16 ::= SEQUENCE
type FailedlogicalchannelidentityR16 struct {
	CellgroupindicationR16       FailedlogicalchannelidentityR16CellgroupindicationR16
	LogicalchannelidentityR16    *utils.INTEGER `lb:0,ub:10`
	LogicalchannelidentityextR16 *utils.INTEGER `lb:0,ub:38`
}
