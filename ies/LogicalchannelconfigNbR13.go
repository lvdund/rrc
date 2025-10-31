package ies

import "rrc/utils"

// LogicalChannelConfig-NB-r13 ::= SEQUENCE
// Extensible
type LogicalchannelconfigNbR13 struct {
	PriorityR13                 *utils.INTEGER `lb:0,ub:16`
	LogicalchannelsrProhibitR13 *utils.BOOLEAN
}
