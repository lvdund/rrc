package ies

import "rrc/utils"

// DAPS-PowerCoordinationInfo-r16 ::= SEQUENCE
type DapsPowercoordinationinfoR16 struct {
	PDapsSourceR16      utils.INTEGER
	PDapsTargetR16      utils.INTEGER
	PowercontrolmodeR16 utils.INTEGER
}
