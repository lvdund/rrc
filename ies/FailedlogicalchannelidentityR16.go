package ies

import "rrc/utils"

// FailedLogicalChannelIdentity-r16 ::= SEQUENCE
type FailedlogicalchannelidentityR16 struct {
	CellgroupindicationR16       utils.ENUMERATED
	LogicalchannelidentityR16    *utils.INTEGER
	LogicalchannelidentityextR16 *utils.INTEGER
}
