package ies

import "rrc/utils"

// CSI-RS-ConfigZP-r11 ::= SEQUENCE
// Extensible
type CsiRsConfigzpR11 struct {
	CsiRsConfigzpidR11    CsiRsConfigzpidR11
	ResourceconfiglistR11 utils.BITSTRING `lb:16,ub:16`
	SubframeconfigR11     utils.INTEGER   `lb:0,ub:154`
}
