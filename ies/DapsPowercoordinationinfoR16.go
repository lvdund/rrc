package ies

import "rrc/utils"

// DAPS-PowerCoordinationInfo-r16 ::= SEQUENCE
type DapsPowercoordinationinfoR16 struct {
	PDapsSourceR16      utils.INTEGER `lb:0,ub:16`
	PDapsTargetR16      utils.INTEGER `lb:0,ub:16`
	PowercontrolmodeR16 utils.INTEGER `lb:0,ub:2`
}
