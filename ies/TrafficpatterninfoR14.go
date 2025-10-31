package ies

import "rrc/utils"

// TrafficPatternInfo-r14 ::= SEQUENCE
type TrafficpatterninfoR14 struct {
	TrafficperiodicityR14       TrafficpatterninfoR14TrafficperiodicityR14
	TimingoffsetR14             utils.INTEGER `lb:0,ub:10239`
	PriorityinfoslR14           *SlPriorityR13
	LogicalchannelidentityulR14 *utils.INTEGER  `lb:0,ub:10`
	MessagesizeR14              utils.BITSTRING `lb:6,ub:6`
}
