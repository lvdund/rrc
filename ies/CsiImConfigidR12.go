package ies

import "rrc/utils"

// CSI-IM-ConfigId-r12 ::= utils.INTEGER (1..maxCSI-IM-r12)
type CsiImConfigidR12 struct {
	Value utils.INTEGER `lb:0,ub:maxCSIImR12`
}
