package ies

import "rrc/utils"

// SL-TrafficPatternInfo-r16 ::= SEQUENCE
type SlTrafficpatterninfoR16 struct {
	TrafficperiodicityR16 SlTrafficpatterninfoR16TrafficperiodicityR16
	TimingoffsetR16       utils.INTEGER   `lb:0,ub:10239`
	MessagesizeR16        utils.BITSTRING `lb:8,ub:8`
	SlQosFlowidentityR16  SlQosFlowidentityR16
}
