package ies

import "rrc/utils"

// CSI-IM-ConfigId-v1250 ::= utils.INTEGER (maxCSI-IM-r12)
type CsiImConfigidV1250 struct {
	Value utils.INTEGER `lb:maxCSIImR12,ub:maxCSIImR12`
}
