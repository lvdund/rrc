package ies

import "rrc/utils"

// CSI-IM-ConfigId-v1310 ::= utils.INTEGER (minCSI-IM-r13..maxCSI-IM-r13)
type CsiImConfigidV1310 struct {
	Value utils.INTEGER `lb:minCSIImR13,ub:maxCSIImR13`
}
