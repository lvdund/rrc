package ies

import "rrc/utils"

// FailedLogicalChannelIdentity-r16-cellGroupIndication-r16 ::= ENUMERATED
type FailedlogicalchannelidentityR16CellgroupindicationR16 struct {
	Value utils.ENUMERATED
}

const (
	FailedlogicalchannelidentityR16CellgroupindicationR16EnumeratedNothing = iota
	FailedlogicalchannelidentityR16CellgroupindicationR16EnumeratedMn
	FailedlogicalchannelidentityR16CellgroupindicationR16EnumeratedSn
)
