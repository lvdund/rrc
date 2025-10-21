package ies

import "rrc/utils"

// FailedLogicalChannelInfo-r15 ::= SEQUENCE
type FailedlogicalchannelinfoR15 struct {
	FailedlogicalchannelidentityR15 *interface{}
	Failuretype                     utils.ENUMERATED
}
