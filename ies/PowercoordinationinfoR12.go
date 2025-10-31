package ies

import "rrc/utils"

// PowerCoordinationInfo-r12 ::= SEQUENCE
type PowercoordinationinfoR12 struct {
	PMenbR12            utils.INTEGER `lb:0,ub:16`
	PSenbR12            utils.INTEGER `lb:0,ub:16`
	PowercontrolmodeR12 utils.INTEGER `lb:0,ub:2`
}
