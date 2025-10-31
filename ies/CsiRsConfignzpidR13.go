package ies

import "rrc/utils"

// CSI-RS-ConfigNZPId-r13 ::= utils.INTEGER (1..maxCSI-RS-NZP-r13)
type CsiRsConfignzpidR13 struct {
	Value utils.INTEGER `lb:0,ub:maxCSIRsNzpR13`
}
