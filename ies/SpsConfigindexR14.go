package ies

import "rrc/utils"

// SPS-ConfigIndex-r14 ::= utils.INTEGER (1..maxConfigSPS-r14)
type SpsConfigindexR14 struct {
	Value utils.INTEGER `lb:0,ub:maxConfigSPSR14`
}
