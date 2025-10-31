package ies

import "rrc/utils"

// CSI-RS-ConfigNZPId-v1310 ::= utils.INTEGER (minCSI-RS-NZP-r13..maxCSI-RS-NZP-r13)
type CsiRsConfignzpidV1310 struct {
	Value utils.INTEGER `lb:minCSIRsNzpR13,ub:maxCSIRsNzpR13`
}
