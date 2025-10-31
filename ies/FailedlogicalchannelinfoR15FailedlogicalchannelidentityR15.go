package ies

import "rrc/utils"

// FailedLogicalChannelInfo-r15-failedLogicalChannelIdentity-r15 ::= SEQUENCE
type FailedlogicalchannelinfoR15FailedlogicalchannelidentityR15 struct {
	CellgroupindicationR15       FailedlogicalchannelinfoR15FailedlogicalchannelidentityR15CellgroupindicationR15
	LogicalchannelidentityR15    *utils.INTEGER `lb:0,ub:10`
	LogicalchannelidentityextR15 *utils.INTEGER `lb:0,ub:38`
}
