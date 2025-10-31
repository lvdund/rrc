package ies

import "rrc/utils"

// CSI-RS-ConfigNZPId-r11 ::= utils.INTEGER (1..maxCSI-RS-NZP-r11)
type CsiRsConfignzpidR11 struct {
	Value utils.INTEGER `lb:0,ub:maxCSIRsNzpR11`
}
