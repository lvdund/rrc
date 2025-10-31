package ies

import "rrc/utils"

// CSI-RS-ConfigZPId-r11 ::= utils.INTEGER (1..maxCSI-RS-ZP-r11)
type CsiRsConfigzpidR11 struct {
	Value utils.INTEGER `lb:0,ub:maxCSIRsZpR11`
}
