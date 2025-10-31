package ies

import "rrc/utils"

// SL-QoS-FlowIdentity-r16 ::= utils.INTEGER (1..maxNrofSL-QFIs-r16)
type SlQosFlowidentityR16 struct {
	Value utils.INTEGER `lb:0,ub:maxNrofSLQfisR16`
}
