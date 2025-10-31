package ies

import "rrc/utils"

// CRI-ConfigIndex-r13 ::= utils.INTEGER (0..1023)
type CriConfigindexR13 struct {
	Value utils.INTEGER `lb:0,ub:1023`
}
