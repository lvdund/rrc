package ies

import "rrc/utils"

// CSI-IM-ConfigId-r11 ::= utils.INTEGER (1..maxCSI-IM-r11)
type CsiImConfigidR11 struct {
	Value utils.INTEGER `lb:0,ub:maxCSIImR11`
}
