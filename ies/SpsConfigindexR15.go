package ies

import "rrc/utils"

// SPS-ConfigIndex-r15 ::= utils.INTEGER (1..maxConfigSPS-r15)
type SpsConfigindexR15 struct {
	Value utils.INTEGER `lb:0,ub:maxConfigSPSR15`
}
