package ies

import "rrc/utils"

// PowerCoordinationInfo-r12 ::= SEQUENCE
type PowercoordinationinfoR12 struct {
	PMenbR12            utils.INTEGER
	PSenbR12            utils.INTEGER
	PowercontrolmodeR12 utils.INTEGER
}
