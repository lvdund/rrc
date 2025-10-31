package ies

import "rrc/utils"

// CSI-IM-ConfigId-r13 ::= utils.INTEGER (1..maxCSI-IM-r13)
type CsiImConfigidR13 struct {
	Value utils.INTEGER `lb:0,ub:maxCSIImR13`
}
